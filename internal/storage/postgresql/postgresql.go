package postgresql

import (
	_ "github.com/lib/pq"
)

//type Storage struct {
//	db *sql.DB
//}

//func New() (*Storage, error) {
//	const op = "storage.postgresql.New"
//
//	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=mysecretpassword dbname=workers sslmode=disable")
//	//postgresql://postgres:mysecretpassword@localhost:5433/users?sslmode=disable
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	stmt, err := db.Prepare(`
//	CREATE TABLE IF NOT EXISTS users(
//	    id INTEGER PRIMARY KEY,
//	    name TEXT NOT NULL,
//	    worker TEXT,
//	    age INT NOT NULL)`)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	_, err = stmt.Exec()
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	stmt1, err := db.Prepare(`CREATE INDEX IF NOT EXISTS isx_name ON users(name)`)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	_, err = stmt1.Exec()
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	return &Storage{db: db}, nil
//}
//
//func (s *Storage) SaveURL(url, alias string) (int, error) {
//	const op = "storage.postgres.sql"
//
//	stmt := "INSERT INTO url (alias, urls) VALUES ($1, $2) RETURNING id"
//
//	var ids int
//	err := s.db.QueryRow(stmt, url, alias).Scan(&ids)
//	if err != nil {
//		return 0, fmt.Errorf("%s: %w", op, err)
//	}
//
//	return ids, nil
//}
//
//func (s *Storage) GetURL(alias string) (string, error) {
//	const op = "storage.postgres.sql"
//
//	stmt := "SELECT url FROM url WHERE alias = $1"
//
//	var resURL string
//	err := s.db.QueryRow(stmt, alias).Scan(&resURL)
//	if err != nil {
//		return "", fmt.Errorf("%s: %w", op, err)
//	}
//	return resURL, nil
//}
//
//func (s *Storage) SaveMessage(message string) (int, error) {
//	const op = "storage.postgres.sql.saveMessage"
//
//	stmt := "INSERT INTO mess (messages) VALUES ($1) RETURNING id"
//
//	var ids int
//	err := s.db.QueryRow(stmt, message).Scan(&ids)
//	if err != nil {
//		return 0, fmt.Errorf("%s: %w", op, err)
//	}
//
//	return ids, nil
//}
