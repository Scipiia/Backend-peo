package save

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type ResultWorkers interface {
	SaveOperationWorkers(req storage.SaveWorkers) error
}

// handlers/executor/save.go
// handlers/executor/handlers.go

func SaveWorkersOperation(log *slog.Logger, result ResultWorkers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.executor.SaveWorkersOperation"

		var req storage.SaveWorkers
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Invalid JSON", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Bad request: invalid JSON", http.StatusBadRequest)
			return
		}

		// 🔹 Проверка: не пусто ли
		if len(req.Assignments) == 0 {
			log.Warn("Empty assignments list", slog.String("op", op))
			http.Error(w, "No assignments provided", http.StatusBadRequest)
			return
		}

		// 🔹 Проверка каждого элемента
		for i, a := range req.Assignments {
			if a.ProductID == 0 {
				log.Error("Missing product_id", slog.Int("index", i), slog.Any("assignment", a))
				http.Error(w, fmt.Sprintf("Assignment %d: product_id is required", i), http.StatusBadRequest)
				return
			}
			if a.EmployeeID == 0 {
				log.Error("Missing employee_id", slog.Int("index", i), slog.Any("assignment", a))
				http.Error(w, fmt.Sprintf("Assignment %d: employee_id is required", i), http.StatusBadRequest)
				return
			}
			if a.OperationName == "" {
				log.Error("Missing operation_name", slog.Int("index", i), slog.Any("assignment", a))
				http.Error(w, fmt.Sprintf("Assignment %d: operation_name is required", i), http.StatusBadRequest)
				return
			}
		}

		log.Info("Received assignments",
			slog.Int("total", len(req.Assignments)),
			slog.Any("sample", req.Assignments[0]),
		)

		// 🔹 Передаём в storage
		err := result.SaveOperationWorkers(req)
		if err != nil {
			log.Error("Failed to save assignments", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		log.Info("Assignments saved successfully",
			slog.Int("saved_count", len(req.Assignments)),
		)

		// 🔹 Ответ
		render.JSON(w, r, map[string]interface{}{
			"status":  "success",
			"saved":   len(req.Assignments),
			"details": req.Assignments, // опционально — для отладки
		})
	}
}
