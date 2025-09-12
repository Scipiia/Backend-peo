package get

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"net/http"
	"strconv"
	"vue-golang/internal/storage"
)

type ResponseOrders struct {
	Orders []*storage.Order `json:"orders"`
	Status string           `json:"status"`
	Error  string           `json:"error"`
}

type GetOrders interface {
	GetOrdersMonth(year int, month int) ([]*storage.Order, error)
}

func GetOrdersFilter(log *slog.Logger, getOrders GetOrders) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.orders.orders.GetOrdersFilter"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Получаем параметры year и month из запроса
		yearStr := r.URL.Query().Get("year")
		monthStr := r.URL.Query().Get("month")

		fmt.Println(yearStr, monthStr)
		log.Info("DATAFONK", yearStr, monthStr)

		if yearStr == "" || monthStr == "" {
			log.Error("Missing year or month in query parameters")
			http.Error(w, "Missing year or month", http.StatusBadRequest)
			return
		}

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			log.Error("Invalid year", slog.String("error", err.Error()))
			http.Error(w, "Invalid year", http.StatusBadRequest)
			return
		}

		month, err := strconv.Atoi(monthStr)
		if err != nil {
			log.Error("Invalid month", slog.String("error", err.Error()))
			http.Error(w, "Invalid month", http.StatusBadRequest)
			return
		}
		//year = 2025
		//month = 6

		orders, err := getOrders.GetOrdersMonth(year, month)
		if err != nil {
			log.Info("sql no rows", err)
			render.JSON(w, r, ResponseOrders{Error: "В базе не найден данный заказ"})
			return
		}

		log.Info("SASAT", orders)

		render.JSON(w, r, ResponseOrders{
			Orders: orders,
			Status: strconv.Itoa(http.StatusOK),
		})
	}
}

////TODO новая логика с мидлеваре

func getOrdersMonth(log *slog.Logger, orders GetOrders, year, month int) ([]*storage.Order, error) {
	const op = "handler.orders.orders.go"

	log = log.With(
		slog.String("op", op),
	)

	ordersMonth, err := orders.GetOrdersMonth(year, month)
	if err != nil {
		log.Info("Order not found", slog.Int("year:", year), slog.Int("month", month), slog.String("error", err.Error()))
		return nil, fmt.Errorf("order-dem-norm not found: %w", err)
	}

	return ordersMonth, nil
}

func OrdersMonthMiddleware(log *slog.Logger, orders GetOrders) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем параметры year и month из запроса
			yearStr := r.URL.Query().Get("year")
			monthStr := r.URL.Query().Get("month")

			if yearStr == "" || monthStr == "" {
				log.Error("Missing year or month in query parameters")
				http.Error(w, "Missing year or month", http.StatusBadRequest)
				return
			}

			year, err := strconv.Atoi(yearStr)
			if err != nil {
				log.Error("Invalid year", slog.String("error", err.Error()))
				http.Error(w, "Invalid year", http.StatusBadRequest)
				return
			}

			month, err := strconv.Atoi(monthStr)
			if err != nil {
				log.Error("Invalid month", slog.String("error", err.Error()))
				http.Error(w, "Invalid month", http.StatusBadRequest)
				return
			}

			log.Info("DATE", year, month)

			ordersMonth, err := getOrdersMonth(log, orders, year, month)
			if err != nil {
				log.Info("sql no rows", err)
				render.JSON(w, r, ResponseOrders{Error: "В базе не найден данный заказ"})
				return
			}

			ctx := context.WithValue(r.Context(), "ordersMonth", ordersMonth)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GenerateExel(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем массив заказов из контекста
		orders, ok := r.Context().Value("ordersMonth").([]*storage.Order)
		if !ok || len(orders) == 0 {
			log.Error("Orders not found in context or empty")
			http.Error(w, "No orders found for the specified period", http.StatusNotFound)
			return
		}

		// Создаем новый файл Excel
		f := excelize.NewFile()
		sheetName := "Orders"
		f.SetSheetName("Sheet1", sheetName)

		// Записываем заголовки
		headers := []string{"ID", "NumFer", "ClassID", "Ordername", "EnginerID"}
		for col, header := range headers {
			cell := fmt.Sprintf("%s1", string('A'+col))
			f.SetCellValue(sheetName, cell, header)
		}

		// Записываем данные заказов
		for i, order := range orders {
			row := i + 2 // Начинаем со второй строки (заголовки в первой строке)
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), order.ID)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), order.OrderNum)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), order.Creator)
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), order.Customer)
			f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), order.DopInfo)
			f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), order.MsNote)
		}

		// Устанавливаем заголовки HTTP-ответа для скачивания файла
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Set("Content-Disposition", "attachment; filename=orders.xlsx")
		w.WriteHeader(http.StatusOK)

		// Отправляем файл пользователю
		if err := f.Write(w); err != nil {
			log.Error("Failed to write Excel file", slog.String("error", err.Error()))
			http.Error(w, "Failed to generate Excel", http.StatusInternalServerError)
		}
	}
}
