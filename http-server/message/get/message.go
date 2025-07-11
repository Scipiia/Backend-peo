package get

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
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

type GetUrls interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, gethyi GetUrls) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.url.message.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		//var resp Response

		alias := "kusokgovna"

		url, err := gethyi.GetURL(alias)
		if err != nil {
			log.Info("Pizda vsemy", err)
			render.JSON(w, r, Response{
				Status: "Pizda ne polychilos",
				Error:  "polnii error",
			})
			return
		}

		render.JSON(w, r, Response{
			Status: "OKEY",
			Error:  "",
			Alias:  url,
		})

		log.Info("PIZDA", Response{})
	}
}
