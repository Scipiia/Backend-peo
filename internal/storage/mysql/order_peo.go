package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
	"vue-golang/utils"
)

func (s *Storage) GetOrderDetails(id int) (*storage.ResultOrderDetails, error) {
	const op = "storage.order-norm-details.sql"

	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	//TODO основная структура
	details := &storage.ResultOrderDetails{}

	stmtDemOrders := "SELECT id, order_num, creator, customer, dop_info, ms_note FROM dem_ready WHERE id = ?"

	//order-norm := &storage.Order{}
	var msNote sql.NullString // Используем NullString для обработки NULL
	details.Order = &storage.Order{}
	err = tx.QueryRow(stmtDemOrders, id).Scan(&details.Order.ID, &details.Order.OrderNum, &details.Order.Creator, &details.Order.Customer, &details.Order.DopInfo, &msNote)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no rows found for query 1: %w", op, err)
		}
		return nil, fmt.Errorf("%s: query 1 failed: %w", op, err)
	}

	// Если значение NULL, заменяем его на пустую строку
	if msNote.Valid {
		details.Order.MsNote = msNote.String
	} else {
		details.Order.MsNote = ""
	}

	stmtDemPrice := `SELECT CAST( p.position AS UNSIGNED ),p.creator, p.name_position, p.kol_vo, i.im_image
		FROM dem_price p 
		left join dem_images i on i.im_ordername = p.numorders and i.im_orderpos = p.position
		WHERE p.numorders LIKE ?
		ORDER BY 1;`

	//TODO тут будет массив данных
	details.OrderDemPrice = []*storage.OrderDemPrice{}

	rows, err := tx.Query(stmtDemPrice, details.Order.OrderNum)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		price := &storage.OrderDemPrice{}

		err := rows.Scan(
			&price.Position,
			&price.Creator,
			&price.NamePosition,
			&price.Count,
			&price.Image,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}

		details.OrderDemPrice = append(details.OrderDemPrice, price)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return details, nil
}

func (s *Storage) GetForm(id int) (*storage.FormPeo, error) {
	const op = "storage.mysql.sql.GetForm"

	stmtFormPeo := "SELECT id, name, fields FROM dem_test_golang_form_json WHERE id = ?"

	// Создаем экземпляр структуры для хранения данных формы
	form := &storage.FormPeo{}

	// Считываем данные из базы данных
	var fieldsJSON string
	err := s.db.QueryRow(stmtFormPeo, id).Scan(&form.ID, &form.Name, &fieldsJSON)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no rows found for query FormPeo: %w", op, err)
		}
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}

	//fmt.Printf("Raw fieldsJSON: %s\n", fieldsJSON)

	// Преобразуем JSON-строку в массив объектов
	err = json.Unmarshal([]byte(fieldsJSON), &form.FieldsPeo)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal JSON for FieldsPeo: %w", op, err)
	}

	return form, nil
}

func (s *Storage) GetWorkers() ([]*storage.Workers, error) {
	const op = "storage.mysql.sql.GetWorkers"
	stmt := "SELECT * FROM dem_test_golang_workers"

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: no rows found for query Workers: %w", op, err)
	}

	defer rows.Close()

	var workers []*storage.Workers

	for rows.Next() {
		worker := &storage.Workers{}

		err := rows.Scan(&worker.ID, &worker.LastName, &worker.Profession)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}

		workers = append(workers, worker)
	}

	return workers, nil

}

