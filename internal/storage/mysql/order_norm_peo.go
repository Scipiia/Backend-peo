package mysql

import (
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetAllProducts() ([]storage.ProductItem, error) {
	const op = "storage.mysql.sql.GetAllProducts"

	query := `
        SELECT 'loggia' AS type, l.order_num, name, count, profil, total_time, order_id, l.id, ao.created_at, ao.updated_at FROM dem_test_golang_result_loggia l
        JOIN dem_test_golang_all_orders ao ON l.order_id = ao.id
        UNION ALL
        SELECT 'vitraj' AS type, v.order_num, name, count, profil, total_time, order_id, v.id, ao.created_at, ao.updated_at FROM dem_test_golang_result_vitraj v
        JOIN dem_test_golang_all_orders ao ON v.order_id = ao.id
        UNION ALL
        SELECT 'door'  AS type, d.order_num, name, count, profil, total_time, order_id, d.id, ao.created_at, ao.updated_at FROM dem_test_golang_result_door d
        JOIN dem_test_golang_all_orders ao ON d.order_id = ao.id
        UNION ALL
        SELECT 'window'  AS type, w.order_num, name, count, profil, total_time, order_id, w.id, ao.created_at, ao.updated_at FROM dem_test_golang_result_window w
        JOIN dem_test_golang_all_orders ao ON w.order_id = ao.id
        UNION ALL
        SELECT 'glyhar'  AS type, g.order_num, name, count, profil, total_time, order_id, g.id, ao.created_at, ao.updated_at FROM dem_test_golang_result_glyhar g
        JOIN dem_test_golang_all_orders ao ON g.order_id = ao.id
        ORDER BY order_num, type
    `

	rows, err := s.db.Query(query)
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
