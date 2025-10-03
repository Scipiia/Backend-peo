package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"vue-golang/internal/storage"
)

func (s *Storage) SaveNormOrder(result storage.OrderNormDetails) (int64, error) {
	const op = "storage.mysql.sql.SaveNormOrder"
	stmt := `INSERT INTO product_instances (order_num, template_code, name, count, total_time, type, part_type, 
            parent_assembly, parent_product_id, customer, position, status, systema, type_izd, profile, sqr) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)`

	exec, err := s.db.Exec(stmt, result.OrderNum, result.TemplateCode, result.Name, result.Count, result.TotalTime,
		result.Type, result.PartType, result.ParentAssembly, result.ParentProductID, result.Customer, result.Position,
		result.Status, result.Systema, result.TypeIzd, result.Profile, result.Sqr)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1452 {
			return 0, fmt.Errorf("%s: Ошибка сохранения нормировки в базу='%s'", op, err)
		}
		return 0, fmt.Errorf("%s: Ошибка сохранения нормировки в базу1='%s'", op, err)
	}

	return exec.LastInsertId()
}

func (s *Storage) SaveNormOperation(OrderID int64, operations []storage.NormOperation) error {
	const op = "storage.mysql.sql.SaveNormOperation"

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin transaction: %w", op, err)
	}

	defer tx.Rollback()

	stmt, _ := tx.Prepare(`
		INSERT INTO operation_values 
			(product_id, operation_name, operation_label, count, value, minutes)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		    operation_name = VALUES(operation_name),
			count = VALUES(count),
			value = VALUES(value)
	`)

	for _, op := range operations {
		_, err := stmt.Exec(OrderID, op.Name, op.Label, op.Count, op.Value, op.Minutes)
		if err != nil {
			return fmt.Errorf("%s: Ошибка сохранения нормировки в базу='%s'", op, err)
		}
	}

	stmt.Close()
	return tx.Commit()
}
