package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"
	"vue-golang/internal/storage"
)

func (s *Storage) GetOrdersMonth(year int, month int) ([]*storage.Order, error) {
	const op = "storage.order-dem-details.GetOrdersMonth.sql"

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

func (s *Storage) GetOrderDetails(id int) (*storage.ResultOrderDetails, error) {
	const op = "storage.order-dem-details.GetOrderDetails.sql"

	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	slog.Any("IDDDDDDDDDDDDD", id)

	//TODO основная структура
	details := &storage.ResultOrderDetails{}

	stmtDemOrders := "SELECT id, order_num, creator, customer, dop_info, ms_note FROM dem_ready WHERE id = ?"

	//order-dem := &storage.Order{}
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

	stmtDemPrice := `
    	SELECT 
        	CAST(p.position AS UNSIGNED),
        	p.creator,
        	p.name_position,
        	p.kol_vo,
       	 	i.im_image,
        	COALESCE(SUM(pl.sqr), 0) AS sqr  -- суммарная площадь по позиции
    	FROM dem_price p
    	LEFT JOIN dem_images i ON i.im_ordername = p.numorders AND i.im_orderpos = p.position
    	LEFT JOIN dem_plan pl ON pl.idorder = ? AND CAST(pl.x AS UNSIGNED) = p.position
    	WHERE p.numorders LIKE ?
    	GROUP BY p.position, p.creator, p.name_position, p.kol_vo, i.im_image
    	ORDER BY 1`

	//TODO тут будет массив данных
	details.OrderDemPrice = []*storage.OrderDemPrice{}

	rows, err := tx.Query(stmtDemPrice, id, details.Order.OrderNum)
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
			&price.Sqr,
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
