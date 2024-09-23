package infra

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DB_NAME")

	//var dsn string = "test:test@tcp(db)/test?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	dsn := fmt.Sprintf("%s:%s@tcp(db)/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo", user, pass, database)

	fmt.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connceted")
	return db
}
