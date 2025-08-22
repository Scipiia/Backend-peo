package updatenormorder

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type OrderDetails interface {
	UpdateGlyhari(ID int, orderNum string, operations map[string]float64, additionalOps []AddOpRequest) error
}

type UpdateNormRequest struct {
	ID                   int                      `json:"id"`   // order_id
	Type                 string                   `json:"type"` // например, "glyhar", "window"
	OrderNum             string                   `json:"order_num"`
	Operations           []storage.OperationValue `json:"operations"` // список операций с новыми значениями
	AdditionalOperations []AddOpRequest           `json:"additional_operations,omitempty"`
}

type AddOpRequest struct {
	Operation string  `json:"operation"`
	Duration  float64 `json:"duration"`
	Comment   *string `json:"comment,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func UpdateNormOrder(log *slog.Logger, storage OrderDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateNormRequest
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("Failed to decode request", slog.String("error", err.Error()))
			render.JSON(w, r, Response{Error: "Invalid request"})
			return
		}

		slog.Info("REEEEEEQQQQQQQQQQQQQQQQQQQ", req)

		if req.ID == 0 {
			render.JSON(w, r, Response{Error: "Order ID is required"})
			return
		}

		var err error
		switch req.Type {
		case "glyhar":
			// Преобразуем []OperationValue → map[string]float64
			ops := make(map[string]float64)
			for _, op := range req.Operations {
				ops[op.ID] = op.Value
			}
			err = storage.UpdateGlyhari(req.ID, req.OrderNum, ops, req.AdditionalOperations)

		// case "window": ...
		// case "door": ...
		default:
			render.JSON(w, r, Response{Error: "Unsupported product type"})
			return
		}

		if err != nil {
			log.Error("Failed to update norm", slog.String("error", err.Error()))
			render.JSON(w, r, Response{Error: "Update failed"})
			return
		}

		render.JSON(w, r, Response{Status: "OK"})
	}
}
