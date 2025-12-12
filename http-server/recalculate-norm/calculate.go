package recalculate_norm

import (
	"context"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"vue-golang/internal/storage"
)

type NormProvider interface {
	CalculateNorm(ctx context.Context, orderID, pos int, typeIzd string) ([]storage.Operation, error)
}

func CalculateNormOperations(log *slog.Logger, calc NormProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.norm.CalculateNormOperations"

		orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Error("Invalid order_id", slog.String("error", err.Error()))
			http.Error(w, "Invalid order_id", http.StatusBadRequest)
			return
		}

		pos, err := strconv.Atoi(r.URL.Query().Get("position"))
		if err != nil {
			log.Error("Invalid pos", slog.String("error", err.Error()))
			http.Error(w, "Invalid pos", http.StatusBadRequest)
			return
		}

		typeIzd := r.URL.Query().Get("type")
		//if err != nil {
		//	log.Error("Invalid pos", slog.String("error", err.Error()))
		//	http.Error(w, "Invalid pos", http.StatusBadRequest)
		//	return
		//}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		norm, err := calc.CalculateNorm(ctx, orderID, pos, typeIzd)
		if err != nil {
			log.Error("Failed to calculate norm", slog.String("error", err.Error()))
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		log.Info("NOOORM", norm)

		render.JSON(w, r, norm)
	}
}
