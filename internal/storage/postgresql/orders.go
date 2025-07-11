package postgresql

//func (s *Storage) GetOrders() ([]*storage.Order, error) {
//	const op = "storage.orders.sql"
//
//	stmt := "SELECT id, name, width, height, param, manager FROM orders LIMIT 10"
//
//	rows, err := s.db.Query(stmt)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	var orders []*storage.Order
//
//	for rows.Next() {
//		s := &storage.Order{}
//
//		err = rows.Scan(&s.ID, &s.Name, &s.Width, &s.Height, &s.Param, &s.Manager)
//		if err != nil {
//			return nil, err
//		}
//
//		orders = append(orders, s)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return orders, err
//}
