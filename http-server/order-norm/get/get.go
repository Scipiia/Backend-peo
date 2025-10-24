package get

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vue-golang/internal/storage"
	"vue-golang/internal/storage/mysql"
)

type ResultGetNorm interface {
	GetNormOrder(id int64) (*storage.GetOrderDetails, error)
	GetNormOrders(orderNum, orderType string) ([]storage.GetOrderDetails, error)
	GetSimpleOrderReport(orderNum string) (*storage.OrderFinalReport, error)
	GetFinalNormOrders() ([]storage.ReportFinalOrders, error)
	GetNormOrdersByOrderNum(orderNum string) ([]*storage.GetOrderDetails, error)
	GetNormOrderIdSub(id int64) ([]*storage.GetOrderDetails, error)

	//TODO new logic
	GetPEOProductsByCategory(filter mysql.ProductFilter) ([]storage.PEOProduct, []storage.GetWorkers, error)
}

func GetNormOrder(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-dem-norm.get.GetNormOrder"

		// Извлекаем id из URL
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		log.Info("Получение нормировки", slog.Int64("id", id))

		norm, err := result.GetNormOrder(id)
		if err != nil {
			if strings.Contains(err.Error(), "не найдена") {
				log.Info("Ошибка реквеста сообщения при вставке в базу заказа сука блять уебище тупорылое DOOR ебаные", err)
				http.Error(w, "Нормировка не найдена", http.StatusNotFound)
				return
			}
			log.Error("Ошибка получения нормировки", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
			return
		}

		// Успешный ответ
		render.JSON(w, r, norm)
	}
}

func GetNormOrdersOrderNum(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.GetNormOrders"

		orderNum := r.URL.Query().Get("order_num")
		//orderNum := chi.URLParam(r, "order_num")

		log.With(
			slog.String("op", op),
			slog.String("order_num", orderNum),
		).Info("Запрос на получение заказов")

		orders, err := result.GetNormOrdersByOrderNum(orderNum)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении нормировок по номеру заказа")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, orders)
	}
}

func GetNormOrders(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.GetNormOrders"

		// Получаем фильтр
		orderNum := r.URL.Query().Get("order_num")
		orderType := r.URL.Query().Get("type")

		log.With(
			slog.String("op", op),
			slog.String("order_num_filter", orderNum),
			slog.String("order_type_filter", orderType),
		).Info("Запрос на получение заказов")

		// Передаём фильтр (может быть пустым)
		items, err := result.GetNormOrders(orderNum, orderType)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении заказов")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		log.With(slog.Int("found", len(items))).Info("Заказы найдены")

		// Возвращаем JSON
		render.JSON(w, r, items)
	}
}

func FinalReportNormOrder(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.FinalReportNormOrder"

		orderNum := chi.URLParam(r, "order_num")

		log.Info("Получение нормировки", slog.String("orderNum", orderNum))

		report, err := result.GetSimpleOrderReport(orderNum)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении заказов по номеру заказа")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		log.Info("ASASASSASTTTTTTT", report)

		render.JSON(w, r, report)
	}
}

func FinalReportNormOrders(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.FinalReportNormOrders"

		// Парсим query-параметры
		fromStr := r.URL.Query().Get("from") // формат: 2025-04-01
		toStr := r.URL.Query().Get("to")
		orderNum := r.URL.Query().Get("order_num")
		typeIzd := r.URL.Query()["type"]

		var from, to time.Time
		//var err error

		parseDate := func(dateStr string, defaultTime time.Time) (time.Time, error) {
			if dateStr == "" {
				return defaultTime, nil
			}
			return time.Parse("2006-01-02", dateStr)
		}

		// По умолчанию: начало и конец текущего месяца
		now := time.Now()
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second) // последний момент предыдущего месяца

		from, err := parseDate(fromStr, startOfMonth)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Warn("Неверный формат from")
			http.Error(w, "Неверный формат даты 'from'", http.StatusBadRequest)
			return
		}

		to, err = parseDate(toStr, endOfMonth)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Warn("Неверный формат to")
			http.Error(w, "Неверный формат даты 'to'", http.StatusBadRequest)
			return
		}

		// Формируем фильтр
		filter := mysql.ProductFilter{
			From:     from,
			To:       to,
			OrderNum: orderNum,
			Type:     typeIzd,
		}

		// Запрашиваем данные
		products, employees, err := result.GetPEOProductsByCategory(filter)
		if err != nil {
			log.With(slog.String("op", op), slog.Any("error", err)).Error("Ошибка при получении изделий")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Отправить как JSON:
		response := map[string]interface{}{
			"employees": employees,
			"products":  products,
		}

		log.Info("REEEESSPPPP", response)

		render.JSON(w, r, response)
	}
}

func DoubleReportOrder(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.DoubleReportOrder"

		// Извлекаем id из URL
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		log.Info("Получение нормировки", slog.Int64("id", id))

		sub, err := result.GetNormOrderIdSub(id)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении заказов по номеру заказа")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, sub)
	}
}
