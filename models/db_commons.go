package modules

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "quanzhang"
    DB_NAME     = "godb"
)

func BorrowDBConn() (*DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    return sql.Open("postgres", dbinfo)
}

func ReleaseDBConn(db *DB) {
	db.Close()
}