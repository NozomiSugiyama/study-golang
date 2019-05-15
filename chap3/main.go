package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("db_hc")
	type Result struct{ Value string }
	var out Result
	db.Raw("SELECT 'hello db' as value ").Scan(&out)
	log.Printf("%v", out)
	fmt.Fprintf(w, out.Value)
}

func main() {
	var err error
	db, err = bootDB(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/wep/systems.db.health.check", handler)
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	fcgi.Serve(l, nil)
}

func bootDB(host, user, pass string) (*gorm.DB, error) {
	var err error

	driver := "mysql"
	protocol := "tcp"
	port := 3306
	name := "study_golang"
	args := "?charset=utf8&parseTime=True&loc=Local"

	con, err := gorm.Open(driver,
		fmt.Sprintf("%s:%s@%s([%s]:%d)/%s%s", user, pass, protocol, host, port, name, args),
	)
	if err != nil {
		return nil, err
	}

	con.DB().SetConnMaxLifetime(time.Second * 10)

	err = con.DB().Ping()
	if err != nil {
		return nil, err
	}

	return con, nil
}
