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
