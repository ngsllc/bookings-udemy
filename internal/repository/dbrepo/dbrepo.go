package dbrepo

// now we can swap out our databse if necessary
import (
	"bookings-udemy/internal/config"
	"bookings-udemy/internal/repository"
	"database/sql"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
