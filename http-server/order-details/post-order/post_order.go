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
}

type SendMessages interface {
	SaveOrder(result storage.DemResult) (int, error)
	SaveWorker(resWorker storage.WorkersResult) error
}

//type RequestWorkers struct {
//
//}

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

		fmt.Printf("ФФФФФФФФФФФФФФФФФФ: %s\n", req)

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

func New(log *slog.Logger, message SendMessages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.message.send.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req storage.DemResult

		fmt.Println("EEEEEEEEEEWWWW", req)

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("Ошибка реквеста")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Empty request",
			})

			return
		}

		// Выводим полученные данные в консоль (для отладки)
		fmt.Printf("Order Number: %s\n", req.OrderNum)
		fmt.Printf("Count: %d\n", req.Count)
		fmt.Printf("Name: %s\n", req.Name)
		fmt.Printf("NapilOBORRTERF: %s\n", req.PodgotovOboryd)
		fmt.Printf("Napil: %s\n", req.NapilImpost)
		fmt.Printf("CountNapil: %s\n", req)
		//fmt.Printf("Form Data: %+v\n", req.FormData)

		saveOrder, err := message.SaveOrder(req)
		if err != nil {
			log.Info("Ошибка реквеста сообщения при вставке в базу данных сука блять уебище тупорылое", err)
			render.JSON(w, r, Response{Error: "da bleeeat"})
			return
		}

		log.Info("message added", slog.Int("id", saveOrder))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			ID:     saveOrder,
		})
	}
}
