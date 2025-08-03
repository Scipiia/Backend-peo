package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"vue-golang/http-server/auth"
	getorder "vue-golang/http-server/order-details/get-order"
	"vue-golang/http-server/order-details/orders"
	"vue-golang/http-server/order-details/post-order"
	"vue-golang/internal/config"
	"vue-golang/internal/storage/mysql"
)

func routes(cfg config.Config, log *slog.Logger, storage *mysql.Storage) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	//ip пользователя
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(middleware.Logger)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Разрешаем запросы с фронтенда
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsHandler.Handler)

	//TODO массив со всеми заказами из дема
	//ordersMiddleware := orders.OrdersMonthMiddleware(log, storage)
	router.Get("/api/orders", orders.New(log, storage))
	router.Handle("/api/generate-excel", orders.OrdersMonthMiddleware(log, storage)(orders.GenerateExel(log)))

	//TODO middleware
	// Middleware для получения данных о заказе
	// Маршруты для Гловяка где он внесет все данные по заказу
	orderDetailsMiddleware := getorder.OrderDetailsMiddleware(log, storage)
	// TODO JSON get
	router.With(orderDetailsMiddleware).Get("/api/orders/order/{id}", getorder.New(log))
	//TODO Генерация Excel-файла
	router.With(orderDetailsMiddleware).Get("/api/orders/order/generate-excel/{id}", getorder.GenerateExcel(log))
	// TODO получение формы для изделия
	router.Get("/api/orders/order/product/form", getorder.GetFormByID(log, storage))

	// TODO отправка формы после нормирования Гловяком с занесением в бд для изделия
	router.Post("/api/orders/order/product/gl", postorder.SaveNormOrderGlyhari(log, storage))
	router.Post("/api/orders/order/product/window", postorder.SaveNormOrderWindow(log, storage))
	router.Post("/api/orders/order/product/door", postorder.SaveNormOrderDoor(log, storage))

	//TODO получение работяг
	// будущие маршруты для мастеров в которых будут назначать работников
	router.Get("/api/master/orders", orders.ResultOrdersNorm(log, storage))
	router.Get("/api/orders/order/product/workers", getorder.GetWorkers(log, storage))
	//TODO добавление работяг с работой в базу
	router.Post("/api/orders/order/assignments", postorder.RequesWorkers(log, storage))

	//TODO получение нормированных деталей заказа
	router.Get("/api/master/orders/order/{id}", getorder.GetNormOrders(log, storage))

	//TODO AUTH
	router.Post("/api/login", auth.Auth(log))

	router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(log))
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value("role").(string)
			if role != "admin" {
				http.Error(w, "Access denied", http.StatusForbidden)
				return
			}
			w.Write([]byte("Admin access granted"))
		})
	})

	return router
}
