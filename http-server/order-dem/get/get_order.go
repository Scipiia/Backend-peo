package get

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"time"
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
	GetOrderDetails(ctx context.Context, id int) (*storage.ResultOrderDetails, error)
}

func GetOrderDetails(log *slog.Logger, order OrderDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.get_orders.GetOrderDetails1"

		idStr := chi.URLParam(r, "id")
		if idStr == "" {
			log.Error("Missing order-dem-norm ID in query parameters")
			http.Error(w, "Missing order-dem-norm ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("Invalid order-dem-norm ID", slog.String("error", err.Error()))
			http.Error(w, "Invalid order-dem-norm ID", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		details, err := order.GetOrderDetails(ctx, id)
		if err != nil {
			log.Error("не удалось получить детали заказа из дема", slog.String("error", err.Error()))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		render.JSON(w, r, ResponseOrder{
			ID:            details.Order.ID,
			OrderNum:      details.Order.OrderNum,
			Creator:       details.Order.Creator,
			Customer:      details.Order.Customer,
			DopInfo:       details.Order.DopInfo,
			MsNote:        details.Order.MsNote,
			OrderDemPrice: details.OrderDemPrice,
			Error:         "",
			Status:        strconv.Itoa(http.StatusOK),
		})
	}
}
