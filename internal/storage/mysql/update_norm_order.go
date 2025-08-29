package mysql

import (
	"fmt"
	"math/rand"
	"vue-golang/internal/storage"
)

func (s *Storage) UpdateNormOrder(ID int64, update storage.GetOrderDetails) error {
	const op = "storage.mysql.sql.UpdateNormOrder"

	stmt := `UPDATE product_instances SET total_time = ?, type = ? WHERE id = ?`
	stmt1 := `DELETE FROM operation_values WHERE product_id = ?`
	stmt3 := `INSERT INTO operation_values (product_id, operation_name, operation_label, count, value, minutes) VALUES (?, ?, ?, ?, ?, ?)`

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	// 1. Обновляем основное изделие
	_, err = tx.Exec(stmt, update.TotalTime, update.Type, ID)
	if err != nil {
		return fmt.Errorf("%s: update product: %w", op, err)
	}

	// 2. Удаляем старые операции
	_, err = tx.Exec(stmt1, ID)
	if err != nil {
		return fmt.Errorf("%s: delete old operations: %w", op, err)
	}

	// 3. Вставляем новые операции
	stmt4, err := tx.Prepare(stmt3)
	if err != nil {
		return fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	defer stmt4.Close()

	for _, operation := range update.Operations {
		opName := operation.Name
		if opName == "" {
			opName = fmt.Sprintf("extra_%d", rand.Intn(10000))
		}

		_, err := stmt4.Exec(ID, opName, operation.Label, operation.Count, operation.Value, operation.Minutes)
		if err != nil {
			return fmt.Errorf("%s: insert operation %s: %w", op, opName, err)
		}
	}

	// 4. Коммит
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit: %w", op, err)
	}

	return nil
}
