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

func (s *Storage) GetNormOrderIdSub(id int64) ([]*storage.GetOrderDetails, error) {
	const op = "storage.mysql.GetNormOrder"

	stmt := `
		SELECT 
			id, name, count, total_time, created_at, updated_at, 
			type, part_type, parent_assembly, parent_product_id, order_num
		FROM product_instances 
		WHERE id = ? OR parent_product_id = ?
		ORDER BY 
			CASE WHEN part_type = 'main' THEN 0 ELSE 1 END, 
			id
	`

	stmtOps := `SELECT operation_name, operation_label, count, value, minutes FROM operation_values WHERE product_id = ?`

	rows, err := s.db.Query(stmt, id, id)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка выполнения: %w", op, err)
	}
	defer rows.Close()

	var results []*storage.GetOrderDetails

	for rows.Next() {
		var detail storage.GetOrderDetails
		var parentAssembly sql.NullString

		err := rows.Scan(
			&detail.ID,
			&detail.Name,
			&detail.Count,
			&detail.TotalTime,
			&detail.CreatedAT,
			&detail.UpdatedAT,
			&detail.Type,
			&detail.PartType,
			&parentAssembly,
			&detail.ParentProductID,
			&detail.OrderNum,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования: %w", op, err)
		}
		if parentAssembly.Valid {
			detail.ParentAssembly = parentAssembly.String
		} else {
			detail.ParentAssembly = ""
		}

		// Операции
		opsRows, err := s.db.Query(stmtOps, detail.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка операций для id=%d: %w", op, detail.ID, err)
		}

		for opsRows.Next() {
			var op storage.NormOperation
			err := opsRows.Scan(&op.Name, &op.Label, &op.Count, &op.Value, &op.Minutes)
			if err != nil {
				opsRows.Close()
				return nil, fmt.Errorf("%s: ошибка сканирования операции: %w", op, err)
			}
			detail.Operations = append(detail.Operations, op)
		}
		opsRows.Close()

		results = append(results, &detail)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: итерация строк: %w", op, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("%s: наряд с id=%d не найден", op, id)
	}

	return results, nil
}

func (s *Storage) GetNormOrdersByOrderNum(orderNum string) ([]*storage.GetOrderDetails, error) {
	const op = "storage.mysql.GetNormOrdersByOrderNum"

	// SQL: получаем все наряды по order_num
	stmt := `
		SELECT 
			id, name, count, total_time, created_at, updated_at, type, part_type, parent_assembly, parent_product_id
		FROM product_instances 
		WHERE order_num = ? 
		ORDER BY 
			CASE WHEN part_type = 'main' THEN 0 ELSE 1 END, 
			id
	`

	// Операции по product_id
	stmtOps := `
		SELECT operation_name, operation_label, count, value, minutes 
		FROM operation_values 
		WHERE product_id = ?
	`

	rows, err := s.db.Query(stmt, orderNum)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка выполнения запроса: %w", op, err)
	}
	defer rows.Close()

	var results []*storage.GetOrderDetails

	for rows.Next() {
		var detail storage.GetOrderDetails
		var parentAssembly sql.NullString // parent_assembly может быть NULL

		err := rows.Scan(
			&detail.ID,
			&detail.Name,
			&detail.Count,
			&detail.TotalTime,
			&detail.CreatedAT,
			&detail.UpdatedAT,
			&detail.Type,
			&detail.PartType,
			&parentAssembly,
			&detail.ParentProductID,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования наряда: %w", op, err)
		}

		// Обработка parent_assembly
		if parentAssembly.Valid {
			detail.ParentAssembly = parentAssembly.String
		} else {
			detail.ParentAssembly = ""
		}

		// Получаем операции для этого наряда
		opsRows, err := s.db.Query(stmtOps, detail.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка получения операций для product_id=%d: %w", op, detail.ID, err)
		}

		for opsRows.Next() {
			var op storage.NormOperation
			if err := opsRows.Scan(&op.Name, &op.Label, &op.Count, &op.Value, &op.Minutes); err != nil {
				opsRows.Close()
				return nil, fmt.Errorf("%s: ошибка сканирования операции: %w", op, err)
			}
			detail.Operations = append(detail.Operations, op)
		}
		opsRows.Close()

		if err := opsRows.Err(); err != nil {
			return nil, fmt.Errorf("%s: ошибка при чтении операций: %w", op, err)
		}

		results = append(results, &detail)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка при итерации строк: %w", op, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("%s: не найдено нарядов для order_num=%s: %w", op, orderNum, sql.ErrNoRows)
	}

	return results, nil
}

func (s *Storage) GetNormOrders(orderNum, orderType string) ([]storage.GetOrderDetails, error) {
	const op = "storage.mysql.GetNormOrders"

	stmt := `SELECT id, order_num, name, count, total_time, created_at, type, part_type, parent_product_id, parent_assembly FROM product_instances 
        	WHERE 1=1 AND (?='' OR order_num LIKE CONCAT('%', ?, '%')) AND (? = '' OR type = ?) AND part_type='main' ORDER BY created_at DESC`

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
			&item.PartType,
			&item.ParentProductID,
			&item.ParentAssembly,
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
