package save

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type ResultWorkers interface {
	SaveOperationExecutors(req storage.SaveExecutorsRequest) error
}

// handlers/executor/save.go
func SaveExecutorsHandler(log *slog.Logger, result ResultWorkers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.executor.SaveExecutorsHandler"

		var req storage.SaveExecutorsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Invalid JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.ProductID == 0 {
			http.Error(w, "product_id is required", http.StatusBadRequest)
			return
		}

		if len(req.Executors) == 0 {
			http.Error(w, "executors list is empty", http.StatusBadRequest)
			return
		}

		log.Info("Назначение исполнителей",
			slog.Int64("product_id", req.ProductID),
			slog.Int("executors_count", len(req.Executors)),
		)

		err := result.SaveOperationExecutors(req)
		if err != nil {
			log.Error("Ошибка сохранения исполнителей", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		log.Info("Исполнители сохранены", slog.Int64("product_id", req.ProductID))

		render.JSON(w, r, map[string]interface{}{
			"status":     "success",
			"product_id": req.ProductID,
			"assigned":   len(req.Executors),
		})
	}
}
