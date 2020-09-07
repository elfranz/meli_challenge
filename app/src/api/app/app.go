package app

import (
	"api/app/items"
	"database/sql"
	"fmt"
	"os"
	"time"

	// handles sql file for schema file
	"github.com/gchaincl/dotsql"
	// web framework
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	r *gin.Engine
)

const (
	port string = ":8080"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	items.Configure(r, db)
	r.Run(port)
}

func configDataBase() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_CONTAINER_NAME"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Could not connect to the db")
	}

	for {
		err := db.Ping()
		if err != nil {
			// added this print so we see if the db is stuck
			fmt.Println("Waiting for db...")
			time.Sleep(1 * time.Second)
			// skips iteration if database does not respond
			continue
		}
		loadSchema(db)
		return db
	}

}

func loadSchema(db *sql.DB) {
	dot, err := dotsql.LoadFromFile("src/api/app/db/schema.sql")
	if err != nil {
		panic("Could not load database schema file")
	}
	res, loaderr := dot.Exec(db, "load-schema")
	_ = res
	if loaderr != nil {
		panic("Could not load schema")
	}
}
