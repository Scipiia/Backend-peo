package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log/slog"
	getform "vue-golang/http-server/get-form"
	getorder "vue-golang/http-server/order-details/get-order"
	"vue-golang/http-server/order-details/orders"
	"vue-golang/http-server/order-details/post-order"
	updatenormorder "vue-golang/http-server/order-details/update-norm-order"
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
	//ordersMiddleware := orders.OrdersMonthMiddleware(log, storage)
	router.Get("/api/orders", orders.New(log, storage))
	router.Handle("/api/generate-excel", orders.OrdersMonthMiddleware(log, storage)(orders.GenerateExel(log)))

	//TODO middleware
	// Middleware для получения данных о заказе
	// Маршруты для Гловяка где он внесет все данные по заказу
	orderDetailsMiddleware := getorder.OrderDetailsMiddleware(log, storage)
	// TODO JSON get ОЩИБКА ТУТ
	router.With(orderDetailsMiddleware).Get("/api/orders/order/{id}", getorder.New(log))

	// TODO отправка формы после нормирования Гловяком с занесением в бд для изделия
	router.Post("/api/orders/order-norm/product/gl", postorder.SaveNormOrderGlyhari(log, storage))
	router.Post("/api/orders/order-norm/product/window", postorder.SaveNormOrderWindow(log, storage))
	router.Post("/api/orders/order-norm/product/door", postorder.SaveNormOrderDoor(log, storage))
	router.Post("/api/orders/order-norm/product/vitraj", postorder.SaveNormOrderVitraj(log, storage))
	router.Post("/api/orders/order-norm/product/loggia", postorder.SaveNormOrderLoggia(log, storage))

	//TODO отправка и внесение изменении обновления нарядов
	router.Put("/api/norm/orders/order-norm/edit/{id}", updatenormorder.UpdateNormOrder(log, storage))
	//TODO генерация excel для всех нормированных нарядов с фильтром
	router.Get("/api/norm/orders/excel", getorder.ExportNormOrders(log, storage))

	//TODO вытягивание всех нормированных нарядов Гловяком
	router.Get("/api/norm/orders", getorder.GetAllNormOrders(log, storage))

	//TODO новая логика с распределением операции YYYYYYYYYYYYYYYYY
	router.Get("/template", getform.GetTemplateByCode(log, storage))
	router.Get("/all_templates", getform.GetAllTemplates(log, storage))

	//TODO сохранение нормированных нарядов
	router.Post("/api/orders/order-norm/form", save.SaveNormOrderOperation(log, storage))

	//TODO get получение нормированного наряда
	router.Get("/api/orders/order/norm/{id}", get.GetNormOrder(log, storage))

	//TODO get получение всех нормированных нарядов
	router.Get("/api/orders/order/norm/all", get.GetNormOrders(log, storage))

	//TODO update обновление нормированного наряда
	router.Put("/api/orders/order/norm/update/{id}", update.UpdateNormOrderOperation(log, storage))

	//TODO назначение сотрудников
	router.Post("/api/workers", saveWorkers.SaveExecutorsHandler(log, storage))
	//TODO получение всех сотрудников
	router.Get("/api/workers/all", getWorkers.GetWorkers(log, storage))

	return router
}
