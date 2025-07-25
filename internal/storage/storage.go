package storage

import "time"

type Order struct {
	ID       int    `json:"id"`
	OrderNum string `json:"order_num"`
	Creator  int    `json:"creator"`
	Customer string `json:"customer"`
	DopInfo  string `json:"dop_info"`
	MsNote   string `json:"ms_note"`
}

type OrderDemPrice struct {
	Position     int    `json:"position"`
	Creator      string `json:"creator"`
	NamePosition string `json:"name_position"`
	Count        int    `json:"count"`
	Image        string `json:"image"`
}

type ImportJSON struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsStudent bool   `json:"is_student"`
}

type ResultOrderDetails struct {
	Order         *Order           `json:"order"`
	OrderDemPrice []*OrderDemPrice `json:"order_dem_price"`
}

type OrderForm struct {
	OrderNum      string  `json:"order_num"`
	OperationName string  `json:"operation_name"`
	Kolvo         int     `json:"kolvo"`
	NormHour      float64 `json:"norm_hour"`
	NormMinutes   float64 `json:"norm_minutes"`
	FIO           string  `json:"fio"`
}

type FormPeo struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	FieldsPeo []FieldPeo `json:"fields_peo"`
}

type FieldPeo struct {
	Type          string   `json:"type"`
	Label         string   `json:"label"`
	Name          string   `json:"name"`
	Required      bool     `json:"required"`
	Value         float64  `json:"value,omitempty"`
	Count         int      `json:"count"`
	Minutes       float64  `json:"minutes"`
	RelatedFields []string `json:"related_fields"`
}

//type Materials struct {
//	Impost   string `json:"impost"`
//	Shtapick string `json:"shtapick"`
//}

type Workers struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	Profession string `json:"profession"`
}

type WorkersResult struct {
	ID            int       `json:"id"`
	OrderNum      string    `json:"order_num"`
	NameIzd       string    `json:"name_izd"`
	OperationName string    `json:"operation_name"`
	WorkerId      int       `json:"worker_id"`
	Value         float64   `json:"value"`
	Count         int       `json:"count"`
	AssignedAt    time.Time `json:"assigned_at"`
}

// 32 операции для глухарей
type DemResult struct {
	ID               int     `json:"id"`
	OrderNum         string  `json:"order_num"`
	Name             string  `json:"name"`
	Count            int     `json:"count"`
	PodgotovOboryd   float64 `json:"podgotov_oboryd"`
	NapilKontr       float64 `json:"napil_kontr"`
	NapilKrishek     float64 `json:"napil_krishek"`
	NapilImpost      float64 `json:"napil_impost"`
	Soedinitel       float64 `json:"soedinitel"`
	PromejSborka     float64 `json:"promej_sborka"`
	ImpostSverlovka  float64 `json:"impost_sverlovka"`
	ImpostFrezerovka float64 `json:"impost_frezerovka"`
	ImpostSborka     float64 `json:"impost_sborka"`
	OpresNastr       float64 `json:"opres_nastr"`
	Opresovka        float64 `json:"opresovka"`
	YstanYplotnitel  float64 `json:"ystan_yplotnitel"`
	Zashivka         float64 `json:"zashivka"`
	Profil           string  `json:"profil"`
	NapilStoikiDo3m  float64 `json:"napil_stoiki_do3m"`
	NapilStoikiBol3m float64 `json:"napil_stoiki_bol3m"`
	NapilRigelDo1m   float64 `json:"napil_rigel_do1m"`
	NapilRigelBol1m  float64 `json:"napil_rigel_bol1m"`
	SverlRigelZamok  float64 `json:"sverl_rigel_zamok"`
	YstanZamkov      float64 `json:"ystan_zamkov"`
	NapilShtapik     float64 `json:"napil_shtapik"`
	Ypakovka         float64 `json:"ypakovka"`
	FrezerRigel      float64 `json:"frezer_rigel"`
	ObrabotRam       float64 `json:"obrabot_ram"`
	HandsSborka      float64 `json:"hands_sborka"`
	FrezerNastr      float64 `json:"frezer_nastr"`
	Shtiftovka       float64 `json:"shtiftovka"`
	YstanovkaZapoln  float64 `json:"ystanovka_zapoln"`
}

type Operation struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type OrderData struct {
	ID         int         `json:"id"`
	OrderNum   string      `json:"order_num"`
	Name       string      `json:"name"`
	Count      int         `json:"count"`
	Profil     string      `json:"profil,omitempty"`
	Operations []Operation `json:"operations"`
}
