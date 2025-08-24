package storage

type OrderDetails struct {
	OrderNum     string          `json:"order_num"`
	TemplateCode string          `json:"template_code"`
	Name         string          `json:"name"`
	Count        int             `json:"count"`
	TotalTime    float64         `json:"total_time"`
	Operations   []NormOperation `json:"operations"`
}

type NormOperation struct {
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	Count   int     `json:"count"`
	Value   float64 `json:"value"`
	Minutes float64 `json:"minutes"`
}
