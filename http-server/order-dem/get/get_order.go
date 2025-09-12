package get

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"vue-golang/internal/storage"
)

type Request struct {
	ID int `json:"id"`
}

type ResponseOrder struct {
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
}

// Общая функция для получения данных о заказе
func getOrderDetails(log *slog.Logger, order OrderDetails, id int) (*storage.ResultOrderDetails, error) {
	const op = "handler.get-norm-details"

	log = log.With(
		slog.String("op", op),
	)

	details, err := order.GetOrderDetails(id)
	if err != nil {
		log.Info("Order not found", slog.Int("id", id), slog.String("error", err.Error()))
		return nil, fmt.Errorf("order-dem not found: %w", err)
	}

	log.Info("PIIIIIIZDA", details)

	return details, nil
}

func OrderDetailsMiddleware(log *slog.Logger, order OrderDetails) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем ID заказа из параметров запроса
			//idStr := r.URL.Query().Get("id")
			idStr := chi.URLParam(r, "id")
			if idStr == "" {
				log.Error("Missing order-dem-norm ID in query parameters")
				http.Error(w, "Missing order-dem-norm ID", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid order-dem-norm ID", http.StatusBadRequest)
				log.Error("Invalid order-dem-norm ID", slog.String("error", err.Error()))
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

// Обработчик GetOrderDetails
func GetOrderDetails(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		details, ok := r.Context().Value("orderDetails").(*storage.ResultOrderDetails)
		if !ok {
			http.Error(w, "Order details not found in context pizda", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, ResponseOrder{
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
