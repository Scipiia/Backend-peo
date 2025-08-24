package postgresql

//func (s *Storage) GetOrderDetails(id int) (*storage.Order, error) {
//	const op = "storage.order-norm-details.sql"
//
//	stmt := "SELECT id, name, width, height, param, manager FROM orders WHERE id = $1"
//
//	res := &storage.Order{}
//	err := s.db.QueryRow(stmt, id).Scan(&res.ID, &res.NumFer, &res.ClassID, &res.Ordername, &res.EnginerID)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, fmt.Errorf("%s: %w", op, err)
//		} else {
//			return nil, err
//		}
//		//return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	return res, nil
//}
