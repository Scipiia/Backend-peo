package get

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type Workers interface {
	GetWorkers() ([]storage.GetWorkers, error)
}

func GetWorkers(log *slog.Logger, worker Workers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-dem-norm.get.GetWorkers"

		workersNew, err := worker.GetWorkers()
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении работяг")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		log.With(slog.Int("found", len(workersNew))).Info("Заказы найдены")

		render.JSON(w, r, workersNew)
	}
}
