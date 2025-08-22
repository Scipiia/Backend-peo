package getorder

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"net/http"
	"time"
	"vue-golang/internal/storage"
)

type AllNormOrders interface {
	GetAllProducts(from, to, orderNum, itemType, profil, name string) ([]storage.ProductItem, error)
}
type ResponseNormOrders struct {
	Orders []storage.ProductItem `json:"orders"`
	Error  string                `json:"error"`
}

func GetAllNormOrders(log *slog.Logger, normOrders AllNormOrders) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		orderNum := r.URL.Query().Get("order_num")
		itemType := r.URL.Query().Get("type")
		profil := r.URL.Query().Get("profil")
		name := r.URL.Query().Get("name")

		orders, err := normOrders.GetAllProducts(from, to, orderNum, itemType, profil, name)
		if err != nil {
			log.Info("Не удалось вытащить заказы из базы данных", err)
			render.JSON(w, r, Response{Error: "da bleeeat, нормированные заказы не удалось вытащить"})
			return
		}
		log.Info("FROM", from)
		log.Info("TOOO", to)
		log.Info("orderNum", orderNum)
		log.Info("profil", profil)
		log.Info("name", name)

		//log.Info("ORDERS", orders)
		// Формируем ответ
		render.JSON(w, r, ResponseNormOrders{
			Orders: orders,
			Error:  "",
		})
	}
}

// ExportNormOrders — выгружает нормированные наряды в Excel
// В шапке: Лоджия, Витраж, Дверь, Окно, Глухарь
// В ячейке: название типа (например, "Лоджия"), если тип совпадает с колонкой
func ExportNormOrders(log *slog.Logger, normOrders AllNormOrders) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Парсим фильтры
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		orderNum := r.URL.Query().Get("order_num")
		itemType := r.URL.Query().Get("type")
		profil := r.URL.Query().Get("profil")
		name := r.URL.Query().Get("name")

		log.Info("Экспорт в Excel",
			"from", from,
			"to", to,
			"order_num", orderNum,
			"type", itemType,
			"profil", profil,
			"name", name,
		)

		// Получаем данные
		products, err := normOrders.GetAllProducts(from, to, orderNum, itemType, profil, name)
		if err != nil {
			log.Error("Не удалось получить данные для экспорта", "error", err)
			http.Error(w, "Не удалось получить данные", http.StatusInternalServerError)
			return
		}

		// Создаём Excel
		file := excelize.NewFile()
		defer func() {
			_ = file.Close()
		}()

		sheetName := "Наряды"
		file.NewSheet(sheetName)
		if _, err := file.GetSheetIndex("Sheet1"); err == nil {
			file.DeleteSheet("Sheet1")
		}
		file.SetActiveSheet(0)

		// --- Сопоставление типа из БД и отображаемого названия ---
		typeMapping := map[string]string{
			"loggia": "Лоджия",
			"vitraj": "Витраж",
			"door":   "Дверь",
			"window": "Окно",
			"glyhar": "Глухарь",
		}

		// Порядок колонок
		itemTypes := []string{"loggia", "vitraj", "door", "window", "glyhar"}

		// --- Заголовки ---
		headers := []string{"№", "Номер наряда", "Наименование", "Кол-во", "Профиль", "Время (ч)", "Дата создания", "Дата обновления"}
		for _, t := range itemTypes {
			headers = append(headers, typeMapping[t]) // "Лоджия", "Витраж" и т.д.
		}

		// Записываем заголовки
		for col, h := range headers {
			cell, _ := excelize.CoordinatesToCellName(col+1, 1)
			file.SetCellValue(sheetName, cell, h)
		}

		// Стиль: жирная шапка
		boldStyle, _ := file.NewStyle(&excelize.Style{
			Font: &excelize.Font{Bold: true},
		})
		file.SetCellStyle(sheetName, "A1", "Z1", boldStyle)

		// Автофильтр
		_ = file.AutoFilter(sheetName, "A1:"+headers[len(headers)-1]+"1", nil)

		// --- Автоширина колонок ---
		_ = file.SetColWidth(sheetName, "A", "A", 5)  // №
		_ = file.SetColWidth(sheetName, "B", "B", 12) // Наряд
		_ = file.SetColWidth(sheetName, "C", "C", 50) // Наименование
		_ = file.SetColWidth(sheetName, "D", "D", 8)  // Кол-во
		_ = file.SetColWidth(sheetName, "E", "E", 30) // Профиль
		_ = file.SetColWidth(sheetName, "F", "F", 10) // Время
		_ = file.SetColWidth(sheetName, "G", "G", 15) // Дата созд
		_ = file.SetColWidth(sheetName, "G", "H", 15) // Дата обнов
		for i := 0; i < len(itemTypes); i++ {
			col, _ := excelize.ColumnNumberToName(i + 8) // H и дальше
			_ = file.SetColWidth(sheetName, col, col, 12)
		}

		// --- Заполняем строки ---
		for i, p := range products {
			rowNum := i + 2

			// №
			file.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNum), i+1)
			// Номер наряда
			file.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNum), p.OrderNum)
			// Наименование
			file.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNum), p.Name)
			// Кол-во
			file.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNum), p.Count)
			// Профиль
			file.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNum), p.Profil)
			// Время (ч)
			file.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNum), fmt.Sprintf("%.1f ч", p.TotalTime))
			// Дата создания
			file.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNum), p.CreatedAt.Format("02.01.2006"))
			// Дата обновления
			file.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNum), p.UpdatedAt.Format("02.01.2006"))

			// Определяем, в какую колонку писать тип
			displayName, exists := typeMapping[p.Type]
			if !exists {
				continue // если тип неизвестен — пропускаем
			}

			// Находим индекс колонки (H = 8, т.к. A=1)
			for j, typeName := range itemTypes {
				if p.Type == typeName {
					colIndex := j + 9 // I — 9-я колонка
					cell, _ := excelize.CoordinatesToCellName(colIndex, rowNum)
					file.SetCellValue(sheetName, cell, displayName) // пишем "Лоджия", а не "loggia"
					break
				}
			}
		}

		// Формируем имя файла
		baseName := "нормированные_наряды"
		var period string

		if from != "" {
			if t, err := time.Parse("2006-01-02", from); err == nil {
				ruMonth := map[time.Month]string{
					time.January:   "январь",
					time.February:  "февраль",
					time.March:     "март",
					time.April:     "апрель",
					time.May:       "май",
					time.June:      "июнь",
					time.July:      "июль",
					time.August:    "август",
					time.September: "сентябрь",
					time.October:   "октябрь",
					time.November:  "ноябрь",
					time.December:  "декабрь",
				}[t.Month()]
				period = fmt.Sprintf("%s_%d", ruMonth, t.Year())
			}
		}
		if period == "" {
			now := time.Now()
			monthName := [...]string{"январь", "февраль", "март", "апрель", "май", "июнь",
				"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}
			period = fmt.Sprintf("%s_%d", monthName[now.Month()-1], now.Year())
		}

		filename := fmt.Sprintf("%s_%s.xlsx", baseName, period)

		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		// --- Отправка файла ---
		//w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		//w.Header().Set("Content-Disposition", `attachment; filename="нормированные_наряды.xlsx"`)

		if err := file.Write(w); err != nil {
			log.Error("Ошибка при записи Excel", "error", err)
			http.Error(w, "Не удалось сгенерировать файл", http.StatusInternalServerError)
			return
		}
	}
}
