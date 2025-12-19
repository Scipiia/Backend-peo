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
	Rules      []Rule      `json:"rules"`
}

type Operation struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Count    float64 `json:"count"`
	Label    string  `json:"label"`
	Value    float64 `json:"value"`
	Minutes  float64 `json:"minutes"`
	Required bool    `json:"required"`
	Group    string  `json:"group"`
}

type Rule struct {
	Operation      string                 `json:"operation"`
	Condition      map[string]interface{} `json:"condition"`
	Mode           string                 `json:"mode"`    // "set", "multiplied", "additive"
	SetValue       float64                `json:"value"`   // fallback
	SetMinutes     float64                `json:"minutes"` // fallback
	ValuePerUnit   float64                `json:"value_per_unit"`
	MinutesPerUnit float64                `json:"minutes_per_unit"`
	// Set — можно удалить, если не используется
}
