package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func GetConnection(DBHost, DBPort, DBName, DBUsername, DBPassword string) (*sqlx.DB, error) {
	DBDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUsername, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("mysql", DBDSN)
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	db = sqldblogger.OpenDriver(DBDSN, db.Driver(), loggerAdapter)
	newDb := sqlx.NewDb(db, "mysql")

	return newDb, err
}
