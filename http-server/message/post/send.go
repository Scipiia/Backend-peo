package post

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Alias  string `json:"alias"`
	Text   string `json:"text"`
	ID     int    `json:"id"`
}

type SendMessages interface {
	SaveMessage(message string) (int, error)
}

func New(log *slog.Logger, message SendMessages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.message.send.New"

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

		log.Info("AAAAAAAAAA", req.Text)
		log.Info("DDDDDDDD", r.Body)

		saveMessage, err := message.SaveMessage(req.Text)
		if err != nil {
			log.Info("Ошибка реквеста сообщения")
			render.JSON(w, r, Response{
				Error: "da bleeeat",
			})
			return
		}

		log.Info("message added", slog.Int("id", saveMessage))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			Alias:  req.Text,
			Text:   req.Text,
			ID:     saveMessage,
		})
	}
}
