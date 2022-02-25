package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnection(DBHost, DBPort, DBName, DBUsername, DBPassword string) (*sqlx.DB, error) {
	DBDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	db, err := sqlx.Open("mysql", DBDSN)

	return db, err
}
