package storage

import "time"

type OrderNormDetails struct {
	OrderNum     string          `json:"order_num"`
	TemplateCode string          `json:"template_code"`
	Name         string          `json:"name"`
	Count        float64         `json:"count"`
	TotalTime    float64         `json:"total_time"`
	Operations   []NormOperation `json:"operations"`
	Type         string          `json:"type"`
}

type NormOperation struct {
	Name    string  `json:"operation_name"`
	Label   string  `json:"operation_label"`
	Count   float64 `json:"count"`
	Value   float64 `json:"value"`
	Minutes float64 `json:"minutes"`
}

type GetOrderDetails struct {
	ID         int64           `json:"id"`
	OrderNum   string          `json:"order_num"`
	Name       string          `json:"name"`
	Count      float64         `json:"count"`
	TotalTime  float64         `json:"total_time"`
	CreatedAT  time.Time       `json:"created_at"`
	UpdatedAT  time.Time       `json:"updated_at"`
	Operations []NormOperation `json:"operations"`
	Type       string          `json:"type"`
}

// types/executor.go
type OperationExecutor struct {
	//ProductID     int64  `json:"product_id"`
	OperationName string  `json:"operation_name"`
	EmployeeID    int64   `json:"employee_id"`
	ActualMinutes float64 `json:"actual_minutes"`
	Notes         string  `json:"notes,omitempty"`
	ActualValue   float64 `json:"actual_value"`
}

type SaveExecutorsRequest struct {
	ProductID int64               `json:"product_id"`
	Executors []OperationExecutor `json:"executors"`
}

type GetWorkers struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
