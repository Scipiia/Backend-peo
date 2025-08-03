package getorder

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"net/http"
	"strconv"
	"vue-golang/internal/storage"
)

type Request struct {
	ID int `json:"id"`
}

type Response struct {
	ID       int    `json:"id"`
	OrderNum string `json:"order_num"`
	Creator  int    `json:"creator"`
	Customer string `json:"customer"`
	DopInfo  string `json:"dop_info"`
	MsNote   string `json:"ms_note"`

	OrderDemPrice []*storage.OrderDemPrice `json:"order_dem_price"`
	ImageBase64   string                   `json:"image_base_64"`

	Error  string `json:"error"`
	Status string `json:"status"`
}

type OrderDetails interface {
	GetOrderDetails(id int) (*storage.ResultOrderDetails, error)
	GetForm(id int) (*storage.FormPeo, error)
	GetWorkers() ([]*storage.Workers, error)
	GetGlyhari(id int) (*storage.OrderData, error)
	GetWindows(id int) (*storage.OrderData, error)
	GetDoor(id int) (*storage.OrderData, error)
}

// Общая функция для получения данных о заказе
func getOrderDetails(log *slog.Logger, order OrderDetails, id int) (*storage.ResultOrderDetails, error) {
	const op = "handler.get-order-details"

	log = log.With(
		slog.String("op", op),
	)

	details, err := order.GetOrderDetails(id)
	if err != nil {
		log.Info("Order not found", slog.Int("id", id), slog.String("error", err.Error()))
		return nil, fmt.Errorf("order not found: %w", err)
	}

	return details, nil
}

func OrderDetailsMiddleware(log *slog.Logger, order OrderDetails) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем ID заказа из параметров запроса
			//idStr := r.URL.Query().Get("id")
			idStr := chi.URLParam(r, "id")
			if idStr == "" {
				log.Error("Missing order ID in query parameters")
				http.Error(w, "Missing order ID", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid order ID", http.StatusBadRequest)
				log.Error("Invalid order ID", slog.String("error", err.Error()))
				return
			}

			// Получаем данные о заказе
			details, err := getOrderDetails(log, order, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			// Сохраняем данные о заказе в контексте
			ctx := context.WithValue(r.Context(), "orderDetails", details)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Обработчик New
func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		details, ok := r.Context().Value("orderDetails").(*storage.ResultOrderDetails)
		if !ok {
			http.Error(w, "Order details not found in context pizda", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, Response{
			ID:            details.Order.ID,
			OrderNum:      details.Order.OrderNum,
			Creator:       details.Order.Creator,
			Customer:      details.Order.Customer,
			DopInfo:       details.Order.DopInfo,
			MsNote:        details.Order.MsNote,
			OrderDemPrice: details.OrderDemPrice, // Массив планов
			Error:         "",
			Status:        strconv.Itoa(http.StatusOK),
		})

	}
}

type ResponseForm struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Fields []storage.FieldPeo `json:"fields"`
}

func GetFormByID(log *slog.Logger, form OrderDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Start Forms")

		//formID := chi.URLParam(r, "id")
		formID := r.URL.Query().Get("idForm")
		log.Info("IDDDD", formID)
		if formID == "" {
			log.Error("Missingsdfsdf order IDFOrm in query parameters")
			http.Error(w, "Missing order ID FORM syka bleat", http.StatusBadRequest)
			return
		}

		log.Info("Fetching form by ID", slog.String("formID", formID))

		id, err := strconv.Atoi(formID)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			log.Error("Invalid order ID", slog.String("error", err.Error()))
			return
		}

		getForm, err := form.GetForm(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Формируем ответ
		render.JSON(w, r, ResponseForm{
			ID:     getForm.ID,
			Name:   getForm.Name,
			Fields: getForm.FieldsPeo,
		})

		log.Info("Successfully fetched form", slog.String("formID", formID))

	}
}

