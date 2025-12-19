package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetTemplateByCode(ctx context.Context, code string) (*storage.Template, error) {
	const op = "storage.mysql.sql.GetFormByCode"

	query := `
		SELECT id, code, name, category, operations, systema, izd, profile, rules
		FROM dem_templates_al 
		WHERE code = ? AND is_active = TRUE
	`

	template := &storage.Template{}

	// Сканируем JSON как строку
	var operationsJSON string
	var rulesJSON string
	err := s.db.QueryRowContext(ctx, query, code).Scan(
		&template.ID,
		&template.Code,
		&template.Name,
		&template.Category,
		&operationsJSON,
		&template.Systema,
		&template.TypeIzd,
		&template.Profile,
		&rulesJSON,
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

	// парсим json правила
	if err := json.Unmarshal([]byte(rulesJSON), &template.Rules); err != nil {
		return nil, fmt.Errorf("%s: ошибка парсинга JSON правил: %w", op, err)
	}

	return template, nil
}

func (s *Storage) GetAllTemplates(ctx context.Context) ([]*storage.Template, error) {
	const op = "storage.mysql.sql.GetAllForms"

	stmt := "SELECT id, code, name, category, systema, izd, profile FROM dem_templates_al WHERE is_active = TRUE"

	rows, err := s.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var templates []*storage.Template

	for rows.Next() {
		template := &storage.Template{}

		err := rows.Scan(&template.ID, &template.Code, &template.Name, &template.Category, &template.Systema, &template.TypeIzd, &template.Profile)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования строки: %w", op, err)
		}

		templates = append(templates, template)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка при итерации по строкам: %w", op, err)
	}

	return templates, nil
}
