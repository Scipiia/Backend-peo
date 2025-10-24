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

type FormsJSON interface {
	GetFormByCode(code string) (*storage.Form, error)
	GetAllForms() ([]*storage.Form, error)
}

type ResponseForm struct {
	ID         int                 `json:"ID"`
	Code       string              `json:"code"`
	Name       string              `json:"name"`
	Category   string              `json:"category"`
	Systema    *string             `json:"systema"`
	TypeIzd    *string             `json:"type_izd"`
	Profile    *string             `json:"profile"`
	Operations []storage.Operation `json:"operations"`
}

func GetFormByCode(log *slog.Logger, provider FormsJSON) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.template.GetFormByCode"

		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		).Info("Fetching template by code")

		code := r.URL.Query().Get("code")
		if code == "" {
			log.With(slog.String("op", op)).Error("Missing 'code' in query parameters")
			http.Error(w, "Missing required query parameter 'code'", http.StatusBadRequest)
			return
		}

		log.With(slog.String("code", code)).Info("Fetching template")

		// Получаем шаблон из хранилища
		template, err := provider.GetFormByCode(code)
		if err != nil {
			if strings.Contains(err.Error(), "не найден") || errors.Is(err, sql.ErrNoRows) {
				log.With(slog.String("op", op), slog.String("code", code)).Warn("Form not found")
				http.Error(w, "Form not found", http.StatusNotFound)
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
		response := ResponseForm{
			ID:         template.ID,
			Code:       template.Code,
			Name:       template.Name,
			Category:   template.Category,
			Systema:    template.Systema,
			TypeIzd:    template.TypeIzd,
			Profile:    template.Profile,
			Operations: template.Operations,
		}

		log.With(slog.String("code", code)).Info("Successfully fetched form")

		// Отправляем JSON
		log.Info("FOOORM1", response)
		render.JSON(w, r, response)
	}
}

type ResponseAllForm struct {
	Forms []*storage.Form
	Error string
}

func GetAllForms(log *slog.Logger, provider FormsJSON) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.template.GetAllForms"

		log.With(slog.String("op", op)).Info("Fetching all templates")

		templates, err := provider.GetAllForms()
		if err != nil {
			log.With(slog.String("op", op), slog.String("error", err.Error())).Error("Failed to fetch templates")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		response := ResponseAllForm{
			Forms: templates,
			Error: "",
		}

		render.JSON(w, r, response)
	}
}