func (s *Storage) SaveWorker(resWorker storage.WorkersResult) error {
	const op = "storage.mysql.sql.saveWorker"
	stmt := `INSERT INTO dem_test_golang_result_worker (order_num, name_izd, operation_name, worker_id, value, count, assigned_at)
			VALUES (?,?,?,?,?,?,?)`

	exec, err := s.db.Exec(stmt, resWorker.OrderNum, resWorker.NameIzd, resWorker.OperationName, resWorker.WorkerId, resWorker.Value, resWorker.Count, resWorker.AssignedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println(exec)
	//insertId, err := exec.LastInsertId()
	//if err != nil {
	//	return 0, fmt.Errorf("%s: %w", op, err)
	//}

	return nil
}

// TODO сохраняются все данные после нормирования, пока только глухари
func (s *Storage) SaveGlyhari(result storage.DemResultGlyhari) (int, error) {
	const op = "storage.mysql.sql.SaveGlyhari"
	// Шаг 1: Получить или создать заказ по order_num
	orderID, err := s.getOrCreateOrder(result.OrderNum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get/create order-norm: %w", op, err)
	}

	// Шаг 2: Подставить orderID в результат
	result.OrderId = orderID

	stmt := `INSERT INTO dem_test_golang_result_glyhar (order_num, name, count,  nast_napil, napil_kontyr, napil_krishek, napil_impost,
	soedinitel, promej_sborka, impost_sverlo, impost_frezer,impost_sborka, opres_nastr, opresovka, ystanovka_yplotn, zashivka, profil,
	napil_stoiki_do3m, napil_stoiki_bol3m, napil_rigel_do1m, napil_rigel_bol1m, sverl_rigel_zamok, ystan_zamkov, napil_shtapik, ypakovka,
    frezer_rigel, obrabot_ram, hands_sborka, frezer_nastr, shtiftovka, ystanovka_zapoln, napil_donnik, adapter_napil, adapter_ystan, ystan_yplotn_falc, order_id, total_time)
	VALUES (?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?)`

	exec, err := s.db.Exec(stmt, result.OrderNum, result.Name, result.Count, result.NastrNapil, result.NapilKontr, result.NapilKrishek,
		result.NapilImpost, result.Soedinitel, result.PromejSborka, result.ImpostSverlovka, result.ImpostFrezerovka, result.ImpostSborka,
		result.OpresNastr, result.Opresovka, result.YstanYplotnitel, result.Zashivka, result.Profil, result.NapilStoikiDo3m, result.NapilStoikiBol3m,
		result.NapilRigelDo1m, result.NapilRigelBol1m, result.SverlRigelZamok, result.YstanZamkov, result.NapilShtapik, result.Ypakovka, result.FrezerRigel,
		result.ObrabotRam, result.HandsSborka, result.FrezerNastr, result.Shtiftovka, result.YstanovkaZapoln, result.NapilDonnik, result.AdapterNapil, result.AdapterYstan,
		result.YstanYplotnFalc, result.OrderId, result.TotalTime)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	insertId, err := exec.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int(insertId), nil
}

func (s *Storage) GetGlyhari(id int) (*storage.OrderData, error) {
	const op = "storage.mysql.sql.GetGlyhari"

	stmt := `SELECT id, order_num, name, count,  nast_napil, napil_kontyr, napil_krishek, napil_impost,
		soedinitel, promej_sborka, impost_sverlo, impost_frezer,impost_sborka, opres_nastr, opresovka, ystanovka_yplotn, zashivka, profil,
		napil_stoiki_do3m, napil_stoiki_bol3m, napil_rigel_do1m, napil_rigel_bol1m, sverl_rigel_zamok, ystan_zamkov, napil_shtapik, ypakovka,
    	frezer_rigel, obrabot_ram, hands_sborka, frezer_nastr, shtiftovka, ystanovka_zapoln, napil_donnik, adapter_napil, adapter_ystan, ystan_yplotn_falc, order_id, total_time
		FROM dem_test_golang_result_glyhar WHERE id = ?`

	var result storage.DemResultGlyhari

	err := s.db.QueryRow(stmt, id).Scan(
		&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.NastrNapil, &result.NapilKontr, &result.NapilKrishek, &result.NapilImpost,
		&result.Soedinitel, &result.PromejSborka, &result.ImpostSverlovka, &result.ImpostFrezerovka, &result.ImpostSborka, &result.OpresNastr,
		&result.Opresovka, &result.YstanYplotnitel, &result.Zashivka, &result.Profil, &result.NapilStoikiDo3m, &result.NapilStoikiBol3m,
		&result.NapilRigelDo1m, &result.NapilRigelBol1m, &result.SverlRigelZamok, &result.YstanZamkov, &result.NapilShtapik, &result.Ypakovka,
		&result.FrezerRigel, &result.ObrabotRam, &result.HandsSborka, &result.FrezerNastr, &result.Shtiftovka, &result.YstanovkaZapoln,
		&result.NapilDonnik, &result.AdapterNapil, &result.AdapterYstan, &result.YstanYplotnFalc, &result.OrderId, &result.TotalTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for order_num: %s", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	//return &result, nil
	return utils.MapToOrderDataGlyhari(&result), nil
}

func (s *Storage) SaveWindows(result storage.DemResultWindow) (int, error) {
	const op = "storage.mysql.sql.SaveWindows"
	// Шаг 1: Получить или создать заказ по order_num
	orderID, err := s.getOrCreateOrder(result.OrderNum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get/create order-norm: %w", op, err)
	}

	// Шаг 2: Подставить orderID в результат
	result.OrderId = orderID

	stmt := `INSERT INTO dem_test_golang_result_window (
  	order_num, name, count, profil,podgot_oboryd, napil_ram_stv, napil_nram_stv, napil_imposta, sverlovka_stoek,obrabot_ram, obrabotka_stv, promej_sborka_stv, 
  	napil_adaptera, promej_sborka_ram,promej_sborka_gl, opres_ram, opres_gl, frezer_stv, ystan_ypl_ram, ystan_ypl_stv,napil_tag, sverlo_tag, ystan_furn, 
	navesh_stv, zashivka, obrab_dop_profil, ystan_adaptera, frezer_impost_pilon, krepl_ygl_ram, glyhar_dorab, yplot_glyhar,zashiv_glyhar, opres_stv,
    razbor_erkera, gl_window, obrabot_shtylp,frezer_pod_shtylp, ystan_shtylp, ystan_impost, nastr_for_opres, napil_ram,napil_nram, napil_dop_profil,
	frezer_dop_profil, obrabotka_stv_rychka,obrabotka_stv_zamok, napil_shtylp, impost_sbor, rezina_ram, podgotov_furn,podgotov_ram, razborka, napil_zashit_prof, 
    impost_frezer, ystan_stekla_zaliv,ystan_krishek, frezer_profil_zamazka, napil_stoiki_do3m, napil_stoiki_bol3m,napil_rigel_do1m, napil_rigel_bol1m, 
    sverlo_rigel_zamok, ystan_zamok,shtiftovka, frezer_rigel, part_sborka, ystan_rez_ygl, opres_ygl, ystan_termovst, napil_shtapik, ypak_sekcii, ypak_rigel,
    ypak_yplotn,ypak_shtapik, napil_donnika,nastr_pbx, meh_obrab_pzr, rabota_pbx,slesar_obr_furn, impost_sverlo, opresovka, sbor_rychka, sbor_petli,plastik_ystn_ram,
    napil_stv, ystan_furn_stv, ystan_furn_ram, compl_otg, ypak_izd, otgryz, rasp_furn, order_id, total_time
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
          ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	exec, err := s.db.Exec(stmt, result.OrderNum, result.Name, result.Count, result.Profil, result.PodgotOboryd, result.NapilRamStv, result.NapilNramStv, result.NapilImposta,
		result.SverlovkaStoek, result.ObrabotRam, result.ObrabotkaStv, result.PromejSborkaStv, result.NapilAdaptera, result.PromejSborkaRam, result.PromejSborkaGl, result.OpresRam,
		result.OpresGl, result.FrezerStv, result.YstanYplRam, result.YstanYplStv, result.NapilTag, result.SverloTag, result.YstanFurn,
		result.NaveshStv, result.Zashivka, result.ObrabDopProfil, result.YstanAdaptera, result.FrezerImpostPilon, result.KreplYglRam, result.GlyharDorab,
		result.YplotGlyhar, result.ZashivGlyhar, result.OpresStv, result.RazborErkera, result.GlWindow, result.ObrabotShtylp, result.FrezerPodShtylp, result.YstanShtylp,
		result.YstanImpost, result.NastrForOpres, result.NapilRam, result.NapilNram, result.NapilDopProfil, result.FrezerDopProfil, result.ObrabotkaStvRychka, result.ObrabotkaStvZamok,
		result.NapilShtylp, result.ImpostSbor, result.RezinaRam, result.PodgotovFurn, result.PodgotovRam, result.Razborka, result.NapilZashitProf, result.ImpostFrezer,
		result.YstanSteklaZaliv, result.YstanKrishek, result.FrezerProfilZamazka, result.NapilStoikiDo3m, result.NapilStoikiBol3m, result.NapilRigelDo1m, result.NapilRigelBol1m,
		result.SverloRigelZamok, result.YstanZamok, result.Shtiftovka, result.FrezerRigel, result.PartSborka, result.YstanRezYgl, result.OpresYgl, result.YstanTermovst, result.NapilShtapik,
		result.YpakSekcii, result.YpakRigel, result.YpakYplotn, result.YpakShtapik, result.NapilDonnika, result.NastrPbx, result.MehObrabPzr, result.RabotaPbx,
		result.SlesarObrFurn, result.ImpostSverlo, result.Opresovka, result.SborRychka, result.SborPetli, result.PlastikYstnRam, result.NapilStv, result.YstanFurnStv,
		result.YstanFurnRam, result.ComplOtg, result.YpakIzd, result.Otgryz, result.RaspFurn, result.OrderId, result.TotalTime,
	)

	if err != nil {
		return 0, fmt.Errorf("%s: failed to insert record in dem result window: %w", op, err)
	}

	insertId, err := exec.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int(insertId), nil
}

func (s *Storage) GetWindows(id int) (*storage.OrderData, error) {
	const op = "storage.mysql.sql.GetWindows"

	stmt := `SELECT 
    	id, order_num, name, count, profil, podgot_oboryd, napil_ram_stv, napil_nram_stv, napil_imposta,
    	sverlovka_stoek, obrabot_ram, obrabotka_stv, promej_sborka_stv, napil_adaptera, promej_sborka_ram,
    	promej_sborka_gl, opres_ram, opres_gl, frezer_stv, ystan_ypl_ram, ystan_ypl_stv, napil_tag,
    	sverlo_tag, ystan_furn, navesh_stv, zashivka, obrab_dop_profil, ystan_adaptera, frezer_impost_pilon,
    	krepl_ygl_ram, glyhar_dorab, yplot_glyhar, zashiv_glyhar, opres_stv, razbor_erkera, gl_window,
    	obrabot_shtylp, frezer_pod_shtylp, ystan_shtylp, ystan_impost, nastr_for_opres, napil_ram,
    	napil_nram, napil_dop_profil, frezer_dop_profil, obrabotka_stv_rychka, obrabotka_stv_zamok,
    	napil_shtylp, impost_sbor, rezina_ram, podgotov_furn, podgotov_ram, razborka, napil_zashit_prof,
    	impost_frezer, ystan_stekla_zaliv, ystan_krishek, frezer_profil_zamazka, napil_stoiki_do3m,
    	napil_stoiki_bol3m, napil_rigel_do1m, napil_rigel_bol1m, sverlo_rigel_zamok, ystan_zamok,
    	shtiftovka, frezer_rigel, part_sborka, ystan_rez_ygl, opres_ygl, ystan_termovst, napil_shtapik,
    	ypak_sekcii, ypak_rigel, ypak_yplotn, ypak_shtapik, napil_donnika, nastr_pbx, meh_obrab_pzr,
    	rabota_pbx, slesar_obr_furn, impost_sverlo, opresovka, sbor_rychka, sbor_petli, plastik_ystn_ram,
   		 napil_stv, ystan_furn_stv, ystan_furn_ram, compl_otg, ypak_izd, otgryz, rasp_furn, order_id, total_time
	FROM dem_test_golang_result_window WHERE id = ?`

	var result storage.DemResultWindow

	err := s.db.QueryRow(stmt, id).Scan(
		&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.Profil, &result.PodgotOboryd,
		&result.NapilRamStv, &result.NapilNramStv, &result.NapilImposta, &result.SverlovkaStoek,
		&result.ObrabotRam, &result.ObrabotkaStv, &result.PromejSborkaStv, &result.NapilAdaptera,
		&result.PromejSborkaRam, &result.PromejSborkaGl, &result.OpresRam, &result.OpresGl,
		&result.FrezerStv, &result.YstanYplRam, &result.YstanYplStv, &result.NapilTag,
		&result.SverloTag, &result.YstanFurn, &result.NaveshStv, &result.Zashivka,
		&result.ObrabDopProfil, &result.YstanAdaptera, &result.FrezerImpostPilon,
		&result.KreplYglRam, &result.GlyharDorab, &result.YplotGlyhar, &result.ZashivGlyhar,
		&result.OpresStv, &result.RazborErkera, &result.GlWindow, &result.ObrabotShtylp,
		&result.FrezerPodShtylp, &result.YstanShtylp, &result.YstanImpost, &result.NastrForOpres,
		&result.NapilRam, &result.NapilNram, &result.NapilDopProfil, &result.FrezerDopProfil,
		&result.ObrabotkaStvRychka, &result.ObrabotkaStvZamok, &result.NapilShtylp,
		&result.ImpostSbor, &result.RezinaRam, &result.PodgotovFurn, &result.PodgotovRam,
		&result.Razborka, &result.NapilZashitProf, &result.ImpostFrezer, &result.YstanSteklaZaliv,
		&result.YstanKrishek, &result.FrezerProfilZamazka, &result.NapilStoikiDo3m,
		&result.NapilStoikiBol3m, &result.NapilRigelDo1m, &result.NapilRigelBol1m,
		&result.SverloRigelZamok, &result.YstanZamok, &result.Shtiftovka, &result.FrezerRigel,
		&result.PartSborka, &result.YstanRezYgl, &result.OpresYgl, &result.YstanTermovst,
		&result.NapilShtapik, &result.YpakSekcii, &result.YpakRigel, &result.YpakYplotn,
		&result.YpakShtapik, &result.NapilDonnika, &result.NastrPbx, &result.MehObrabPzr,
		&result.RabotaPbx, &result.SlesarObrFurn, &result.ImpostSverlo, &result.Opresovka,
		&result.SborRychka, &result.SborPetli, &result.PlastikYstnRam, &result.NapilStv,
		&result.YstanFurnStv, &result.YstanFurnRam, &result.ComplOtg, &result.YpakIzd,
		&result.Otgryz, &result.RaspFurn, &result.OrderId, &result.TotalTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for order_num: %s", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	//return &result, nil
	return utils.MapToOrderDataWindow(&result), nil
}

func (s *Storage) SaveDoor(result storage.DemResultDoor) (int, error) {
	const op = "storage.mysql.sql.SaveDoor"

	// Шаг 1: Получить или создать заказ по order_num
	orderID, err := s.getOrCreateOrder(result.OrderNum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get/create order-norm: %w", op, err)
	}

	result.OrderId = orderID

	// Подготовка SQL-запроса для таблицы dem_test_golang_result_door
	stmt := `INSERT INTO dem_test_golang_result_door (
        order_num, name, count, profil, napil_ram_stv, petli_obr, zamok_obr, shping_obr_stv, shping_obr_ram, ystan_zakld,sverl_otv_shtift, frezer_stoek_rigel, sborka_ram, 
		shitf_ram, opres_stv, sbor_stv, shift_stv, nanes_klei, ystan_ypl_ram, ystan_ypl_stv, yst_zamok_nakl, yst_shping_otv, sbor_petli_ram, sbor_yst_porog, naveshiv,
        zashiv, opres_ram, sbor_petli_stv, nastr_stanok, nastr_pbx, meh_obrab_pzr, rabota_pbx, frezer_nastr, frezer_porog_sborka, frezer_yst_shtyp, opres_nastr, 
        opres, podg_derj_shetki, yst_porog_ypl_dr, naveshiv_stv, yst_zapoln, impost_napil, impost_frezer, impost_sverlo, impost_yst, impost_shtift, ypl_falc, napil_nalich, 
        napil_ram, napil_stv, kontr_sbork, sverl_ram, sverl_zink, zashit_pl, sbor_ram, napil_yst_krish_stv, napil_yst_krish_ram, sbor_petli, yst_ptli_ram_stv, rezka_plast, brysok, 
        izg_pritv, obr_pritv, yst_pritv, obrabotka_all, ystan_pln_petli, yst_fetr, rezina, frezer_shping, gl, fortochka, ypak, order_id, total_time
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
              ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Выполнение запроса с передачей всех значений
	res, err := s.db.Exec(stmt,
		result.OrderNum, result.Name, result.Count, result.Profil, result.NapilRamStv,
		result.PetliObr, result.ZamokObr, result.ShpingObrStv, result.ShpingObrRam, result.YstanZakld, result.SverlOtvShtift, result.FrezerStoekRigel, result.SborkaRam,
		result.ShitfRam, result.OpresStv, result.SborStv, result.ShiftStv, result.NanesKlei, result.YstanYplRam, result.YstanYplStv, result.YstZamokNakl, result.YstShpingOtv,
		result.SborPetliRam, result.SborYstPorog, result.Naveshiv, result.Zashiv, result.OpresRam, result.SborPetliStv, result.NastrStanok, result.NastrPbx,
		result.MehObrabPzr, result.RabotaPbx, result.FrezerNastr, result.FrezerPorogSborka, result.FrezerYstShtyp, result.OpresNastr, result.Opres, result.PodgDerjShetki,
		result.YstPorogYplDr, result.NaveshivStv, result.YstZapoln, result.ImpostNapil, result.ImpostFrezer, result.ImpostSverlo, result.ImpostYst, result.ImpostShtift,
		result.YplFalc, result.NapilNalich, result.NapilRam, result.NapilStv, result.KontrSbork, result.SverlRam, result.SverlZink, result.ZashitPl, result.SborRam,
		result.NapilYstKrishStv, result.NapilYstKrishRam, result.SborPetli, result.YstPtliRamStv, result.RezkaPlast, result.Brysok, result.IzgPritv, result.ObrPritv,
		result.YstPritv, result.ObrabotkaAll, result.YstanPlnPetli, result.YstFetr, result.Rezina, result.FrezerShping, result.Gl, result.Fortochka,
		result.Upak, result.OrderId, result.TotalTime,
	)

	if err != nil {
		return 0, fmt.Errorf("%s: failed to insert record into dem_test_golang_result_door: %w", op, err)
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return int(insertID), nil
}

func (s *Storage) GetDoor(id int) (*storage.OrderData, error) {
	const op = "storage.mysql.sql.GetDoor"

	stmt := `SELECT 
        id, order_num, name, count, profil,
        napil_ram_stv, petli_obr, zamok_obr, shping_obr_stv, shping_obr_ram,
        ystan_zakld, sverl_otv_shtift, frezer_stoek_rigel, sborka_ram, shitf_ram,
        opres_stv, sbor_stv, shift_stv, nanes_klei, ystan_ypl_ram,
        ystan_ypl_stv, yst_zamok_nakl, yst_shping_otv, sbor_petli_ram, sbor_yst_porog,
        naveshiv, zashiv, opres_ram, sbor_petli_stv, nastr_stanok,
        nastr_pbx, meh_obrab_pzr, rabota_pbx, frezer_nastr, frezer_porog_sborka,
        frezer_yst_shtyp, opres_nastr, opres, podg_derj_shetki, yst_porog_ypl_dr,
        naveshiv_stv, yst_zapoln, impost_napil, impost_frezer, impost_sverlo,
        impost_yst, impost_shtift, ypl_falc, napil_nalich, napil_ram,
        napil_stv, kontr_sbork, sverl_ram, sverl_zink, zashit_pl,
        sbor_ram, napil_yst_krish_stv, napil_yst_krish_ram, sbor_petli, yst_ptli_ram_stv,
        rezka_plast, brysok, izg_pritv, obr_pritv, yst_pritv,
        obrabotka_all, ystan_pln_petli, yst_fetr, rezina, frezer_shping,
        gl, fortochka, ypak, order_id, total_time
    FROM dem_test_golang_result_door 
    WHERE id = ?`

	var result storage.DemResultDoor

	err := s.db.QueryRow(stmt, id).Scan(&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.Profil, &result.NapilRamStv,
		&result.PetliObr, &result.ZamokObr, &result.ShpingObrStv, &result.ShpingObrRam, &result.YstanZakld, &result.SverlOtvShtift, &result.FrezerStoekRigel,
		&result.SborkaRam, &result.ShitfRam, &result.OpresStv, &result.SborStv, &result.ShiftStv, &result.NanesKlei, &result.YstanYplRam, &result.YstanYplStv,
		&result.YstZamokNakl, &result.YstShpingOtv, &result.SborPetliRam, &result.SborYstPorog, &result.Naveshiv, &result.Zashiv, &result.OpresRam, &result.SborPetliStv,
		&result.NastrStanok, &result.NastrPbx, &result.MehObrabPzr, &result.RabotaPbx, &result.FrezerNastr, &result.FrezerPorogSborka, &result.FrezerYstShtyp, &result.OpresNastr,
		&result.Opres, &result.PodgDerjShetki, &result.YstPorogYplDr, &result.NaveshivStv, &result.YstZapoln, &result.ImpostNapil, &result.ImpostFrezer, &result.ImpostSverlo,
		&result.ImpostYst, &result.ImpostShtift, &result.YplFalc, &result.NapilNalich, &result.NapilRam, &result.NapilStv, &result.KontrSbork, &result.SverlRam, &result.SverlZink,
		&result.ZashitPl, &result.SborRam, &result.NapilYstKrishStv, &result.NapilYstKrishRam, &result.SborPetli, &result.YstPtliRamStv, &result.RezkaPlast, &result.Brysok,
		&result.IzgPritv, &result.ObrPritv, &result.YstPritv, &result.ObrabotkaAll, &result.YstanPlnPetli, &result.YstFetr, &result.Rezina, &result.FrezerShping,
		&result.Gl, &result.Fortochka, &result.Upak, &result.OrderId, &result.TotalTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for id: %d", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	// Маппинг в OrderData (через утилиту)
	return utils.MapToOrderDataDoor(&result), nil
}

func (s *Storage) SaveVitrag(result storage.DemResultVitraj) (int, error) {
	const op = "storage.mysql.sql.SaveVitrag"

	// Шаг 1: Получить или создать заказ по order_num
	orderID, err := s.getOrCreateOrder(result.OrderNum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get/create order-norm: %w", op, err)
	}

	result.OrderID = orderID

	stmt := `INSERT INTO dem_test_golang_result_vitraj (order_num, name, count, profil, podgot_oboryd, napil_stk_do3m, napil_stk_bol3m, napil_stk_bol5m, pzr, nastr_pbx,
        drenaj_otv, zashelk_obr, priset_pr_obr, ypl_yst, fetr_yst,steklo_yst, razmet_otv, sverlo_otv, naklad_yst, panel,naprav, ydlin, upak, napil_rigel_do1m, napil_rigel_bol1m,
        nakld_napil, napil_krish_st, napil_krish_rg, kompl_krish, napil_adapt,frezer_rigel, fukel_yst, stoiki_pbx, zamok_yst, zamok_yst_bolt,rigel_sverlo_zamok, naklad_sverlo, 
        ypl_yst_rigel, ypl_yst_nakld, ypl_yst_stoik,adapt_yst_stoik, adapt_yst_rigel, termovst_yst_stoiki, termovst_yst_rigel, sbor_nog, birki, upak_stoik, upak_pet, upak_rigel,
        upak_rigel_2m, upak_krish, upak_nakld, upak_ypl_nog, upak_kronsht, vinos_gotov_izd,yst_stik_zakld, obn_ysov_do4m, obn_ysov_bol4m, otm_rezin, obn_nastr,
        tryba_prof, kompl_shtapik, otv_vo, nastr_stanok_rigel, frezer_stoiki_pr, nastr_stanok_stoiki_1, frezer_stoiki_ygl, nastr_stanok_stoiki_2, zashit_pl, yst_pritv,
        yst_kapel, sbor_sekci, germet_ypl, upak_soed, upak_shtapik, upak_sekcii, obrezanie, vinos_sekcii, ygol_korob, pbx_vo, clear_vo, yst_komptr_st, yst_komptr_rg, upak_ygl, napil_ygl,
        narez_termovst, narez_kompensr, upak_termovst, rigel_sverlo_zamok_3m, rigel_nsverlo_zamok,obj_ysov_termvst, yst_kps, napil_stk_bol6m, napil_shtapik, rez_profil,
        rez_sekcii, part_sb_seck, part_sb_mark_seck, part_sb_stoek, yst_formr, order_id, total_time
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
              ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
              ?, ?)`

	res, err := s.db.Exec(stmt,
		result.OrderNum, result.Name, result.Count, result.Profil, result.PodgotOboryd, result.NapilStkDo3m, result.NapilStkBol3m,
		result.NapilStkBol5m, result.PZR, result.NastrPbx, result.DrenajOtv, result.ZashelkObr, result.PrisetPrObr, result.YplYst,
		result.FetrYst, result.StekloYst, result.RazmetOtv, result.SverloOtv, result.NakladYst, result.Panel, result.Naprav,
		result.Ydlin, result.Upak, result.NapilRigelDo1m, result.NapilRigelBol1m, result.NakldNapil, result.NapilKriskhSt,
		result.NapilKriskhRg, result.KomplKriskh, result.NapilAdapt, result.FrezerRigel, result.FukelYst, result.StoikiPbx,
		result.ZamokYst, result.ZamokYstBolt, result.RigelSverloZamok, result.NakladSverlo, result.YplYstRigel,
		result.YplYstNakld, result.YplYstStoik, result.AdaptYstStoik, result.AdaptYstRigel, result.TermovstYstStoiki,
		result.TermovstYstRigel, result.SborNog, result.Birki, result.UpakStoik, result.UpakPet, result.UpakRigel,
		result.UpakRigel2m, result.UpakKriskh, result.UpakNakld, result.UpakYplNog, result.UpakKronsht, result.VinosGotovIzd,
		result.YstStikZakld, result.ObnYsovDo4m, result.ObnYsovBol4m, result.OtmRezin, result.ObnNastr, result.TrybaProf,
		result.KomplShtapik, result.OtvVO, result.NastrStanokRigel, result.FrezerStoikiPr, result.NastrStanokStoiki1,
		result.FrezerStoikiYgl, result.NastrStanokStoiki2, result.ZashitPl, result.YstPritv, result.YstKapel, result.SborSekci,
		result.GermetYpl, result.UpakSoed, result.UpakShtapik, result.UpakSekcii, result.Obrezanie, result.VinosSekcii,
		result.YgolKorob, result.PbxVO, result.ClearVO, result.YstKomptrSt, result.YstKomptrRg, result.UpakYgl, result.NapilYgl,
		result.NarezTermovst, result.NarezKompensr, result.UpakTermovst, result.RigelSverloZamok3m, result.RigelNsverloZamok,
		result.ObjYsovTermvst, result.YstKPS, result.NapilStkBol6m, result.NapilShtapik, result.RezProfil, result.RezSekcii,
		result.PartSbSeck, result.PartSbMarkSeck, result.PartSbStoek, result.YstFormr, result.OrderID, result.TotalTime,
	)

	if err != nil {
		return 0, fmt.Errorf("%s: failed to execute insert query: %w", op, err)
	}

	// Получаем ID вставленной записи
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return int(id), nil
}

func (s *Storage) GetVitraj(id int) (*storage.OrderData, error) {
	const op = "storage.mysql.sql.GetVitraj"

	query := `SELECT 
        id, order_num, name, count, profil, podgot_oboryd, napil_stk_do3m, napil_stk_bol3m, napil_stk_bol5m, pzr, nastr_pbx,
        drenaj_otv, zashelk_obr, priset_pr_obr, ypl_yst, fetr_yst, steklo_yst, razmet_otv, sverlo_otv, naklad_yst, panel,
        naprav, ydlin, upak, napil_rigel_do1m, napil_rigel_bol1m, nakld_napil, napil_krish_st, napil_krish_rg, kompl_krish,
        napil_adapt, frezer_rigel, fukel_yst, stoiki_pbx, zamok_yst, zamok_yst_bolt, rigel_sverlo_zamok, naklad_sverlo,
        ypl_yst_rigel, ypl_yst_nakld, ypl_yst_stoik, adapt_yst_stoik, adapt_yst_rigel, termovst_yst_stoiki, termovst_yst_rigel,
        sbor_nog, birki, upak_stoik, upak_pet, upak_rigel, upak_rigel_2m, upak_krish, upak_nakld, upak_ypl_nog, upak_kronsht,
        vinos_gotov_izd, yst_stik_zakld, obn_ysov_do4m, obn_ysov_bol4m, otm_rezin, obn_nastr, tryba_prof, kompl_shtapik,
        otv_vo, nastr_stanok_rigel, frezer_stoiki_pr, nastr_stanok_stoiki_1, frezer_stoiki_ygl, nastr_stanok_stoiki_2,
        zashit_pl, yst_pritv, yst_kapel, sbor_sekci, germet_ypl, upak_soed, upak_shtapik, upak_sekcii, obrezanie,
        vinos_sekcii, ygol_korob, pbx_vo, clear_vo, yst_komptr_st, yst_komptr_rg, upak_ygl, napil_ygl, narez_termovst,
        narez_kompensr, upak_termovst, rigel_sverlo_zamok_3m, rigel_nsverlo_zamok, obj_ysov_termvst, yst_kps,
        napil_stk_bol6m, napil_shtapik, rez_profil, rez_sekcii, part_sb_seck, part_sb_mark_seck, part_sb_stoek, yst_formr,
        order_id, total_time
    FROM dem_test_golang_result_vitraj
    WHERE id = ?`

	var result storage.DemResultVitraj

	err := s.db.QueryRow(query, id).Scan(&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.Profil, &result.PodgotOboryd, &result.NapilStkDo3m,
		&result.NapilStkBol3m, &result.NapilStkBol5m, &result.PZR, &result.NastrPbx, &result.DrenajOtv, &result.ZashelkObr, &result.PrisetPrObr, &result.YplYst, &result.FetrYst,
		&result.StekloYst, &result.RazmetOtv, &result.SverloOtv, &result.NakladYst, &result.Panel, &result.Naprav, &result.Ydlin, &result.Upak, &result.NapilRigelDo1m,
		&result.NapilRigelBol1m, &result.NakldNapil, &result.NapilKriskhSt, &result.NapilKriskhRg, &result.KomplKriskh, &result.NapilAdapt, &result.FrezerRigel, &result.FukelYst,
		&result.StoikiPbx, &result.ZamokYst, &result.ZamokYstBolt, &result.RigelSverloZamok, &result.NakladSverlo, &result.YplYstRigel, &result.YplYstNakld, &result.YplYstStoik,
		&result.AdaptYstStoik, &result.AdaptYstRigel, &result.TermovstYstStoiki, &result.TermovstYstRigel, &result.SborNog, &result.Birki, &result.UpakStoik, &result.UpakPet,
		&result.UpakRigel, &result.UpakRigel2m, &result.UpakKriskh, &result.UpakNakld, &result.UpakYplNog, &result.UpakKronsht, &result.VinosGotovIzd, &result.YstStikZakld,
		&result.ObnYsovDo4m, &result.ObnYsovBol4m, &result.OtmRezin, &result.ObnNastr, &result.TrybaProf, &result.KomplShtapik, &result.OtvVO, &result.NastrStanokRigel,
		&result.FrezerStoikiPr, &result.NastrStanokStoiki1, &result.FrezerStoikiYgl, &result.NastrStanokStoiki2, &result.ZashitPl, &result.YstPritv, &result.YstKapel, &result.SborSekci,
		&result.GermetYpl, &result.UpakSoed, &result.UpakShtapik, &result.UpakSekcii, &result.Obrezanie, &result.VinosSekcii, &result.YgolKorob, &result.PbxVO, &result.ClearVO,
		&result.YstKomptrSt, &result.YstKomptrRg, &result.UpakYgl, &result.NapilYgl, &result.NarezTermovst, &result.NarezKompensr, &result.UpakTermovst, &result.RigelSverloZamok3m,
		&result.RigelNsverloZamok, &result.ObjYsovTermvst, &result.YstKPS, &result.NapilStkBol6m, &result.NapilShtapik, &result.RezProfil, &result.RezSekcii, &result.PartSbSeck,
		&result.PartSbMarkSeck, &result.PartSbStoek, &result.YstFormr, &result.OrderID, &result.TotalTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for id: %d", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	// возвращает читабельный вид через доп функцию в пакете utils
	return utils.MapToOrderDataVitraj(&result), nil
}

func (s *Storage) SaveLoggia(result storage.DemResultLoggia) (int, error) {
	const op = "storage.mysql.sql.SaveLoggia"

	// Шаг 1: Получить или создать заказ по order_num
	orderID, err := s.getOrCreateOrder(result.OrderNum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get/create order-norm: %w", op, err)
	}

	result.OrderID = orderID

	// Подготовка SQL-запроса INSERT
	stmt := `INSERT INTO dem_test_golang_result_loggia (
        order_num, name, count, profil, podgot_oboryd, napil_stk_do3m, napil_rigel_do1m, napil_rigel_bol1m, napil_shtapik, kompl_shtapik,napil_adapt_tr, rigel_frezer, 
        frezer_rigel_zamok, zamok_yst, rez_pgm_st, rez_pgm_rg, frez_vo, part_sb_sekci, yst_pritv, yst_formir,birki, upak_ram, upak_stoik, upak_rigel, upak_rigel_2m,
        upak_shtapik, upak_adapt_tr, upak_ypl_nog, vinos_gotov_izd, napil_pr_virav, napil_ram, napil_stv, napil_pritv, napil_soed, frez_ram, frez_stv, frez_pritv, 
        clear_sverl_stv, clear_sverl_ysil_stv, kraska, yst_rolik, yst_zashel, rezina, sbor_stv, sbor_pritv,podg_kompl, shtamp_st_stv, podg_rezin, sbor_ypl, yst_zaklep,
        order_id, total_time
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Выполнение запроса с передачей всех значений
	res, err := s.db.Exec(stmt, result.OrderNum, result.Name, result.Count, result.Profil, result.PodgotOboryd, result.NapilStkDo3m, result.NapilRigelDo1m,
		result.NapilRigelBol1m, result.NapilShtapik, result.KomplShtapik, result.NapilAdaptTr, result.RigelFrezer, result.FrezerRigelZamok, result.ZamokYst, result.RezPgmSt,
		result.RezPgmRg, result.FrezVo, result.PartSbSekci, result.YstPritv, result.YstFormir, result.Birki, result.UpakRam, result.UpakStoik, result.UpakRigel, result.UpakRigel2m,
		result.UpakShtapik, result.UpakAdaptTr, result.UpakYplNog, result.VinosGotovIzd, result.NapilPrVirav, result.NapilRam, result.NapilStv, result.NapilPritv, result.NapilSoed,
		result.FrezRam, result.FrezStv, result.FrezPritv, result.ClearSverlStv, result.ClearSverlYsilStv, result.Kraska, result.YstRolik, result.YstZashel, result.Rezina,
		result.SborStv, result.SborPritv, result.PodgKompl, result.ShtampStStv, result.PodgRezin, result.SborYpl, result.YstZaklep, result.OrderID, result.TotalTime,
	)

	if err != nil {
		return 0, fmt.Errorf("%s: failed to execute insert query: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return int(id), nil
}

func (s *Storage) GetLoggia(id int) (*storage.OrderData, error) {
	const op = "storage.mysql.sql.GetLoggia"

	query := `SELECT id, order_num, name, count, profil, podgot_oboryd,napil_stk_do3m, napil_rigel_do1m, napil_rigel_bol1m, napil_shtapik, kompl_shtapik,
        napil_adapt_tr, rigel_frezer, frezer_rigel_zamok, zamok_yst, rez_pgm_st, rez_pgm_rg, frez_vo, part_sb_sekci, yst_pritv, yst_formir,
        birki, upak_ram, upak_stoik, upak_rigel, upak_rigel_2m, upak_shtapik, upak_adapt_tr, upak_ypl_nog, vinos_gotov_izd, napil_pr_virav,
        napil_ram, napil_stv, napil_pritv, napil_soed, frez_ram, frez_stv, frez_pritv, clear_sverl_stv, clear_sverl_ysil_stv, kraska,yst_rolik, yst_zashel, rezina, 
        sbor_stv, sbor_pritv, podg_kompl, shtamp_st_stv, podg_rezin, sbor_ypl, yst_zaklep, order_id, total_time
    FROM dem_test_golang_result_loggia
    WHERE id = ?`

	var result storage.DemResultLoggia

	err := s.db.QueryRow(query, id).Scan(&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.Profil, &result.PodgotOboryd, &result.NapilStkDo3m,
		&result.NapilRigelDo1m, &result.NapilRigelBol1m, &result.NapilShtapik, &result.KomplShtapik, &result.NapilAdaptTr, &result.RigelFrezer, &result.FrezerRigelZamok,
		&result.ZamokYst, &result.RezPgmSt, &result.RezPgmRg, &result.FrezVo, &result.PartSbSekci, &result.YstPritv, &result.YstFormir, &result.Birki, &result.UpakRam,
		&result.UpakStoik, &result.UpakRigel, &result.UpakRigel2m, &result.UpakShtapik, &result.UpakAdaptTr, &result.UpakYplNog, &result.VinosGotovIzd, &result.NapilPrVirav,
		&result.NapilRam, &result.NapilStv, &result.NapilPritv, &result.NapilSoed, &result.FrezRam, &result.FrezStv, &result.FrezPritv, &result.ClearSverlStv,
		&result.ClearSverlYsilStv, &result.Kraska, &result.YstRolik, &result.YstZashel, &result.Rezina, &result.SborStv, &result.SborPritv, &result.PodgKompl,
		&result.ShtampStStv, &result.PodgRezin, &result.SborYpl, &result.YstZaklep, &result.OrderID, &result.TotalTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for id: %d", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	// Преобразуем в универсальную структуру OrderData через утилиту
	return utils.MapToOrderDataLoggia(&result), nil
}

func (s *Storage) getOrCreateOrder(orderNum string) (int, error) {
	var orderID int

	// 1. Попробуем найти существующий заказ
	err := s.db.QueryRow(`SELECT id FROM dem_test_golang_all_orders WHERE order_num = ?`, orderNum).Scan(&orderID)

	if err == nil {
		return orderID, nil // Нашли — возвращаем id
	}

	if err == sql.ErrNoRows {
		// 2. Заказа нет — создаём
		result, err := s.db.Exec(`INSERT INTO dem_test_golang_all_orders (order_num) VALUES (?)`, orderNum)
		if err != nil {
			return 0, err
		}
		id, _ := result.LastInsertId()
		return int(id), nil
	}

	return 0, err
}
