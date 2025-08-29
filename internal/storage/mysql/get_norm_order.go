package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetNormOrder(id int64) (*storage.GetOrderDetails, error) {
	const op = "storage.mysql.GetNormOrder"

	stmt := "SELECT order_num, name, count, total_time, created_at, updated_at, type FROM product_instances WHERE id = ?"

	stmt1 := "SELECT operation_name, operation_label, count, value, minutes FROM operation_values WHERE product_id = ?"

	var res storage.GetOrderDetails

	fmt.Println("SDSDSDSDSDDS", id)

	err := s.db.QueryRow(stmt, id).Scan(&res.OrderNum, &res.Name, &res.Count, &res.TotalTime, &res.CreatedAT, &res.UpdatedAT, &res.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: нормировка не найдена: %w", op, err)
		}
		return nil, fmt.Errorf("%s: ошибка запроса: %w", op, err)
	}

	rows, err := s.db.Query(stmt1, id)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка получения операций: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var op storage.NormOperation
		err := rows.Scan(&op.Name, &op.Label, &op.Count, &op.Value, &op.Minutes)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования операции: %w", op, err)
		}
		res.Operations = append(res.Operations, op)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка при итерации: %w", op, err)
	}

	return &res, nil
}

func (s *Storage) GetNormOrders(orderNum, orderType string) ([]storage.GetOrderDetails, error) {
	const op = "storage.mysql.GetNormOrders"

	stmt := `SELECT id, order_num, name, count, total_time, created_at, type FROM product_instances WHERE 1=1 AND (?='' OR order_num LIKE CONCAT('%', ?, '%')) AND (? = '' OR type = ?) ORDER BY created_at DESC`

	rows, err := s.db.Query(stmt, orderNum, orderNum, orderType, orderType)
	if err != nil {
		return nil, fmt.Errorf("%s: выполнение запроса: %w", op, err)
	}
	defer rows.Close()

	var items []storage.GetOrderDetails
	for rows.Next() {
		var item storage.GetOrderDetails
		err = rows.Scan(
			&item.ID,
			&item.OrderNum,
			&item.Name,
			&item.Count,
			&item.TotalTime,
			&item.CreatedAT,
			&item.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: сканирование: %w", op, err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка итерации: %w", op, err)
	}

	return items, nil
}
