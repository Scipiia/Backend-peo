package get

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"vue-golang/internal/storage"
)

type ResultGetNorm interface {
	GetNormOrder(id int64) (*storage.GetOrderDetails, error)
	GetNormOrders(orderNum, orderType string) ([]storage.GetOrderDetails, error)
}

func GetNormOrder(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.GetNormOrder"

		// Извлекаем id из URL
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		log.Info("Получение нормировки", slog.Int64("id", id))

		norm, err := result.GetNormOrder(id)
		if err != nil {
			if strings.Contains(err.Error(), "не найдена") {
				log.Info("Ошибка реквеста сообщения при вставке в базу заказа сука блять уебище тупорылое DOOR ебаные", err)
				http.Error(w, "Нормировка не найдена", http.StatusNotFound)
				return
			}
			log.Error("Ошибка получения нормировки", slog.String("op", op), slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
			return
		}

		// Успешный ответ
		render.JSON(w, r, norm)
	}
}

func GetNormOrders(log *slog.Logger, result ResultGetNorm) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order-norm.get.GetNormOrders"

		// Получаем фильтр
		orderNum := r.URL.Query().Get("order_num")
		orderType := r.URL.Query().Get("type")

		log.With(
			slog.String("op", op),
			slog.String("order_num_filter", orderNum),
			slog.String("order_type_filter", orderType),
		).Info("Запрос на получение заказов")

		// Передаём фильтр (может быть пустым)
		items, err := result.GetNormOrders(orderNum, orderType)
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Ошибка при получении заказов")
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		log.With(slog.Int("found", len(items))).Info("Заказы найдены")

		// Возвращаем JSON
		render.JSON(w, r, items)
	}
}
