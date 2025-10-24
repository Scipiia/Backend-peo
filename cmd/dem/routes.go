package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log/slog"
	getform "vue-golang/http-server/get-form"
	getorder "vue-golang/http-server/order-dem/get"
	"vue-golang/http-server/order-norm/get"
	"vue-golang/http-server/order-norm/save"
	"vue-golang/http-server/order-norm/update"
	getWorkers "vue-golang/http-server/workers/get"
	saveWorkers "vue-golang/http-server/workers/save"
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
	router.Get("/api/orders", getorder.GetOrdersFilter(log, storage))

	//TODO middleware
	// Middleware для получения данных о заказе
	// Маршруты для Гловяка где он внесет все данные по заказу
	orderDetailsMiddleware := getorder.OrderDetailsMiddleware(log, storage)
	router.With(orderDetailsMiddleware).Get("/api/orders/order/{id}", getorder.GetOrderDetails(log))

	//TODO новая логика с распределением операции YYYYYYYYYYYYYYYYY
	router.Get("/template", getform.GetFormByCode(log, storage))
	router.Get("/all_templates", getform.GetAllForms(log, storage))

	//TODO сохранение нормированных нарядов
	router.Post("/api/orders/order-norm/form", save.SaveNormOrderOperation(log, storage))

	//TODO get получение нормированного наряда
	router.Get("/api/orders/order/norm/{id}", get.GetNormOrder(log, storage))
	//TODO получение нескольких заказов нормирования(связанных между собой) НННННННННННННННННН
	router.Get("/api/orders/order-norm/by-order", get.GetNormOrdersOrderNum(log, storage))
	router.Get("/api/orders/order-norm/{id}", get.DoubleReportOrder(log, storage))

	//TODO get получение всех нормированных нарядов
	router.Get("/api/orders/order/norm/all", get.GetNormOrders(log, storage))

	//TODO update обновление нормированного наряда
	router.Put("/api/orders/order/norm/update/{id}", update.UpdateNormOrderOperation(log, storage))

	//TODO назначение сотрудников
	router.Post("/api/workers", saveWorkers.SaveWorkersOperation(log, storage))
	//TODO получение всех сотрудников
	router.Get("/api/workers/all", getWorkers.GetWorkers(log, storage))

	//TODO финальные маршруты для всех готовых заказов и возможность провалиться в них
	router.Get("/api/allians/{order_num}", get.FinalReportNormOrder(log, storage))
	router.Get("/api/all_final_order", get.FinalReportNormOrders(log, storage))

	//TODO финальное обновление
	router.Put("/api/final/update/{id}", update.UpdateFinalOrder(log, storage))

	return router
}
