package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const op = "storage.postgresql.New"

	db, err := sql.Open("mysql", "root:@tcp(mysql-8.0:3306)/test_new_logic?parseTime=true")
	//ubuntu
	//db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test_new_logic?parseTime=true")
	//db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/basa_zapas")
	//postgresql://postgres:mysecretpassword@localhost:5433/users?sslmode=disable
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
