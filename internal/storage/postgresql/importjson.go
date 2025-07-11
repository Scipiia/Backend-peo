package postgresql

//func (s *Storage) ImportJSON(data storage.ImportJSON) error {
//	const op = "storage.postgres.exportjson.sql"
//	stmsInsert1 := "INSERT INTO person (name, age, is_student) VALUES ($1, $2, $3)"
//
//	fmt.Println(stmsInsert1)
//	//var kop = &storage.ImportJSON{}
//
//	//row := s.db.QueryRow(stmsInsert1, &kop.Name, &kop.Age, &kop.IsStudent)
//
//	_, err := s.db.Exec(stmsInsert1, &data.Name, &data.Age, &data.IsStudent)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
