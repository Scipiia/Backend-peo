package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"vue-golang/internal/storage"
)

func (s *Storage) GetOrdersMonth(year int, month int, search string) ([]*storage.Order, error) {
	const op = "storage.order-dem-details.GetOrdersMonth.sql"

	var stmt string
	var args []interface{}

	// Если есть поиск — ищем по всем заказам (глобальный поиск)
	if search != "" {
		stmt = `
			SELECT id, order_num, creator, customer, dop_info, ms_note 
			FROM dem_ready 
			WHERE order_num LIKE ?
		`
		args = append(args, "%"+search+"%")
	} else {
		// Иначе фильтруем по месяцу
		startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		endOfMonth := startOfMonth.AddDate(0, 1, 0)
		startUnix := startOfMonth.Unix()
		endUnix := endOfMonth.Unix()

		stmt = `
			SELECT id, order_num, creator, customer, dop_info, ms_note 
			FROM dem_ready 
			WHERE CAST(creation_date AS UNSIGNED) >= ? 
			  AND CAST(creation_date AS UNSIGNED) < ?
		`
		args = []interface{}{startUnix, endUnix}
	}

	// Дополнительно можно всегда искать только "готовые" заказы (если нужно)
	stmt += " AND order_num LIKE '%Q6%'" // например

	rows, err := s.db.Query(stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var orders []*storage.Order
	for rows.Next() {
		var order storage.Order
		var msNote sql.NullString

		err := rows.Scan(&order.ID, &order.OrderNum, &order.Creator, &order.Customer, &order.DopInfo, &msNote)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if msNote.Valid {
			order.MsNote = msNote.String
		} else {
			order.MsNote = ""
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
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

	//TODO основная структура
	details := &storage.ResultOrderDetails{}

	stmtDemOrders := "SELECT id, order_num, creator, customer, dop_info, ms_note FROM dem_ready WHERE id = ?"

	var msNote sql.NullString // Используем NullString для обработки NULL
	details.Order = &storage.Order{}
	err = tx.QueryRow(stmtDemOrders, id).Scan(&details.Order.ID, &details.Order.OrderNum, &details.Order.Creator, &details.Order.Customer, &details.Order.DopInfo, &msNote)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no rows found for query dem orders: %w", op, err)
		}
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
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
       	 	pl.color,
        	COALESCE(SUM(pl.sqr), 0) AS sqr
    	FROM dem_price p
    	LEFT JOIN dem_images i ON i.im_ordername = p.numorders AND i.im_orderpos = p.position
    	LEFT JOIN dem_plan pl ON pl.idorder = ? AND CAST(pl.x AS UNSIGNED) = p.position
    	WHERE p.numorders LIKE ?
    	GROUP BY p.position, p.creator, p.name_position, p.kol_vo, i.im_image, pl.color
    	ORDER BY 1`

	//TODO тут будет массив данных
	details.OrderDemPrice = []*storage.OrderDemPrice{}

	rows, err := tx.Query(stmtDemPrice, id, details.Order.OrderNum)
	if err != nil {
		return nil, fmt.Errorf("%s:query failed no rows in dem price %w", op, err)
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
			&price.Color,
			&price.Sqr,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row dem price: %w", op, err)
		}

		details.OrderDemPrice = append(details.OrderDemPrice, price)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return details, nil
}
