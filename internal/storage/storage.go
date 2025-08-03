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
type DemResultGlyhari struct {
	ID               int     `json:"id"`
	OrderNum         string  `json:"order_num"`
	Name             string  `json:"name"`
	Count            int     `json:"count"`
	NastrNapil       float64 `json:"nast_napil"`
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
	NapilDonnik      float64 `json:"napil_donnik"`
	AdapterNapil     float64 `json:"adapter_napil"`
	AdapterYstan     float64 `json:"adapter_ystan"`
	YstanYplotnFalc  float64 `json:"ystan_yplotn_falc"`
	OrderId          int     `json:"order_id"`
	TotalTime        float64 `json:"total_time"`
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

type DemResultWindow struct {
	ID                  int     `json:"id"`
	OrderNum            string  `json:"order_num"`
	Name                string  `json:"name"`
	Count               int     `json:"count"`
	Profil              string  `json:"profil"`
	PodgotOboryd        float64 `json:"podgot_oboryd"`
	NapilRamStv         float64 `json:"napil_ram_stv"`
	NapilNramStv        float64 `json:"napil_nram_stv"`
	NapilImposta        float64 `json:"napil_imposta"`
	SverlovkaStoek      float64 `json:"sverlovka_stoek"`
	ObrabotRam          float64 `json:"obrabot_ram"`
	ObrabotkaStv        float64 `json:"obrabotka_stv"`
	PromejSborkaStv     float64 `json:"promej_sborka_stv"`
	NapilAdaptera       float64 `json:"napil_adaptera"`
	PromejSborkaRam     float64 `json:"promej_sborka_ram"`
	PromejSborkaGl      float64 `json:"promej_sborka_gl"`
	OpresRam            float64 `json:"opres_ram"`
	OpresGl             float64 `json:"opres_gl"`
	FrezerStv           float64 `json:"frezer_stv"`
	YstanYplRam         float64 `json:"ystan_ypl_ram"`
	YstanYplStv         float64 `json:"ystan_ypl_stv"`
	NapilTag            float64 `json:"napil_tag"`
	SverloTag           float64 `json:"sverlo_tag"`
	YstanFurn           float64 `json:"ystan_furn"`
	NaveshStv           float64 `json:"navesh_stv"`
	Zashivka            float64 `json:"zashivka"`
	ObrabDopProfil      float64 `json:"obrab_dop_profil"`
	YstanAdaptera       float64 `json:"ystan_adaptera"`
	FrezerImpostPilon   float64 `json:"frezer_impost_pilon"`
	KreplYglRam         float64 `json:"krepl_ygl_ram"`
	GlyharDorab         float64 `json:"glyhar_dorab"`
	YplotGlyhar         float64 `json:"yplot_glyhar"`
	ZashivGlyhar        float64 `json:"zashiv_glyhar"`
	OpresStv            float64 `json:"opres_stv"`
	RazborErkera        float64 `json:"razbor_erkera"`
	GlWindow            float64 `json:"gl_window"`
	ObrabotShtylp       float64 `json:"obrabot_shtylp"`
	FrezerPodShtylp     float64 `json:"frezer_pod_shtylp"`
	YstanShtylp         float64 `json:"ystan_shtylp"`
	YstanImpost         float64 `json:"ystan_impost"`
	NastrForOpres       float64 `json:"nastr_for_opres"`
	NapilRam            float64 `json:"napil_ram"`
	NapilNram           float64 `json:"napil_nram"`
	NapilDopProfil      float64 `json:"napil_dop_profil"`
	FrezerDopProfil     float64 `json:"frezer_dop_profil"`
	ObrabotkaStvRychka  float64 `json:"obrabotka_stv_rychka"`
	ObrabotkaStvZamok   float64 `json:"obrabotka_stv_zamok"`
	NapilShtylp         float64 `json:"napil_shtylp"`
	ImpostSbor          float64 `json:"impost_sbor"`
	RezinaRam           float64 `json:"rezina_ram"`
	PodgotovFurn        float64 `json:"podgotov_furn"`
	PodgotovRam         float64 `json:"podgotov_ram"`
	Razborka            float64 `json:"razborka"`
	NapilZashitProf     float64 `json:"napil_zashit_prof"`
	ImpostFrezer        float64 `json:"impost_frezer"`
	YstanSteklaZaliv    float64 `json:"ystan_stekla_zaliv"`
	YstanKrishek        float64 `json:"ystan_krishek"`
	FrezerProfilZamazka float64 `json:"frezer_profil_zamazka"`
	NapilStoikiDo3m     float64 `json:"napil_stoiki_do3m"`
	NapilStoikiBol3m    float64 `json:"napil_stoiki_bol3m"`
	NapilRigelDo1m      float64 `json:"napil_rigel_do1m"`
	NapilRigelBol1m     float64 `json:"napil_rigel_bol1m"`
	SverloRigelZamok    float64 `json:"sverlo_rigel_zamok"`
	YstanZamok          float64 `json:"ystan_zamok"`
	Shtiftovka          float64 `json:"shtiftovka"`
	FrezerRigel         float64 `json:"frezer_rigel"`
	PartSborka          float64 `json:"part_sborka"`
	YstanRezYgl         float64 `json:"ystan_rez_ygl"`
	OpresYgl            float64 `json:"opres_ygl"`
	YstanTermovst       float64 `json:"ystan_termovst"`
	NapilShtapik        float64 `json:"napil_shtapik"`
	YpakSekcii          float64 `json:"ypak_sekcii"`
	YpakRigel           float64 `json:"ypak_rigel"`
	YpakYplotn          float64 `json:"ypak_yplotn"`
	YpakShtapik         float64 `json:"ypak_shtapik"`
	NapilDonnika        float64 `json:"napil_donnika"`
	NastrPbx            float64 `json:"nastr_pbx"`
	MehObrabPzr         float64 `json:"meh_obrab_pzr"`
	RabotaPbx           float64 `json:"rabota_pbx"`
	SlesarObrFurn       float64 `json:"slesar_obr_furn"`
	ImpostSverlo        float64 `json:"impost_sverlo"`
	Opresovka           float64 `json:"opresovka"`
	SborRychka          float64 `json:"sbor_rychka"`
	SborPetli           float64 `json:"sbor_petli"`
	PlastikYstnRam      float64 `json:"plastik_ystn_ram"`
	NapilStv            float64 `json:"napil_stv"`
	YstanFurnStv        float64 `json:"ystan_furn_stv"`
	YstanFurnRam        float64 `json:"ystan_furn_ram"`
	ComplOtg            float64 `json:"compl_otg"`
	YpakIzd             float64 `json:"ypak_izd"`
	Otgryz              float64 `json:"otgryz"`
	RaspFurn            float64 `json:"rasp_furn"`
	OrderId             int     `json:"order_id"`
	TotalTime           float64 `json:"total_time"`
}

type DemResultDoor struct {
	ID                int     `json:"id"`
	OrderNum          string  `json:"order_num"`
	Name              string  `json:"name"`
	Count             int     `json:"count"`
	Profil            string  `json:"profil"`
	NapilRamStv       float64 `json:"napil_ram_stv"`
	PetliObr          float64 `json:"petli_obr"`
	ZamokObr          float64 `json:"zamok_obr"`
	ShpingObrStv      float64 `json:"shping_obr_stv"`
	ShpingObrRam      float64 `json:"shping_obr_ram"`
	YstanZakld        float64 `json:"ystan_zakld"`
	SverlOtvShtift    float64 `json:"sverl_otv_shtift"`
	FrezerStoekRigel  float64 `json:"frezer_stoek_rigel"`
	SborkaRam         float64 `json:"sborka_ram"`
	ShitfRam          float64 `json:"shitf_ram"`
	OpresStv          float64 `json:"opres_stv"`
	SborStv           float64 `json:"sbor_stv"`
	ShiftStv          float64 `json:"shift_stv"`
	NanesKlei         float64 `json:"nanes_klei"`
	YstanYplRam       float64 `json:"ystan_ypl_ram"`
	YstanYplStv       float64 `json:"ystan_ypl_stv"`
	YstZamokNakl      float64 `json:"yst_zamok_nakl"`
	YstShpingOtv      float64 `json:"yst_shping_otv"`
	SborPetliRam      float64 `json:"sbor_petli_ram"`
	SborYstPorog      float64 `json:"sbor_yst_porog"`
	Naveshiv          float64 `json:"naveshiv"`
	Zashiv            float64 `json:"zashiv"`
	OpresRam          float64 `json:"opres_ram"`
	SborPetliStv      float64 `json:"sbor_petli_stv"`
	NastrStanok       float64 `json:"nastr_stanok"`
	NastrPbx          float64 `json:"nastr_pbx"`
	MehObrabPzr       float64 `json:"meh_obrab_pzr"`
	RabotaPbx         float64 `json:"rabota_pbx"`
	FrezerNastr       float64 `json:"frezer_nastr"`
	FrezerPorogSborka float64 `json:"frezer_porog_sborka"`
	FrezerYstShtyp    float64 `json:"frezer_yst_shtyp"`
	OpresNastr        float64 `json:"opres_nastr"`
	Opres             float64 `json:"opres"`
	PodgDerjShetki    float64 `json:"podg_derj_shetki"`
	YstPorogYplDr     float64 `json:"yst_porog_ypl_dr"`
	NaveshivStv       float64 `json:"naveshiv_stv"`
	YstZapoln         float64 `json:"yst_zapoln"`
	ImpostNapil       float64 `json:"impost_napil"`
	ImpostFrezer      float64 `json:"impost_frezer"`
	ImpostSverlo      float64 `json:"impost_sverlo"`
	ImpostYst         float64 `json:"impost_yst"`
	ImpostShtift      float64 `json:"impost_shtift"`
	YplFalc           float64 `json:"ypl_falc"`
	NapilNalich       float64 `json:"napil_nalich"`
	NapilRam          float64 `json:"napil_ram"`
	NapilStv          float64 `json:"napil_stv"`
	KontrSbork        float64 `json:"kontr_sbork"`
	SverlRam          float64 `json:"sverl_ram"`
	SverlZink         float64 `json:"sverl_zink"`
	ZashitPl          float64 `json:"zashit_pl"`
	SborRam           float64 `json:"sbor_ram"`
	NapilYstKrishStv  float64 `json:"napil_yst_krish_stv"`
	NapilYstKrishRam  float64 `json:"napil_yst_krish_ram"`
	SborPetli         float64 `json:"sbor_petli"`
	YstPtliRamStv     float64 `json:"yst_ptli_ram_stv"`
	RezkaPlast        float64 `json:"rezka_plast"`
	Brysok            float64 `json:"brysok"`
	IzgPritv          float64 `json:"izg_pritv"`
	ObrPritv          float64 `json:"obr_pritv"`
	YstPritv          float64 `json:"yst_pritv"`
	ObrabotkaAll      float64 `json:"obrabotka_all"`
	YstanPlnPetli     float64 `json:"ystan_pln_petli"`
	YstFetr           float64 `json:"yst_fetr"`
	Rezina            float64 `json:"rezina"`
	FrezerShping      float64 `json:"frezer_shping"`
	Gl                float64 `json:"gl"`
	Fortochka         float64 `json:"fortochka"`
	Upak              float64 `json:"ypak"`
	OrderId           int     `json:"order_id"`
	TotalTime         float64 `json:"total_time"`
}
