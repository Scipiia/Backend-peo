package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
)

// storage/mysql/sql/templates.go

func (s *Storage) GetFormByCode(code string) (*storage.Form, error) {
	const op = "storage.mysql.sql.GetFormByCode"

	query := `
		SELECT id, code, name, category, operations 
		FROM templates 
		WHERE code = ? AND is_active = TRUE
	`

	template := &storage.Form{}

	// Сканируем JSON как строку
	var operationsJSON string
	err := s.db.QueryRow(query, code).Scan(
		&template.ID,
		&template.Code,
		&template.Name,
		&template.Category,
		&operationsJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: шаблон с code='%s' не найден: %w", op, code, err)
		}
		return nil, fmt.Errorf("%s: выполнение запроса завершилось ошибкой: %w", op, err)
	}

	// Парсим JSON операций
	if err := json.Unmarshal([]byte(operationsJSON), &template.Operations); err != nil {
		return nil, fmt.Errorf("%s: ошибка парсинга JSON операций: %w", op, err)
	}

	return template, nil
}

func (s *Storage) GetAllForms() ([]*storage.Form, error) {
	const op = "storage.mysql.sql.GetAllForms"

	stmt := "SELECT id, code, name, category FROM templates WHERE is_active = TRUE"

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var templates []*storage.Form
	//var operationsJSON string

	for rows.Next() {
		template := &storage.Form{}

		err := rows.Scan(&template.ID, &template.Code, &template.Name, &template.Category)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования строки: %w", op, err)
		}

		//if err := json.Unmarshal([]byte(operationsJSON), &template.Operations); err != nil {
		//	return nil, fmt.Errorf("%s: ошибка парсинга JSON операций: %w", op, err)
		//}

		templates = append(templates, template)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка при итерации по строкам: %w", op, err)
	}

	return templates, nil
}
