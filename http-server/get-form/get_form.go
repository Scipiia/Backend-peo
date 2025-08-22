package getform

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strings"
	"vue-golang/internal/storage"
)

// pkg/http-server/handlers/template/get.go

type TemplateProvider interface {
	GetTemplateByCode(code string) (*storage.Template, error)
	GetAllTemplates() ([]*storage.Template, error)
}

type ResponseTemplate struct {
	ID         int                  `json:"ID"`
	Code       string               `json:"code"`
	Name       string               `json:"name"`
	Category   string               `json:"category"`
	Operations []storage.Operation1 `json:"operations"`
}

func GetTemplateByCode(log *slog.Logger, provider TemplateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.template.GetTemplateByCode"

		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		).Info("Fetching template by code")

		// Получаем code из query: ?code=GLUHOE_OKNO_BASIC
		code := r.URL.Query().Get("code")
		if code == "" {
			log.With(slog.String("op", op)).Error("Missing 'code' in query parameters")
			http.Error(w, "Missing required query parameter 'code'", http.StatusBadRequest)
			return
		}

		log.With(slog.String("code", code)).Info("Fetching template")

		// Получаем шаблон из хранилища
		template, err := provider.GetTemplateByCode(code)
		if err != nil {
			if strings.Contains(err.Error(), "не найден") || errors.Is(err, sql.ErrNoRows) {
				log.With(slog.String("op", op), slog.String("code", code)).Warn("Template not found")
				http.Error(w, "Template not found", http.StatusNotFound)
				return
			}

			log.With(
				slog.String("op", op),
				slog.String("code", code),
				slog.String("error", err.Error()),
			).Error("Failed to fetch template")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Формируем ответ
		response := ResponseTemplate{
			ID:         template.ID,
			Code:       template.Code,
			Name:       template.Name,
			Category:   template.Category,
			Operations: template.Operations,
		}

		log.With(slog.String("code", code)).Info("Successfully fetched template")

		// Отправляем JSON
		render.JSON(w, r, response)
	}
}

type ResponseAllTemplates struct {
	Template []*storage.Template
	Error    string
}

func GetAllTemplates(log *slog.Logger, provider TemplateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.template.GetAllTemplates"

		log.With(slog.String("op", op)).Info("Fetching all templates")

		templates, err := provider.GetAllTemplates()
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Failed to fetch templates")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		response := ResponseAllTemplates{
			Template: templates,
			Error:    "",
		}

		render.JSON(w, r, response)
	}
}
