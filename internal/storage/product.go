package storage

import "time"

// УДАЛИТЬ НЕ ПРИГОДИТСЯ
// models/product.go
type ProductInstance struct {
	ID           int64                 `json:"id"`
	OrderNum     string                `json:"order_num"`
	TemplateCode string                `json:"template_code"`
	Name         string                `json:"name"`
	Count        int                   `json:"count"`
	TotalTime    float64               `json:"total_time"`
	Operations   []OperationValueInput `json:"operations"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

type OperationValueInput struct {
	OperationName string  `json:"operation_name"`
	ActualValue   float64 `json:"actual_value"`
}
