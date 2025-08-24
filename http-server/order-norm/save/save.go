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
	SaveNormOrder(result storage.OrderDetails) (int64, error)
	SaveNormOperation(OrderID int64, operations []storage.NormOperation) error
}

type Response struct {
	OrderID int64 `json:"order_id"`
	Status  string
	Error   string
}

func SaveNormOrderOperation(log *slog.Logger, res ResultNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.normirovka.SaveNormOrderOperation"

		//var req RequestNormData
		var req storage.OrderDetails
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error("Неверный JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Неверные данные", http.StatusBadRequest)
			return
		}

		// --- 🔍 ПРОВЕРКА: операции со значением 0 ---
		//var zeroOps []string
		//for _, op := range req.Operations {
		//	// Если value == 0, добавляем в список
		//	if op.Value == 0 {
		//		zeroOps = append(zeroOps, op.Name)
		//	}
		//}
		//
		//if len(zeroOps) > 0 {
		//	log.Warn("Попытка сохранить операции со значением 0",
		//		slog.String("op", op),
		//		slog.Any("zero_ops", zeroOps),
		//		slog.String("order_num", req.OrderNum),
		//	)
		//
		//	// Ответ с понятной ошибкой
		//	render.JSON(w, r, Response{
		//		Error: "Обнаружены операции со значением 0: " + strings.Join(zeroOps, ", "),
		//	})
		//	return
		//}
		// --- ✅ Проверка пройдена ---

		// Сохраняем в БД
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
