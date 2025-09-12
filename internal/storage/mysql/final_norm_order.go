package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetFinalNormOrders() ([]storage.ReportFinalOrders, error) {
	const op = "storage.mysql.GetFinalNormOrders"

	query := `
			SELECT DISTINCT
				pi.order_num,
				MIN(pi.created_at) AS first_created,
				COUNT(DISTINCT pi.id) AS product_count
			FROM product_instances pi
			JOIN operation_executors oe ON pi.id = oe.product_id
			GROUP BY pi.order_num
			ORDER BY first_created DESC
		`

	rows, err := s.db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: заказы не найдены: %w", op, err)
		}
		return nil, fmt.Errorf("%s: ошибка выполнения запроса: %w", op, err)
	}
	defer rows.Close()

	var items []storage.ReportFinalOrders
	for rows.Next() {
		var item storage.ReportFinalOrders
		err := rows.Scan(&item.OrderNum, &item.FirstCreated, &item.IzdCount)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования строк для получения всех готовых изделии: %w", op, err)
		}
		items = append(items, item)
	}

	return items, err
}

func (s *Storage) GetSimpleOrderReport(orderNum string) (*storage.OrderFinalReport, error) {
	const op = "storage.mysql.GetSimpleOrderReport"

	query := `
		SELECT
			pi.id,
			pi.order_num,
			pi.name,
			t.name AS template_name,
			ov.operation_name,
			ov.operation_label,
			ov.minutes AS norm_minutes,
			e.name AS employee_name,
			oe.actual_minutes,
			oe.actual_value
		FROM product_instances pi
		JOIN templates t ON pi.template_code = t.code
		JOIN operation_values ov ON pi.id = ov.product_id
		LEFT JOIN operation_executors oe ON ov.product_id = oe.product_id AND ov.operation_name = oe.operation_name
		LEFT JOIN employees e ON oe.employee_id = e.id
		WHERE pi.order_num = ?
		ORDER BY pi.id, ov.operation_name;
	`

	rows, err := s.db.Query(query, orderNum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: заказы не найдены: %w", op, err)
		}
		return nil, fmt.Errorf("%s: ошибка выполнения запроса: %w", op, err)
	}
	defer rows.Close()

	report := &storage.OrderFinalReport{
		OrderNum: orderNum,
		Izdelie:  []storage.IzdelieInfo{},
	}

	// 🔑 Мапа для быстрого доступа к изделию по ID
	productMap := make(map[int64]*storage.IzdelieInfo)

	for rows.Next() {
		var (
			productID      int64
			productName    string
			templateName   string
			operationName  string
			operationLabel string
			normMinutes    float64
			employeeName   sql.NullString
			actualMinutes  sql.NullFloat64
			actualValue    sql.NullFloat64
		)

		// 🔽 Сканируем все поля из строки
		err := rows.Scan(
			&productID,
			&orderNum, // можно не использовать, но нужно прочитать
			&productName,
			&templateName,
			&operationName,
			&operationLabel,
			&normMinutes,
			&employeeName,
			&actualMinutes,
			&actualValue,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования строки: %w", op, err)
		}

		// 🔍 ШАГ 1: Получить или создать изделие
		izd, exists := productMap[productID]
		if !exists {
			izd = &storage.IzdelieInfo{
				ID:           productID,
				Name:         productName,
				TemplateName: templateName,
				Operations:   []storage.OperationsNorm{},
			}
			productMap[productID] = izd
		}

		// 🔍 ШАГ 2: Найти операцию в этом изделии
		var opNorm *storage.OperationsNorm
		for i := range izd.Operations {
			if izd.Operations[i].OperationName == operationName {
				opNorm = &izd.Operations[i]
				break
			}
		}

		// Если операция ещё не добавлена — создаём
		if opNorm == nil {
			opNorm = &storage.OperationsNorm{
				OperationName:  operationName,
				OperationLabel: operationLabel,
				NormMinutes:    normMinutes,
				Executors:      []storage.Workers{},
			}
			izd.Operations = append(izd.Operations, *opNorm)
			// обновляем указатель, потому что слайс мог перераспределиться
			opNorm = &izd.Operations[len(izd.Operations)-1]
		}

		// 🔍 ШАГ 3: Добавить исполнителя, если есть (то есть если e.name NOT NULL)
		if employeeName.Valid {
			worker := storage.Workers{
				WorkerName:    employeeName.String,
				ActualMinutes: actualMinutes.Float64, // будет 0, если NULL
				ActualValue:   actualValue.Float64,   // будет 0, если NULL
			}
			opNorm.Executors = append(opNorm.Executors, worker)
		}
	}

	// 🔁 Проверка на ошибки после цикла
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка при чтении строк: %w", op, err)
	}

	// 📦 Преобразуем мапу в срез
	for _, izd := range productMap {
		report.Izdelie = append(report.Izdelie, *izd)
	}

	return report, nil
}
