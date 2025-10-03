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

// TODO new logic

// PEOProduct — изделие с агрегированными данными по сотрудникам
type PEOProduct struct {
	ID              int64     `json:"id"`
	OrderNum        string    `json:"order_num"`
	Customer        string    `json:"customer"`
	TotalTime       float64   `json:"total_time"` // "площадь"
	CreatedAt       time.Time `json:"created_at"`
	Status          string    `json:"status"`
	PartType        string    `json:"part_type"`
	Type            string    `json:"type"`
	ParentProductID *int64    `json:"parent_product_id"`
	ParentAssembly  string    `json:"parent_assembly"`
	CustomerType    string    `json:"customer_type"`
	Systema         string    `json:"systema"`
	TypeIzd         string    `json:"type_izd"`
	Profile         string    `json:"profile"`
	Count           int       `json:"count"`
	Sqr             float64   `json:"sqr"`

	// Дополнительные поля, которые ты добавишь позже
	// Material       *string `json:"material"`
	// ProfileType    *string `json:"profile_type"`

	// Мапа: employee_id → суммарные минуты
	EmployeeMinutes map[int64]float64 `json:"employee_minutes"`
}

// PEOEmployee — сотрудник для заголовков таблицы
type PEOEmployee struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
