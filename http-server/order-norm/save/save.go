package save

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"vue-golang/internal/storage"
)

type ResultNorm interface {
	SaveNormOrder(result storage.OrderNormDetails) (int64, error)
	SaveNormOperation(OrderID int64, operations []storage.NormOperation) error
}

type Response struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

func SaveNormOrderOperation(log *slog.Logger, res ResultNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.normirovka.SaveNormOrderOperation"

		//var req RequestNormData
		var req storage.OrderNormDetails
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error("Неверный JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Неверные данные", http.StatusBadRequest)
			return
		}

		orderID, err := res.SaveNormOrder(req)
		if err != nil {
			log.Info("Ошибка реквеста сообщения при вставке в базу заказа сука блять уебище тупорылое DOOR ебаные", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		log.Info("RREEQWQWWQ", req.Operations)
		// Сохраняем операции
		err = res.SaveNormOperation(orderID, req.Operations)
		if err != nil {
			log.Info("Ошибка реквеста сообщения при вставке в базу операции сука блять уебище тупорылое ROOOT ебаные", err)
			render.JSON(w, r, Response{Error: "da bleeeat1"})
			return
		}

		log.Info("message added", slog.Int64("id", orderID))

		render.JSON(w, r, Response{
			OrderID: orderID,
			Status:  strconv.Itoa(http.StatusOK),
			Error:   "",
		})
	}
}
