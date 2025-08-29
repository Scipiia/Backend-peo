package mysql

import (
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetWorkersNew() ([]storage.GetWorkers, error) {
	const op = "storage.mysql.GetWorkersNew"

	stmt := "SELECT id, name, code FROM employees WHERE is_active = TRUE"

	var workers []storage.GetWorkers

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка получения операций: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var worker storage.GetWorkers

		err := rows.Scan(&worker.ID, &worker.Name, &worker.Code)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row for workers: %w", op, err)
		}

		workers = append(workers, worker)
	}

	return workers, nil
}

func (s *Storage) SaveOperationExecutors(req storage.SaveExecutorsRequest) error {
	const op = "storage.mysql.SaveOperationExecutors"

	// 🔹 Начинаем транзакцию
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin: %w", op, err)
	}
	defer tx.Rollback()

	// 🔹 Подготавливаем запрос
	stmt, err := tx.Prepare(`
		INSERT INTO operation_executors
		(product_id, operation_name, employee_id, actual_minutes, notes, actual_value)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			actual_minutes = VALUES(actual_minutes),
			notes = VALUES(notes),
			actual_value = VALUES(actual_value),
			updated_at = CURRENT_TIMESTAMP
	`)
	if err != nil {
		return fmt.Errorf("%s: prepare: %w", op, err)
	}
	defer stmt.Close()

	for _, e := range req.Executors {
		// Защита: actual_minutes не может быть 0
		//if e.ActualMinutes == 0 {
		//	e.ActualMinutes = 1 // или пропускаем?
		//}

		_, err := stmt.Exec(
			req.ProductID,
			e.OperationName,
			e.EmployeeID,
			e.ActualMinutes,
			e.Notes,
			e.ActualValue,
		)
		if err != nil {
			return fmt.Errorf("%s: insert executor: %w", op, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit: %w", op, err)
	}

	return nil
}
