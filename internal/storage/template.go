package storage

// models/get_template.go
type Form struct {
	ID         int         `json:"ID"`
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Count    float64 `json:"count"`
	Label    string  `json:"label"`
	Value    float64 `json:"value"` // в часах
	Minutes  float64 `json:"minutes"`
	Required bool    `json:"required"`
}
