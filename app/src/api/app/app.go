package app

import (
	"api/app/items"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// Needed to sql lite 3
	// _ "github.com/mattn/go-sqlite3"
	"github.com/gchaincl/dotsql"
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
	os.Remove("./foo.db")
	// db, err := sql.Open("mysql", "./foo.db")
	// TODO: Configure this, i'm behind on the tests right now
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_CONTAINER_NAME"), os.Getenv("DB_NAME")))
	if err != nil {
		panic("Could not connect to the db")
	}

	for {
		err := db.Ping()
		if err != nil {
			time.Sleep(1 * time.Second)
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
