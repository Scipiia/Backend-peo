package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type Request struct {
	URL   string `json:"urls"`
	Alias string `json:"alias"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Alias  string `json:"alias"`
}

//type Template interface {
//	Render(w http.ResponseWriter, r *http.Request, name string, td *templates.TemplData)
//}

type URLSaver interface {
	SaveURL(url, alias string) (int, error)
	//NewTemplateCache(dir string) (map[string]*template.Template, error)
	//RenderHTML(w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.url.save.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом
			// Обработаем её отдельно
			log.Error("request body is empty")

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "empty request",
			})

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", err)

			render.JSON(w, r, Response{Error: "invalid request"})

			return
		}

		//alias := "kusokgovna"

		id, err := urlSaver.SaveURL(req.URL, req.Alias)
		if err != nil {
			log.Info("Pizda vsemy", err)
			render.JSON(w, r, Response{Error: "url ne dobavilsia bleat suka"})
			return
		}

		log.Info("url added", slog.Int("id", id))

		render.JSON(w, r, Response{
			Status: strconv.Itoa(http.StatusOK),
			Error:  "",
			Alias:  req.Alias,
		})

		//template.Render(w, r, "home.page.tmpl", nil)

	}
}
