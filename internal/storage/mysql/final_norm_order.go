package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
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

// TODO new logic

func (s *Storage) GetPEOProductsByCategory() ([]storage.PEOProduct, []storage.GetWorkers, error) {
	const op = "storage.mysql.GetPEOProductsByCategory"

	// Шаг 1: Получаем всех сотрудников нужной бригады
	//employees, err := s.getEmployeesByTeam()
	//if err != nil {
	//	return nil, nil, fmt.Errorf("%s: %w", op, err)
	//}

	employees, err := s.GetAllWorkers()
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(employees) == 0 {
		// Если нет сотрудников — возвращаем пустые изделия
		return []storage.PEOProduct{}, []storage.GetWorkers{}, nil
	}

	// Шаг 2: Получаем ID сотрудников для фильтрации
	employeeIDs := make([]int64, len(employees))
	for i, emp := range employees {
		employeeIDs[i] = emp.ID
	}

	// Шаг 3: Получаем изделия нужной категории (только assigned, текущий месяц)
	//start, end := getCurrentMonthRange()
	//typePlaceholders := placeholders(len(types))

	//type IN (` + typePlaceholders + `)
	queryProducts := `
		SELECT 
			p.id, p.order_num, p.customer, p.total_time, p.created_at, p.status, p.part_type, p.type, p.parent_product_id, p.parent_assembly, c.short_name_customer,
			p.systema, p.type_izd, p.profile, p.count, p.sqr
		FROM product_instances p 
		LEFT JOIN customer c ON p.customer = c.name
		WHERE status = 'assigned'
		  AND created_at >= ?
		  AND created_at < ?
		ORDER BY order_num, created_at
	`

	start, end := getCurrentMonthRange()

	rowsProducts, err := s.db.Query(queryProducts, start, end)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: ошибка получения изделий: %w", op, err)
	}

	defer rowsProducts.Close()

	// Собираем изделия
	products := make(map[int64]*storage.PEOProduct)
	var productList []storage.PEOProduct

	for rowsProducts.Next() {
		var id int64
		var orderNum, customer, status, partType, Type, parentAssembly, customerType, systema, typeIzd, profile string
		var totalTime, sqr float64
		var createdAt time.Time
		var count int
		var parentProductID sql.NullInt64

		err := rowsProducts.Scan(&id, &orderNum, &customer, &totalTime, &createdAt, &status, &partType, &Type, &parentProductID, &parentAssembly,
			&customerType, &systema, &typeIzd, &profile, &count, &sqr)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: scan product: %w", op, err)
		}

		// Обработка NULL для строк

		if customerType == "" {
			customerType = "не определено" // или ""
		}

		if systema == "" {
			systema = "не определено" // или ""
		}

		if typeIzd == "" {
			typeIzd = "не определено" // или ""
		}

		if profile == "" {
			profile = "не определено" // или ""
		}

		// Преобразуем в *int64
		var parentID *int64 = nil
		if parentProductID.Valid {
			parentID = &parentProductID.Int64
		}

		p := storage.PEOProduct{
			ID:              id,
			OrderNum:        orderNum,
			Customer:        customer,
			TotalTime:       totalTime,
			CreatedAt:       createdAt,
			Status:          status,
			PartType:        partType,
			Type:            Type,
			ParentProductID: parentID,
			ParentAssembly:  parentAssembly,
			CustomerType:    customerType,
			Systema:         systema,
			TypeIzd:         typeIzd,
			Profile:         profile,
			Count:           count,
			Sqr:             sqr,
			EmployeeMinutes: make(map[int64]float64),
		}

		products[p.ID] = &p
		productList = append(productList, p)
	}

	// Шаг 4: Получаем все operation_executors для этих изделий и нужных сотрудников
	if len(productList) == 0 {
		return productList, employees, nil
	}

	productIDs := make([]int64, len(productList))
	for i, p := range productList {
		productIDs[i] = p.ID
	}

	queryExecutors := `
		SELECT product_id, employee_id, actual_minutes
		FROM operation_executors
		WHERE product_id IN (` + placeholders(len(productIDs)) + `)
		  AND employee_id IN (` + placeholders(len(employeeIDs)) + `)
	`

	args := make([]interface{}, 0, len(productIDs)+len(employeeIDs))
	for _, id := range productIDs {
		args = append(args, id)
	}
	for _, id := range employeeIDs {
		args = append(args, id)
	}

	rowsExecutors, err := s.db.Query(queryExecutors, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: ошибка получения исполнителей: %w", op, err)
	}
	defer rowsExecutors.Close()

	// Агрегируем минуты по изделию и сотруднику
	for rowsExecutors.Next() {
		var productID, employeeID int64
		var minutes float64
		err := rowsExecutors.Scan(&productID, &employeeID, &minutes)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: ошибка сканирования исполнителя: %w", op, err)
		}
		//fmt.Println("GGGGGGG", minutes)
		if p, ok := products[productID]; ok {
			p.EmployeeMinutes[employeeID] += minutes
			fmt.Printf("KKKK product=%d, employee=%d, current=%f, adding=%f\n",
				productID, employeeID, p.EmployeeMinutes[employeeID], minutes)
		}
	}

	return productList, employees, nil
}

func (s *Storage) getEmployeesByTeam() ([]storage.PEOEmployee, error) {
	query := `SELECT id, name FROM employees WHERE is_active = TRUE ORDER BY name`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения сотрудников бригады %s: ", err)
	}
	defer rows.Close()

	var emps []storage.PEOEmployee
	for rows.Next() {
		var e storage.PEOEmployee
		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			return nil, fmt.Errorf("ошибка сканирования сотрудника: %w", err)
		}
		emps = append(emps, e)
	}
	return emps, nil
}

func getCurrentMonthRange() (start, end time.Time) {
	now := time.Now()
	start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	end = start.AddDate(0, 1, 0)
	return start, end
}

// placeholders генерирует строку вида "?,?,?"
func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	items := make([]string, n)
	for i := range items {
		items[i] = "?"
	}
	return strings.Join(items, ",")
}
