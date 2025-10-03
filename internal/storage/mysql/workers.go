package mysql

import (
	"fmt"
	"log/slog"
	"vue-golang/internal/storage"
)

func (s *Storage) GetAllWorkers() ([]storage.GetWorkers, error) {
	const op = "storage.mysql.GetWorkers"

	stmt := "SELECT id, name FROM employees WHERE is_active = TRUE"

	var workers []storage.GetWorkers

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка получения операций: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var worker storage.GetWorkers

		err := rows.Scan(&worker.ID, &worker.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row for workers: %w", op, err)
		}

		workers = append(workers, worker)
	}

	return workers, nil
}

// storage/mysql/operations.go

func (s *Storage) SaveOperationWorkers(req storage.SaveWorkers) error {
	const op = "storage.mysql.SaveOperationWorkers"

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	// 🔥 УДАЛЯЕМ ВСЕ НАЗНАЧЕНИЯ ДЛЯ СБОРКИ: корень + прямые дети
	_, err = tx.Exec(`
		DELETE FROM operation_executors
		WHERE product_id = ? 
		   OR product_id IN (
		       SELECT * FROM (
		           SELECT id FROM product_instances WHERE parent_product_id = ?
		       ) AS tmp
		   )
	`, req.RootProductID, req.RootProductID)
	if err != nil {
		return fmt.Errorf("%s: delete old assignments for root=%d: %w", op, req.RootProductID, err)
	}

	stmt, err := tx.Prepare(`
        INSERT INTO operation_executors 
        (product_id, operation_name, employee_id, actual_minutes, notes, actual_value)
        VALUES (?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            actual_minutes = VALUES(actual_minutes),
            actual_value = VALUES(actual_value),
            notes = VALUES(notes),
            updated_at = CURRENT_TIMESTAMP
    `)
	if err != nil {
		return fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	defer stmt.Close()

	for _, a := range req.Assignments {
		_, err := stmt.Exec(
			a.ProductID,
			a.OperationName,
			a.EmployeeID,
			a.ActualMinutes,
			a.Notes,
			a.ActualValue,
		)
		if err != nil {
			return fmt.Errorf("%s: insert assignment for product_id=%d, op=%s: %w", op, a.ProductID, a.OperationName, err)
		}
	}

	fmt.Println("ROOOOOOT", req.RootProductID)
	// 2. Если указано — обновляем статус всей сборки
	if req.UpdateStatus != "" && req.RootProductID != 0 {
		// Обновляем main + все его sub
		result, err := tx.Exec(`
            UPDATE product_instances 
            SET status = ? 
            WHERE id = ? OR parent_product_id = ?
        `, req.UpdateStatus, req.RootProductID, req.RootProductID)

		if err != nil {
			return fmt.Errorf("%s: update status for root_id=%d: %w", op, req.RootProductID, err)
		}

		// Проверим, сколько строк затронуто (для лога)
		count, _ := result.RowsAffected()
		slog.Info("Status updated in transaction",
			slog.String("op", op),
			slog.Int64("root_id", req.RootProductID),
			slog.String("status", req.UpdateStatus),
			slog.Int64("rows_affected", count),
		)

		fmt.Println("SSASAS", req.UpdateStatus)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit transaction: %w", op, err)
	}

	return nil
}
