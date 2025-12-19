package mysql

import (
	"context"
	"fmt"
	"vue-golang/internal/storage"
)

func (s *Storage) GetOrderMaterials(ctx context.Context, id, pos int) ([]*storage.KlaesMaterials, error) {
	const op = "storage.order-dem-details.GetOrderMaterials.sql"

	stmt := `SELECT idorders, articul_mat, name_mat, width, height FROM dem_klaes_materials 
            	WHERE idorders= ? AND position=? AND LOWER(TRIM(name_mat)) IN ('импост', 'доп. импост')`

	rows, err := s.db.QueryContext(ctx, stmt, id, pos)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка выполнения запроса для получения материалов %w", op, err)
	}

	defer rows.Close()

	var materials []*storage.KlaesMaterials

	for rows.Next() {
		var material storage.KlaesMaterials

		err := rows.Scan(&material.OrderID, &material.ArticulMat, &material.NameMat, &material.Width, &material.Height)
		if err != nil {
			return nil, fmt.Errorf("%s: ошибка сканирования строк материалов %w", op, err)
		}

		materials = append(materials, &material)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: ошибка сканирования строк для получения материалов %w", op, err)
	}

	return materials, nil
}
