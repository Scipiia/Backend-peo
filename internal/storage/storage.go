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
	Position     int      `json:"position"`
	Creator      string   `json:"creator"`
	NamePosition string   `json:"name_position"`
	Count        *float64 `json:"count"`
	Image        *string  `json:"image"`
}

type ImportJSON struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsStudent bool   `json:"is_student"`
}

type ResultOrderDetails struct {
	Order         *Order           `json:"order-norm"`
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
	NapilKontr       float64 `json:"napil_kontyr"`
	NapilKrishek     float64 `json:"napil_krishek"`
	NapilImpost      float64 `json:"napil_impost"`
	Soedinitel       float64 `json:"soedinitel"`
	PromejSborka     float64 `json:"promej_sborka"`
	ImpostSverlovka  float64 `json:"impost_sverlo"`
	ImpostFrezerovka float64 `json:"impost_frezer"`
	ImpostSborka     float64 `json:"impost_sborka"`
	OpresNastr       float64 `json:"opres_nastr"`
	Opresovka        float64 `json:"opresovka"`
	YstanYplotnitel  float64 `json:"ystanovka_yplotn"`
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

type DemResultVitraj struct {
	ID                 int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderNum           string  `json:"order_num" gorm:"column:order_num"`
	Name               string  `json:"name" gorm:"column:name"`
	Count              int     `json:"count" gorm:"column:count"`
	Profil             string  `json:"profil" gorm:"column:profil"`
	PodgotOboryd       float64 `json:"podgot_oboryd" gorm:"column:podgot_oboryd"`
	NapilStkDo3m       float64 `json:"napil_stk_do3m" gorm:"column:napil_stk_do3m"`
	NapilStkBol3m      float64 `json:"napil_stk_bol3m" gorm:"column:napil_stk_bol3m"`
	NapilStkBol5m      float64 `json:"napil_stk_bol5m" gorm:"column:napil_stk_bol5m"`
	PZR                float64 `json:"pzr" gorm:"column:pzr"`
	NastrPbx           float64 `json:"nastr_pbx" gorm:"column:nastr_pbx"`
	DrenajOtv          float64 `json:"drenaj_otv" gorm:"column:drenaj_otv"`
	ZashelkObr         float64 `json:"zashelk_obr" gorm:"column:zashelk_obr"`
	PrisetPrObr        float64 `json:"priset_pr_obr" gorm:"column:priset_pr_obr"`
	YplYst             float64 `json:"ypl_yst" gorm:"column:ypl_yst"`
	FetrYst            float64 `json:"fetr_yst" gorm:"column:fetr_yst"`
	StekloYst          float64 `json:"steklo_yst" gorm:"column:steklo_yst"`
	RazmetOtv          float64 `json:"razmet_otv" gorm:"column:razmet_otv"`
	SverloOtv          float64 `json:"sverlo_otv" gorm:"column:sverlo_otv"`
	NakladYst          float64 `json:"naklad_yst" gorm:"column:naklad_yst"`
	Panel              float64 `json:"panel" gorm:"column:panel"`
	Naprav             float64 `json:"naprav" gorm:"column:naprav"`
	Ydlin              float64 `json:"ydlin" gorm:"column:ydlin"`
	Upak               float64 `json:"upak" gorm:"column:upak"`
	NapilRigelDo1m     float64 `json:"napil_rigel_do1m" gorm:"column:napil_rigel_do1m"`
	NapilRigelBol1m    float64 `json:"napil_rigel_bol1m" gorm:"column:napil_rigel_bol1m"`
	NakldNapil         float64 `json:"nakld_napil" gorm:"column:nakld_napil"`
	NapilKriskhSt      float64 `json:"napil_krish_st" gorm:"column:napil_krish_st"`
	NapilKriskhRg      float64 `json:"napil_krish_rg" gorm:"column:napil_krish_rg"`
	KomplKriskh        float64 `json:"kompl_krish" gorm:"column:kompl_krish"`
	NapilAdapt         float64 `json:"napil_adapt" gorm:"column:napil_adapt"`
	FrezerRigel        float64 `json:"frezer_rigel" gorm:"column:frezer_rigel"`
	FukelYst           float64 `json:"fukel_yst" gorm:"column:fukel_yst"`
	StoikiPbx          float64 `json:"stoiki_pbx" gorm:"column:stoiki_pbx"`
	ZamokYst           float64 `json:"zamok_yst" gorm:"column:zamok_yst"`
	ZamokYstBolt       float64 `json:"zamok_yst_bolt" gorm:"column:zamok_yst_bolt"`
	RigelSverloZamok   float64 `json:"rigel_sverlo_zamok" gorm:"column:rigel_sverlo_zamok"`
	NakladSverlo       float64 `json:"naklad_sverlo" gorm:"column:naklad_sverlo"`
	YplYstRigel        float64 `json:"ypl_yst_rigel" gorm:"column:ypl_yst_rigel"`
	YplYstNakld        float64 `json:"ypl_yst_nakld" gorm:"column:ypl_yst_nakld"`
	YplYstStoik        float64 `json:"ypl_yst_stoik" gorm:"column:ypl_yst_stoik"`
	AdaptYstStoik      float64 `json:"adapt_yst_stoik" gorm:"column:adapt_yst_stoik"`
	AdaptYstRigel      float64 `json:"adapt_yst_rigel" gorm:"column:adapt_yst_rigel"`
	TermovstYstStoiki  float64 `json:"termovst_yst_stoiki" gorm:"column:termovst_yst_stoiki"`
	TermovstYstRigel   float64 `json:"termovst_yst_rigel" gorm:"column:termovst_yst_rigel"`
	SborNog            float64 `json:"sbor_nog" gorm:"column:sbor_nog"`
	Birki              float64 `json:"birki" gorm:"column:birki"`
	UpakStoik          float64 `json:"upak_stoik" gorm:"column:upak_stoik"`
	UpakPet            float64 `json:"upak_pet" gorm:"column:upak_pet"`
	UpakRigel          float64 `json:"upak_rigel" gorm:"column:upak_rigel"`
	UpakRigel2m        float64 `json:"upak_rigel_2m" gorm:"column:upak_rigel_2m"`
	UpakKriskh         float64 `json:"upak_krish" gorm:"column:upak_krish"`
	UpakNakld          float64 `json:"upak_nakld" gorm:"column:upak_nakld"`
	UpakYplNog         float64 `json:"upak_ypl_nog" gorm:"column:upak_ypl_nog"`
	UpakKronsht        float64 `json:"upak_kronsht" gorm:"column:upak_kronsht"`
	VinosGotovIzd      float64 `json:"vinos_gotov_izd" gorm:"column:vinos_gotov_izd"`
	YstStikZakld       float64 `json:"yst_stik_zakld" gorm:"column:yst_stik_zakld"`
	ObnYsovDo4m        float64 `json:"obn_ysov_do4m" gorm:"column:obn_ysov_do4m"`
	ObnYsovBol4m       float64 `json:"obn_ysov_bol4m" gorm:"column:obn_ysov_bol4m"`
	OtmRezin           float64 `json:"otm_rezin" gorm:"column:otm_rezin"`
	ObnNastr           float64 `json:"obn_nastr" gorm:"column:obn_nastr"`
	TrybaProf          float64 `json:"tryba_prof" gorm:"column:tryba_prof"`
	KomplShtapik       float64 `json:"kompl_shtapik" gorm:"column:kompl_shtapik"`
	OtvVO              float64 `json:"otv_vo" gorm:"column:otv_vo"`
	NastrStanokRigel   float64 `json:"nastr_stanok_rigel" gorm:"column:nastr_stanok_rigel"`
	FrezerStoikiPr     float64 `json:"frezer_stoiki_pr" gorm:"column:frezer_stoiki_pr"`
	NastrStanokStoiki1 float64 `json:"nastr_stanok_stoiki_1" gorm:"column:nastr_stanok_stoiki_1"`
	FrezerStoikiYgl    float64 `json:"frezer_stoiki_ygl" gorm:"column:frezer_stoiki_ygl"`
	NastrStanokStoiki2 float64 `json:"nastr_stanok_stoiki_2" gorm:"column:nastr_stanok_stoiki_2"`
	ZashitPl           float64 `json:"zashit_pl" gorm:"column:zashit_pl"`
	YstPritv           float64 `json:"yst_pritv" gorm:"column:yst_pritv"`
	YstKapel           float64 `json:"yst_kapel" gorm:"column:yst_kapel"`
	SborSekci          float64 `json:"sbor_sekci" gorm:"column:sbor_sekci"`
	GermetYpl          float64 `json:"germet_ypl" gorm:"column:germet_ypl"`
	UpakSoed           float64 `json:"upak_soed" gorm:"column:upak_soed"`
	UpakShtapik        float64 `json:"upak_shtapik" gorm:"column:upak_shtapik"`
	UpakSekcii         float64 `json:"upak_sekcii" gorm:"column:upak_sekcii"`
	Obrezanie          float64 `json:"obrezanie" gorm:"column:obrezanie"`
	VinosSekcii        float64 `json:"vinos_sekcii" gorm:"column:vinos_sekcii"`
	YgolKorob          float64 `json:"ygol_korob" gorm:"column:ygol_korob"`
	PbxVO              float64 `json:"pbx_vo" gorm:"column:pbx_vo"`
	ClearVO            float64 `json:"clear_vo" gorm:"column:clear_vo"`
	YstKomptrSt        float64 `json:"yst_komptr_st" gorm:"column:yst_komptr_st"`
	YstKomptrRg        float64 `json:"yst_komptr_rg" gorm:"column:yst_komptr_rg"`
	UpakYgl            float64 `json:"upak_ygl" gorm:"column:upak_ygl"`
	NapilYgl           float64 `json:"napil_ygl" gorm:"column:napil_ygl"`
	NarezTermovst      float64 `json:"narez_termovst" gorm:"column:narez_termovst"`
	NarezKompensr      float64 `json:"narez_kompensr" gorm:"column:narez_kompensr"`
	UpakTermovst       float64 `json:"upak_termovst" gorm:"column:upak_termovst"`
	RigelSverloZamok3m float64 `json:"rigel_sverlo_zamok_3m" gorm:"column:rigel_sverlo_zamok_3m"`
	RigelNsverloZamok  float64 `json:"rigel_nsverlo_zamok" gorm:"column:rigel_nsverlo_zamok"`
	ObjYsovTermvst     float64 `json:"obj_ysov_termvst" gorm:"column:obj_ysov_termvst"`
	YstKPS             float64 `json:"yst_kps" gorm:"column:yst_kps"`
	NapilStkBol6m      float64 `json:"napil_stk_bol6m" gorm:"column:napil_stk_bol6m"`
	NapilShtapik       float64 `json:"napil_shtapik" gorm:"column:napil_shtapik"`
	RezProfil          float64 `json:"rez_profil" gorm:"column:rez_profil"`
	RezSekcii          float64 `json:"rez_sekcii" gorm:"column:rez_sekcii"`
	PartSbSeck         float64 `json:"part_sb_seck" gorm:"column:part_sb_seck"`
	PartSbMarkSeck     float64 `json:"part_sb_mark_seck" gorm:"column:part_sb_mark_seck"`
	PartSbStoek        float64 `json:"part_sb_stoek" gorm:"column:part_sb_stoek"`
	YstFormr           float64 `json:"yst_formr" gorm:"column:yst_formr"`
	OrderID            int     `json:"order_id" gorm:"column:order_id"`
	TotalTime          float64 `json:"total_time" gorm:"column:total_time"`
}

type DemResultLoggia struct {
	ID                int     `json:"id" gorm:"column:id;primaryKey"`
	OrderNum          string  `json:"order_num" gorm:"column:order_num"`
	Name              string  `json:"name" gorm:"column:name"`
	Count             int     `json:"count" gorm:"column:count"`
	Profil            string  `json:"profil" gorm:"column:profil"`
	PodgotOboryd      float64 `json:"podgot_oboryd" gorm:"column:podgot_oboryd"`
	NapilStkDo3m      float64 `json:"napil_stk_do3m" gorm:"column:napil_stk_do3m"`
	NapilRigelDo1m    float64 `json:"napil_rigel_do1m" gorm:"column:napil_rigel_do1m"`
	NapilRigelBol1m   float64 `json:"napil_rigel_bol1m" gorm:"column:napil_rigel_bol1m"`
	NapilShtapik      float64 `json:"napil_shtapik" gorm:"column:napil_shtapik"`
	KomplShtapik      float64 `json:"kompl_shtapik" gorm:"column:kompl_shtapik"`
	NapilAdaptTr      float64 `json:"napil_adapt_tr" gorm:"column:napil_adapt_tr"`
	RigelFrezer       float64 `json:"rigel_frezer" gorm:"column:rigel_frezer"`
	FrezerRigelZamok  float64 `json:"frezer_rigel_zamok" gorm:"column:frezer_rigel_zamok"`
	ZamokYst          float64 `json:"zamok_yst" gorm:"column:zamok_yst"`
	RezPgmSt          float64 `json:"rez_pgm_st" gorm:"column:rez_pgm_st"`
	RezPgmRg          float64 `json:"rez_pgm_rg" gorm:"column:rez_pgm_rg"`
	FrezVo            float64 `json:"frez_vo" gorm:"column:frez_vo"`
	PartSbSekci       float64 `json:"part_sb_sekci" gorm:"column:part_sb_sekci"`
	YstPritv          float64 `json:"yst_pritv" gorm:"column:yst_pritv"`
	YstFormir         float64 `json:"yst_formir" gorm:"column:yst_formir"`
	Birki             float64 `json:"birki" gorm:"column:birki"`
	UpakRam           float64 `json:"upak_ram" gorm:"column:upak_ram"`
	UpakStoik         float64 `json:"upak_stoik" gorm:"column:upak_stoik"`
	UpakRigel         float64 `json:"upak_rigel" gorm:"column:upak_rigel"`
	UpakRigel2m       float64 `json:"upak_rigel_2m" gorm:"column:upak_rigel_2m"`
	UpakShtapik       float64 `json:"upak_shtapik" gorm:"column:upak_shtapik"`
	UpakAdaptTr       float64 `json:"upak_adapt_tr" gorm:"column:upak_adapt_tr"`
	UpakYplNog        float64 `json:"upak_ypl_nog" gorm:"column:upak_ypl_nog"`
	VinosGotovIzd     float64 `json:"vinos_gotov_izd" gorm:"column:vinos_gotov_izd"`
	NapilPrVirav      float64 `json:"napil_pr_virav" gorm:"column:napil_pr_virav"`
	NapilRam          float64 `json:"napil_ram" gorm:"column:napil_ram"`
	NapilStv          float64 `json:"napil_stv" gorm:"column:napil_stv"`
	NapilPritv        float64 `json:"napil_pritv" gorm:"column:napil_pritv"`
	NapilSoed         float64 `json:"napil_soed" gorm:"column:napil_soed"`
	FrezRam           float64 `json:"frez_ram" gorm:"column:frez_ram"`
	FrezStv           float64 `json:"frez_stv" gorm:"column:frez_stv"`
	FrezPritv         float64 `json:"frez_pritv" gorm:"column:frez_pritv"`
	ClearSverlStv     float64 `json:"clear_sverl_stv" gorm:"column:clear_sverl_stv"`
	ClearSverlYsilStv float64 `json:"clear_sverl_ysil_stv" gorm:"column:clear_sverl_ysil_stv"`
	Kraska            float64 `json:"kraska" gorm:"column:kraska"`
	YstRolik          float64 `json:"yst_rolik" gorm:"column:yst_rolik"`
	YstZashel         float64 `json:"yst_zashel" gorm:"column:yst_zashel"`
	Rezina            float64 `json:"rezina" gorm:"column:rezina"`
	SborStv           float64 `json:"sbor_stv" gorm:"column:sbor_stv"`
	SborPritv         float64 `json:"sbor_pritv" gorm:"column:sbor_pritv"`
	PodgKompl         float64 `json:"podg_kompl" gorm:"column:podg_kompl"`
	ShtampStStv       float64 `json:"shtamp_st_stv" gorm:"column:shtamp_st_stv"`
	PodgRezin         float64 `json:"podg_rezin" gorm:"column:podg_rezin"`
	SborYpl           float64 `json:"sbor_ypl" gorm:"column:sbor_ypl"`
	YstZaklep         float64 `json:"yst_zaklep" gorm:"column:yst_zaklep"`
	OrderID           int     `json:"order_id" gorm:"column:order_id"`
	TotalTime         float64 `json:"total_time" gorm:"column:total_time"`
}

type ProductItem struct {
	Type      string    `json:"type"` // "loggia", "vitraj", и т.д.
	OrderNum  string    `json:"order_num"`
	Name      string    `json:"name"`
	Count     int       `json:"count"`
	Profil    string    `json:"profil"`
	TotalTime float64   `json:"total_time"`
	OrderID   int64     `json:"order_id"`
	ResultID  int       `json:"result_id"` // ID в таблице результата (например, loggia.id)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type OperationValue struct {
	ID    string  `json:"id"`    // соответствует имени поля в БД: "nast_napil", "napil" и т.д.
	Value float64 `json:"value"` // новое значение времени (в минутах)
}

//type UpdateGlyhariRequest struct {
//	OrderNum   string             `json:"order_num"`
//	Operations map[string]float64 `json:"operations"` // например: {"nast_napil": 0.08}
//}

type AdditionalOperation struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	Operation string    `json:"operation"`
	Duration  float64   `json:"duration"`
	Comment   *string   `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
