package get

import (
	"context"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"vue-golang/internal/storage"
)

type MaterialProvider interface {
	GetOrderMaterials(ctx context.Context, id, pos int) ([]*storage.KlaesMaterials, error)
}

func GetMaterials(log *slog.Logger, material MaterialProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.materials.GetMaterials"

		idStr := r.URL.Query().Get("id")
		positionStr := r.URL.Query().Get("position")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("Invalid id", slog.String("error", err.Error()))
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		position, err := strconv.Atoi(positionStr)
		if err != nil {
			log.Error("Invalid position", slog.String("error", err.Error()))
			http.Error(w, "Invalid position", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		materials, err := material.GetOrderMaterials(ctx, id, position)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении заказов из дема")
			render.JSON(w, r, "В базе не найдено материалов")
			return
		}

		render.JSON(w, r, materials)
	}
}