type ResponseWorkers struct {
	Workers []*storage.Workers
}

func GetWorkers(log *slog.Logger, worker OrderDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		workers, err := worker.GetWorkers()
		if err != nil {
			log.Info("Не удалось вытащить работяг и базы данных", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		log.Info("WORKERS", workers)
		// Формируем ответ
		render.JSON(w, r, ResponseWorkers{
			Workers: workers,
		})
	}
}

type ResponseNormData struct {
	OrderNormData *storage.OrderData
}

func GetNormOrders(log *slog.Logger, normData OrderDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if idStr == "" {
			log.Error("Missing order ID in query parameters")
			http.Error(w, "Missing order ID", http.StatusBadRequest)
			return
		}

		productType := r.URL.Query().Get("type")
		if productType == "" {
			log.Error("Missing 'type' parameter")
			render.JSON(w, r, Response{Error: "Missing 'type' parameter"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			log.Error("Invalid order ID", slog.String("error", err.Error()))
			return
		}

		fmt.Println("TTTTTTYYYYYPE", productType)
		fmt.Println("TTTTTTYYYYYPE111111", id)

		var result *storage.OrderData
		var Fetherr error

		switch productType {
		case "glyhar":
			result, Fetherr = normData.GetGlyhari(id)
		case "window":
			//TODO ljltkfnm nen
			result, Fetherr = normData.GetWindows(id)
		case "door":
			result, Fetherr = normData.GetDoor(id)
		}

		//glyhari, err := normData.GetGlyhari(id)
		//if err != nil {
		//	log.Info("Не удалось вытащить нормированный наряд из базы данных", err)
		//	render.JSON(w, r, Response{Error: "da bleeeat"})
		//	return
		//}
		//log.Info("GGGGLYHARI", glyhari)
		log.Info("RRRRRRRRRRRRRRRREEEE", result)

		if Fetherr != nil {
			log.Error("Product not found", slog.Int("id", id), slog.String("type", productType), Fetherr)
			render.JSON(w, r, Response{Error: "Product not found"})
			return
		}

		//render.JSON(w, r, ResponseNormData{
		//	OrderNormData: result,
		//})
		render.JSON(w, r, result)
	}
}

// Обработчик GenerateExcel
func GenerateExcel(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем детали заказа из контекста
		details, ok := r.Context().Value("orderDetails").(*storage.Order)
		if !ok {
			log.Error("Order details not found in context")
			http.Error(w, "Order details not found in context", http.StatusInternalServerError)
			return
		}

		log.Info("Generating Excel file for order", slog.Int("orderID", details.ID))

		// Создаем новый Excel-файл
		f := excelize.NewFile()
		sheetName := "OrderDetails"
		f.SetSheetName("Sheet1", sheetName)

		// Записываем заголовки
		headers := []string{"Field", "Value"}
		for col, header := range headers {
			cell := fmt.Sprintf("%s1", string('A'+col))
			f.SetCellValue(sheetName, cell, header)
		}

		// Подготавливаем данные для записи
		type Field struct {
			Name  string
			Value string
		}

		fields := []Field{
			{"ID", strconv.Itoa(details.ID)},
			{"OrderNum", details.OrderNum},
			{"Creator", strconv.Itoa(details.Creator)},
			{"Customer", details.Customer},
			{"DopInfo", details.DopInfo},
			{"MsNote", details.MsNote},
		}

		// Записываем данные
		row := 2
		for _, field := range fields {
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), field.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), field.Value)
			row++
		}

		// Устанавливаем заголовки HTTP
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Set("Content-Disposition", "attachment; filename=order_details.xlsx")
		w.WriteHeader(http.StatusOK)

		// Отправляем файл клиенту
		if err := f.Write(w); err != nil {
			log.Error("Failed to write Excel file", slog.String("error", err.Error()))
			http.Error(w, "Failed to generate Excel file", http.StatusInternalServerError)
			return
		}
	}
}
