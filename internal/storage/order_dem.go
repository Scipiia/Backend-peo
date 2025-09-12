package storage

type Order struct {
	ID       int    `json:"id"`
	OrderNum string `json:"order_num"`
	Creator  int    `json:"creator"`
	Customer string `json:"customer"`
	DopInfo  string `json:"dop_info"`
	MsNote   string `json:"ms_note"`
}

type OrderDemPrice struct {
	Position     int      `json:"position"`
	Creator      string   `json:"creator"`
	NamePosition string   `json:"name_position"`
	Count        *float64 `json:"count"`
	Image        *string  `json:"image"`
	Sqr          float64  `json:"sqr"`
}

type ResultOrderDetails struct {
	Order         *Order           `json:"order_dem_norm"`
	OrderDemPrice []*OrderDemPrice `json:"order_dem_price"`
}
