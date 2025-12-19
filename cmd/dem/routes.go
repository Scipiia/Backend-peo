package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"strings"
	gettemplate "vue-golang/http-server/get-template"
	getmaterials "vue-golang/http-server/materials/get"
	getorder "vue-golang/http-server/order-dem/get"
	"vue-golang/http-server/order-norm/get"
	"vue-golang/http-server/order-norm/save"
	"vue-golang/http-server/order-norm/update"
	recalculate_norm "vue-golang/http-server/recalculate-norm"
	getWorkers "vue-golang/http-server/workers/get"
	saveWorkers "vue-golang/http-server/workers/save"
	"vue-golang/internal/config"
	"vue-golang/internal/service"
	"vue-golang/internal/storage/mysql"
)

func routes(cfg config.Config, log *slog.Logger, storage *mysql.Storage, service *service.NormService) *chi.Mux {
	router := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Разрешаем запросы с фронтенда
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsHandler.Handler)

	router.Use(middleware.RequestID)
	//ip пользователя
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./dist")) // путь к билду Vue

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Если путь начинается с /api — уже должен был быть обработан выше
		// Но на всякий случай: если дошли сюда — это не API
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		// Отдаем index.html для Vue Router
		fs.ServeHTTP(w, r)
	})

	//TODO массив со всеми заказами из дема
	router.Get("/api/orders", getorder.GetOrdersFilter(log, storage))

	// Маршруты для Гловяка где он внесет все данные по заказу
	router.Get("/api/orders/order/{id}", getorder.GetOrderDetails(log, storage))

	//TODO получение шаблонов
	router.Get("/api/template", gettemplate.GetTemplatesByCode(log, storage))
	router.Get("/api/all_templates", gettemplate.GetAllTemplates(log, storage))

	//TODO сохранение нормированных нарядов
	router.Post("/api/orders/order-norm/template", save.SaveNormOrderOperation(log, storage))

	//TODO обновление статуса нормировки(отмена)
	router.Post("/api/orders/cancel", update.UpdateCancelStatus(log, storage))

	//TODO get получение нормированного наряда
	router.Get("/api/orders/order/norm/{id}", get.GetNormOrder(log, storage))
	//TODO получение нескольких заказов нормирования(связанных между собой)
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

	//Материалы к заказу
	router.Get("/api/materials", getmaterials.GetMaterials(log, storage))
	router.Get("/api/materials/ubub", recalculate_norm.CalculateNormOperations(log, service))

	return router
}
