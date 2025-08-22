package mysql

import (
	"context"
	"fmt"
	"strings"
	"time"
	updatenormorder "vue-golang/http-server/order-details/update-norm-order"
	"vue-golang/internal/storage"
)

func (s *Storage) GetAllProducts(from, to, orderNum, itemType, profil, name string) ([]storage.ProductItem, error) {
	const op = "storage.mysql.sql.GetAllProducts"

	query := `
        SELECT type, order_num, name, count, profil, total_time, order_id, result_id, created_at, updated_at
        FROM (
            -- 1. Лоджии
            SELECT 
                'loggia' AS type,
                l.order_num,
                l.name,
                l.count,
                l.profil,
                l.total_time,
                l.order_id,
                l.id AS result_id,
                ao.created_at,
                ao.updated_at
            FROM dem_test_golang_result_loggia l
            JOIN dem_test_golang_all_orders ao ON l.order_id = ao.id
            WHERE 1=1
                AND (? = '' OR ao.created_at >= ?)
                AND (? = '' OR ao.created_at < DATE_ADD(?, INTERVAL 1 DAY))
                AND (? = '' OR l.order_num LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR ? = 'loggia')
                AND (? = '' OR l.profil LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR l.name LIKE CONCAT('%', ?, '%'))

            UNION ALL
             -- 2. Витражи
            SELECT 
                'vitraj' AS type,
                v.order_num,
                v.name,
                v.count,
                v.profil,
                v.total_time,
                v.order_id,
                v.id AS result_id,
                ao.created_at,
                ao.updated_at
            FROM dem_test_golang_result_vitraj v
            JOIN dem_test_golang_all_orders ao ON v.order_id = ao.id
            WHERE 1=1
                AND (? = '' OR ao.created_at >= ?)
                AND (? = '' OR ao.created_at < DATE_ADD(?, INTERVAL 1 DAY))
                AND (? = '' OR v.order_num LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR ? = 'vitraj')
                AND (? = '' OR v.profil LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR v.name LIKE CONCAT('%', ?, '%'))

            UNION ALL
             -- 3. Двери
            SELECT 
                'door' AS type,
                d.order_num,
                d.name,
                d.count,
                d.profil,
                d.total_time,
                d.order_id,
                d.id AS result_id,
                ao.created_at,
                ao.updated_at
            FROM dem_test_golang_result_door d
            JOIN dem_test_golang_all_orders ao ON d.order_id = ao.id
            WHERE 1=1
                AND (? = '' OR ao.created_at >= ?)
                AND (? = '' OR ao.created_at < DATE_ADD(?, INTERVAL 1 DAY))
                AND (? = '' OR d.order_num LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR ? = 'door')
                AND (? = '' OR d.profil LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR d.name LIKE CONCAT('%', ?, '%'))

            UNION ALL
            -- 4. Окна
            SELECT 
                'window' AS type,
                w.order_num,
                w.name,
                w.count,
                w.profil,
                w.total_time,
                w.order_id,
                w.id AS result_id,
                ao.created_at,
                ao.updated_at
            FROM dem_test_golang_result_window w
            JOIN dem_test_golang_all_orders ao ON w.order_id = ao.id
            WHERE 1=1
                AND (? = '' OR ao.created_at >= ?)
                AND (? = '' OR ao.created_at < DATE_ADD(?, INTERVAL 1 DAY))
                AND (? = '' OR w.order_num LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR ? = 'window')
                AND (? = '' OR w.profil LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR w.name LIKE CONCAT('%', ?, '%'))

            UNION ALL
             -- 5. Глухари
            SELECT 
                'glyhar' AS type,
                g.order_num,
                g.name,
                g.count,
                g.profil,
                g.total_time,
                g.order_id,
                g.id AS result_id,
                ao.created_at,
                ao.updated_at
            FROM dem_test_golang_result_glyhar g
            JOIN dem_test_golang_all_orders ao ON g.order_id = ao.id
            WHERE 1=1
                AND (? = '' OR ao.created_at >= ?)
                AND (? = '' OR ao.created_at < DATE_ADD(?, INTERVAL 1 DAY))
                AND (? = '' OR g.order_num LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR ? = 'glyhar')
                AND (? = '' OR g.profil LIKE CONCAT('%', ?, '%'))
                AND (? = '' OR g.name LIKE CONCAT('%', ?, '%'))
        ) AS combined
        ORDER BY order_num, type
    `

	//Подготавливаем 5 наборов по 8 параметров = 40 параметров
	args := make([]interface{}, 0, 40)

	// Каждый блок требует одни и те же 8 параметров
	for i := 0; i < 5; i++ {
		args = append(args,
			from, from,
			to, to,
			orderNum, orderNum,
			itemType, itemType,
			profil, profil,
			name, name,
		)
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to query: %w", op, err)
	}
	defer rows.Close()

	var products []storage.ProductItem

	for rows.Next() {
		var p storage.ProductItem
		err := rows.Scan(&p.Type, &p.OrderNum, &p.Name, &p.Count, &p.Profil, &p.TotalTime, &p.OrderID, &p.ResultID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: row iteration error: %w", op, err)
	}

	return products, nil
}

// storage/psql.go
func (s *Storage) UpdateGlyhari(ID int, orderNum string, operations map[string]float64, additionalOps []updatenormorder.AddOpRequest) error {
	const op = "storage.mysql.sql.UpdateGlyhari"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}
	defer tx.Rollback()

	// --- Шаг 1: Формируем SET часть для операций ---
	var setParts []string
	var args []interface{}

	for field, value := range operations {
		//if s.isValidGlyharField(field) {
		setParts = append(setParts, fmt.Sprintf("`%s` = ?", field))
		args = append(args, value)
		//}
	}

	// --- Шаг 2: Пересчитываем total_time ---
	var totalTime float64
	for _, value := range operations {
		totalTime += value
	}
	// Добавляем total_time в SET и args
	setParts = append(setParts, "`total_time` = ?")
	args = append(args, totalTime)

	// --- Шаг 3: Добавляем order_id в WHERE ---
	args = append(args, ID)

	query := fmt.Sprintf(
		`UPDATE dem_test_golang_result_glyhar SET %s WHERE id = ?`,
		strings.Join(setParts, ", "),
	)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: failed to update glyhar: %w", op, err)
	}

	// --- 2. Добавляем/обновляем доп. операции ---
	for _, addOp := range additionalOps {
		// Проверим, нет ли уже такой операции (по имени и order_id)?
		// Или просто добавим новую?
		// Пока — просто добавим (можно потом улучшить)
		_, err := tx.ExecContext(ctx,
			`INSERT INTO dem_test_golang_additional_operations (order_id, operation_name, duration, comment) VALUES (?, ?, ?, ?)`,
			ID, addOp.Operation, addOp.Duration, addOp.Comment,
		)
		if err != nil {
			return fmt.Errorf("%s: failed to insert additional operation: %w", op, err)
		}
	}

	// --- Шаг 4: Обновляем updated_at в основной таблице ---
	_, err = tx.ExecContext(ctx,
		`UPDATE dem_test_golang_all_orders SET updated_at = NOW() WHERE order_num LIKE ?`,
		orderNum,
	)
	if err != nil {
		return fmt.Errorf("%s: failed to update timestamp: %w", op, err)
	}

	// --- Шаг 5: Коммитим ---
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return nil
}

//func (s *Storage) isValidGlyharField(field string) bool {
//	validFields := map[string]bool{
//		"nast_napil": true, "napil": true, "napil_krishek": true, "napil_impost": true,
//		"soedinitel": true, "promej_sborka": true, "impost_sverlo": true, "impost_frezer": true,
//		"impost_sborka": true, "opres_nastr": true, "opresovka": true, "ystan_yplotn": true,
//		"zashivka": true, "napil_stoiki_do3m": true, "napil_stoiki_bol3m": true,
//		"napil_rigel_do1m": true, "napil_rigel_bol1m": true, "sverl_rigel_zamok": true,
//		"ystan_zamkov": true, "napil_shtapik": true, "ypakovka": true, "frezer_rigel": true,
//		"obrabot_ram": true, "hands_sborka": true, "frezer_nastr": true, "shtiftovka": true,
//		"ystanovka_zapoln": true, "napil_donnik": true, "adapter_napil": true, "adapter_ystan": true,
//		"ystan_yplotn_falc": true, "total_time": true,
//	}
//	return validFields[field]
//}
