package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetOrderDetails(id int) (*storage.ResultOrderDetails, error) {
	const op = "storage.order-details.sql"

	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	//TODO основная структура
	details := &storage.ResultOrderDetails{}

	stmtDemOrders := "SELECT id, order_num, creator, customer, dop_info, ms_note FROM dem_ready WHERE id = ?"

	//order := &storage.Order{}
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
func (s *Storage) SaveOrder(result storage.DemResult) (int, error) {
	const op = "storage.mysql.sql.saveOrder"
	stmt := `INSERT INTO dem_test_golang_rezult_glyhar (order_num, name, count,  nast_napil, napil, napil_krishek, napil_impost,
	soedinitel, promej_sborka, impost_sverlo, impost_frezer,impost_sborka, opres_nastr, opresovka, ystan_yplotn, zashivka, profil,
	napil_stoiki_do3m, napil_stoiki_bol3m, napil_rigel_do1m, napil_rigel_bol1m, sverl_rigel_zamok, ystan_zamkov, napil_shtapik, ypakovka,
    frezer_rigel, obrabot_ram, hands_sborka, frezer_nastr, shtiftovka, ystanovka_zapoln)
	VALUES (?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	exec, err := s.db.Exec(stmt, result.OrderNum, result.Name, result.Count, result.PodgotovOboryd, result.NapilKontr, result.NapilKrishek,
		result.NapilImpost, result.Soedinitel, result.PromejSborka, result.ImpostSverlovka, result.ImpostFrezerovka, result.ImpostSborka,
		result.OpresNastr, result.Opresovka, result.YstanYplotnitel, result.Zashivka, result.Profil, result.NapilStoikiDo3m, result.NapilStoikiBol3m,
		result.NapilRigelDo1m, result.NapilRigelBol1m, result.SverlRigelZamok, result.YstanZamkov, result.NapilShtapik, result.Ypakovka, result.FrezerRigel,
		result.ObrabotRam, result.HandsSborka, result.FrezerNastr, result.Shtiftovka, result.YstanovkaZapoln)
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

	stmt := `SELECT id, order_num, name, count,  nast_napil, napil, napil_krishek, napil_impost,
		soedinitel, promej_sborka, impost_sverlo, impost_frezer,impost_sborka, opres_nastr, opresovka, ystan_yplotn, zashivka, profil,
		napil_stoiki_do3m, napil_stoiki_bol3m, napil_rigel_do1m, napil_rigel_bol1m, sverl_rigel_zamok, ystan_zamkov, napil_shtapik, ypakovka,
    	frezer_rigel, obrabot_ram, hands_sborka, frezer_nastr, shtiftovka, ystanovka_zapoln FROM dem_test_golang_rezult_glyhar WHERE id LIKE ?`

	var result storage.DemResult
	err := s.db.QueryRow(stmt, id).Scan(
		&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.PodgotovOboryd, &result.NapilKontr, &result.NapilKrishek, &result.NapilImpost,
		&result.Soedinitel, &result.PromejSborka, &result.ImpostSverlovka, &result.ImpostFrezerovka, &result.ImpostSborka, &result.OpresNastr,
		&result.Opresovka, &result.YstanYplotnitel, &result.Zashivka, &result.Profil, &result.NapilStoikiDo3m, &result.NapilStoikiBol3m,
		&result.NapilRigelDo1m, &result.NapilRigelBol1m, &result.SverlRigelZamok, &result.YstanZamkov, &result.NapilShtapik, &result.Ypakovka,
		&result.FrezerRigel, &result.ObrabotRam, &result.HandsSborka, &result.FrezerNastr, &result.Shtiftovka, &result.YstanovkaZapoln,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no data found for order_num: %s", op, id)
		}
		return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
	}

	operations := []storage.Operation{}

	// Добавляем только те, где значение > 0
	if result.PodgotovOboryd > 0 {
		operations = append(operations, storage.Operation{"podgotov_oboryd", "Настройка оборудования для напиловки", result.PodgotovOboryd})
	}
	if result.NapilKontr > 0 {
		operations = append(operations, storage.Operation{"napil_kontr", "Напиловка контура", result.NapilKontr})
	}
	if result.NapilKrishek > 0 {
		operations = append(operations, storage.Operation{"napil_krishek", "Напиловка крышек", result.NapilKrishek})
	}
	if result.NapilImpost > 0 {
		operations = append(operations, storage.Operation{"napil_impost", "Напиловка импоста", result.NapilImpost})
	}
	if result.Soedinitel > 0 {
		operations = append(operations, storage.Operation{"soedinitel", "Соединитель", result.Soedinitel})
	}
	if result.PromejSborka > 0 {
		operations = append(operations, storage.Operation{"promej_sborka", "Промежуточная сборка", result.PromejSborka})
	}
	if result.ImpostSverlovka > 0 {
		operations = append(operations, storage.Operation{"impost_sverlovka", "Импост (сверловка)", result.ImpostSverlovka})
	}
	if result.ImpostFrezerovka > 0 {
		operations = append(operations, storage.Operation{"impost_frezerovka", "Импост (фрезеровка)", result.ImpostFrezerovka})
	}
	if result.ImpostSborka > 0 {
		operations = append(operations, storage.Operation{"impost_sborka", "Импост (сборка)", result.ImpostSborka})
	}
	if result.OpresNastr > 0 {
		operations = append(operations, storage.Operation{"opres_nastr", "Опресовка (настройка)", result.OpresNastr})
	}
	if result.Opresovka > 0 {
		operations = append(operations, storage.Operation{"opresovka", "Опресовка", result.Opresovka})
	}
	if result.YstanYplotnitel > 0 {
		operations = append(operations, storage.Operation{"ystan_yplotnitel", "Установка уплотнителей, штифтовка", result.YstanYplotnitel})
	}
	if result.Zashivka > 0 {
		operations = append(operations, storage.Operation{"zashivka", "Зашивка", result.Zashivka})
	}
	if result.NapilStoikiDo3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_do3m", "Напиловка стойки до 3м", result.NapilStoikiDo3m})
	}
	if result.NapilStoikiBol3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_bol3m", "Напиловка стойки больше 3м", result.NapilStoikiBol3m})
	}
	if result.NapilRigelDo1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_do1m", "Напиловка ригеля до 1м", result.NapilRigelDo1m})
	}
	if result.NapilRigelBol1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_bol1m", "Напиловка ригеля больше 1м", result.NapilRigelBol1m})
	}
	if result.SverlRigelZamok > 0 {
		operations = append(operations, storage.Operation{"sverl_rigel_zamok", "Сверловка ригеля под замок", result.SverlRigelZamok})
	}
	if result.YstanZamkov > 0 {
		operations = append(operations, storage.Operation{"ystan_zamkov", "Установка замков", result.YstanZamkov})
	}
	if result.NapilShtapik > 0 {
		operations = append(operations, storage.Operation{"napil_shtapik", "Напиловка штапика", result.NapilShtapik})
	}
	if result.Ypakovka > 0 {
		operations = append(operations, storage.Operation{"ypakovka", "Упаковка", result.Ypakovka})
	}
	if result.FrezerRigel > 0 {
		operations = append(operations, storage.Operation{"frezer_rigel", "Фрезеровка ригеля", result.FrezerRigel})
	}
	if result.ObrabotRam > 0 {
		operations = append(operations, storage.Operation{"obrabot_ram", "Обработка рам", result.ObrabotRam})
	}
	if result.HandsSborka > 0 {
		operations = append(operations, storage.Operation{"hands_sborka", "Ручная сборка", result.HandsSborka})
	}
	if result.FrezerNastr > 0 {
		operations = append(operations, storage.Operation{"frezer_nastr", "Фрезеровка (настройка)", result.FrezerNastr})
	}
	if result.Shtiftovka > 0 {
		operations = append(operations, storage.Operation{"shtiftovka", "Штифтовка", result.Shtiftovka})
	}
	if result.YstanovkaZapoln > 0 {
		operations = append(operations, storage.Operation{"ystanovka_zapoln", "Установка заполнения", result.YstanovkaZapoln})
	}

	resultFinish := &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Profil:     result.Profil,
		Operations: operations,
	}

	return resultFinish, nil
}
