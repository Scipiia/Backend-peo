package getorder

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type AllNormOrders interface {
	GetAllProducts() ([]storage.ProductItem, error)
}
type ResponseNormOrders struct {
	Orders []storage.ProductItem `json:"orders"`
	Error  string                `json:"error"`
}

func GetAllNormOrders(log *slog.Logger, normOrders AllNormOrders) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		orders, err := normOrders.GetAllProducts()
		if err != nil {
			log.Info("Не удалось вытащить заказы из базы данных", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		log.Info("ORDERS", orders)
		// Формируем ответ
		render.JSON(w, r, ResponseNormOrders{
			Orders: orders,
			Error:  "",
		})
	}
}
