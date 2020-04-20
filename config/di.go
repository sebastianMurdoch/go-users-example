package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	di_injector "github.com/sebastianMurdoch/di-injector"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

func NewHerokuContainer() di_injector.DiContainer {
	c := di_injector.NewDiContainer()
	c.AddToDependencies(db(), gin.Default())
	return c
}

func db() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	/* Used for localhost tests */
	if err != nil {
		db = sqlx.MustConnect("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR (255) NOT NULL);`
	db.MustExec(schema)
	return db
}
