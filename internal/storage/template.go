package storage

type Template struct {
	ID         int         `json:"ID"`
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Systema    *string     `json:"systema"`
	TypeIzd    *string     `json:"type_izd"`
	Profile    *string     `json:"profile"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Count    float64 `json:"count"`
	Label    string  `json:"label"`
	Value    float64 `json:"value"`
	Minutes  float64 `json:"minutes"`
	Required bool    `json:"required"`
}
