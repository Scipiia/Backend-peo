package storage

import "time"

type OrderFinalReport struct {
	OrderNum string        `json:"order_num"`
	Izdelie  []IzdelieInfo `json:"izdelie"`
}

type IzdelieInfo struct {
	ID           int64            `json:"id"`
	Name         string           `json:"name"`
	TemplateName string           `json:"template_name"`
	Operations   []OperationsNorm `json:"operations"`
}

type OperationsNorm struct {
	OperationName  string    `json:"operation_name"`
	OperationLabel string    `json:"operation_label"`
	NormMinutes    float64   `json:"norm_minutes"`
	Executors      []Workers `json:"executors"`
}

type Workers struct {
	WorkerName    string  `json:"worker_name"`
	ActualMinutes float64 `json:"actual_minutes"`
	ActualValue   float64 `json:"actual_value"`
}

type ReportFinalOrders struct {
	OrderNum     string    `json:"order_num"`
	IzdCount     int       `json:"izd_count"`
	FirstCreated time.Time `json:"first_created"`
}
