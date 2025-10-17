package update

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"vue-golang/internal/storage"
)

type ResultUpdateNorm interface {
	UpdateNormOrder(ID int64, update storage.UpdateOrderDetails) error
	UpdateFinalOrder(ID int64, update storage.UpdateFinalOrderDetails) error
}

func UpdateNormOrderOperation(log *slog.Logger, update ResultUpdateNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.norm.UpdateNormHandler"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var req storage.UpdateOrderDetails
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Invalid JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		log.Info("TTTOOOOOTALTIIIIME", req.TotalTime)

		log.Info("Обновление нормировки", slog.Int64("id", id))

		err = update.UpdateNormOrder(id, req)
		if err != nil {
			log.Error("Ошибка обновления", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Save error", http.StatusInternalServerError)
			return
		}

		log.Info("Нормировка обновлена", slog.Int64("id", id))

		render.JSON(w, r, map[string]interface{}{
			"status":  "success",
			"norm_id": id,
		})
	}
}

func UpdateFinalOrder(log *slog.Logger, update ResultUpdateNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.norm.UpdateFinalOrder"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		fmt.Println("IDDDD", id)

		var req storage.UpdateFinalOrderDetails
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Invalid JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		fmt.Println("REQQQQ", req)

		log.Info("Обновление final нормировки", slog.Int64("id", id))

		err = update.UpdateFinalOrder(id, req)
		if err != nil {
			log.Error("Ошибка обновления", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Save error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"status": "success",
		})
	}
}
