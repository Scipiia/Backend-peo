package postorder

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"vue-golang/internal/storage"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     int    `json:"id"`
	Data   map[string]interface{}
}

type SendMessages interface {
	SaveGlyhari(result storage.DemResultGlyhari) (int, error)
	SaveWorker(resWorker storage.WorkersResult) error
	SaveWindows(result storage.DemResultWindow) (int, error)
}

//type RequestWorkers struct {
//
//}

// TODO данные для занесения работяг в базу
func RequesWorkers(log *slog.Logger, workers SendMessages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.postorder.send.RequestWorkers"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req []storage.WorkersResult

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Ошибка реквеста", err)
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})
			return
		}

		//fmt.Printf("ФФФФФФФФФФФФФФФФФФ: %s\n", req)

		for _, item := range req {
			item.AssignedAt = time.Now() // Если дата не пришла, установите текущую
			err := workers.SaveWorker(item)
			if err != nil {
				log.Error("Ошибка при сохранении", slog.Any("error", err))
				render.JSON(w, r, Response{Error: "Ошибка базы данных"})
				return
			}
		}

		//log.Info("message added", slog.Int("id", worker))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
		})
	}
}

// TODO получаем данные для занесения в таблицу глухарей
func SaveNormOrderGlyhari(log *slog.Logger, message SendMessages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.message.SaveNormOrderGlyhari"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req storage.DemResultGlyhari

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Ошибка реквеста")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})
			return
		}

		log.Info("YYYYYYYY !!!!!!!!!!!!11111")
		saveOrder, err := message.SaveGlyhari(req)
		if err != nil {
			log.Info("Ошибка реквеста сообщения при вставке в базу данных сука блять уебище тупорылое", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		fmt.Printf("Form Data: %+v\n", req)
		//fmt.Printf("Form Data: %+v\n", req.FormId)

		log.Info("message added", slog.Int("id", saveOrder))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			ID:     saveOrder,
			Data: map[string]interface{}{
				"order_num": req.OrderNum,
				"ID":        saveOrder,
			},
		})
	}
}

// TODO получаем данные для занесения в таблицу окон
func SaveNormOrderWindow(log *slog.Logger, message SendMessages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.message.SaveNormOrderWindow"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req storage.DemResultWindow

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Ошибка реквеста")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})
			return
		}
		fmt.Printf("Form Data: %+v\n", req)

		log.Info("YYYYYYYY !!!!!!!!!!!!11111", req)
		saveOrder, err := message.SaveWindows(req)
		if err != nil {
			log.Info("Ошибка реквеста сообщения при вставке в базу данных сука блять уебище тупорылое окна ебаные", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		//fmt.Printf("Form Data: %+v\n", req)
		//fmt.Printf("Form Data: %+v\n", req.FormId)

		log.Info("message added", slog.Int("id", saveOrder))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			ID:     saveOrder,
			Data: map[string]interface{}{
				"order_num": req.OrderNum,
				"ID":        saveOrder,
			},
		})
	}
}
