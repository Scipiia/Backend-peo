package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"vue-golang/internal/storage"
)

func (s *Storage) GetOrdersMonth(year int, month int) ([]*storage.Order, error) {
	const op = "storage.order-details.sql"

	// Определяем начало и конец месяца
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	// Преобразуем даты в Unix-время
	startUnix := startOfMonth.Unix()
	endUnix := endOfMonth.Unix()

	// SQL-запрос для получения заказов за месяц
	stmt := `SELECT id, order_num, creator, customer, dop_info, ms_note 
        FROM dem_ready 
        WHERE CAST(creation_date AS UNSIGNED) >= ? AND CAST(creation_date AS UNSIGNED) <= ?
    `

	// Выполняем запрос
	rows, err := s.db.Query(stmt, startUnix, endUnix)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	// Собираем результаты
	var orders []*storage.Order

	for rows.Next() {
		order := &storage.Order{}
		var msNote sql.NullString // Используем NullString для обработки NULL

		err := rows.Scan(&order.ID, &order.OrderNum, &order.Creator, &order.Customer, &order.DopInfo, &msNote)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		// Обрабатываем msNote
		if msNote.Valid {
			order.MsNote = msNote.String // Если значение не NULL, сохраняем строку
		} else {
			order.MsNote = "" // Если значение NULL, сохраняем пустую строку
		}

		orders = append(orders, order)
	}

	log.Println(orders, "PIZDPA APIKDOSDSD")
	// Проверяем ошибки после цикла
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}

func (s *Storage) GetReadyOrders() ([]*storage.DemResult, error) {
	const op = "storage.order-details.sql"

	stmt := `
    SELECT 
        id, order_num, name, count, nast_napil, napil, napil_krishek, napil_impost,
        soedinitel, promej_sborka, impost_sverlo, impost_frezer, impost_sborka,
        opres_nastr, opresovka, ystan_yplotn, zashivka, profil
    FROM dem_test_golang_rezult_glyhar`

	// Выполняем запрос
	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	// Собираем результаты
	var res []*storage.DemResult

	for rows.Next() {
		result := &storage.DemResult{}

		err := rows.Scan(&result.ID, &result.OrderNum, &result.Name, &result.Count, &result.PodgotovOboryd, &result.NapilKontr, &result.NapilKrishek,
			&result.NapilImpost, &result.Soedinitel, &result.PromejSborka, &result.ImpostSverlovka, &result.ImpostFrezerovka,
			&result.ImpostSborka, &result.OpresNastr, &result.Opresovka, &result.YstanYplotnitel, &result.Zashivka, &result.Profil)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		res = append(res, result)
	}

	return res, nil
}
