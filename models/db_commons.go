package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func BorrowDBConn() (*sql.DB, error) {
	db_user := beego.AppConfig.String("mysqluser")
	db_pass := beego.AppConfig.String("mysqlpass")
	db_name := beego.AppConfig.String("mysqldb")
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        db_user, db_pass, db_name)
    return sql.Open("postgres", dbinfo)
}

func ReleaseDBConn(db *sql.DB) {
	db.Close()
}