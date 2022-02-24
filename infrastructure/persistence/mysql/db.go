package mysql

import (
	"fmt"

	"github.com/PetengDedet/fortune-post-api/domain/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	MenuRepo repository.MenuRepository
	db       *sqlx.DB
}

func NewRepositories(DBHost, DBPort, DBName, DBUsername, DBPassword string) (*Repositories, error) {
	DBDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	db, err := sqlx.Open("mysql", DBDSN)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		MenuRepo: NewMenuRepository(db),
		db:       db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}
