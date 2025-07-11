package form_peo

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type Request struct {
	OrderNum string  `json:"order_num" validate:"required"`
	Num      int     `json:"num" validate:"required"`
	Sum      float64 `json:"sum" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Text   string `json:"text"`
	ID     int    `json:"id"`
}

type SendData interface {
	SaveOrderDetails(orderNum string, num int, sum float64) (int, error)
}

func New(log *slog.Logger, data SendData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.data-peo.send.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Ошибка реквеста")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", err)

			render.JSON(w, r, Response{Error: "invalid request"})

			return
		}

		log.Info("SDSDDS", req)
		details, err := data.SaveOrderDetails(req.OrderNum, req.Num, req.Sum)
		if err != nil {
			log.Info("Ошибка реквеста сообщения", err)
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})
			return
		}

		log.Info("message added", slog.Int("id", details))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			Text:   "ahyenchik",
			ID:     details,
		})

	}
}
