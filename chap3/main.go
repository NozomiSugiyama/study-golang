package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"

	"github.com/NozomiSugiyama/study-golang/chap3/handler"
	"github.com/jinzhu/gorm"

	"github.com/NozomiSugiyama/study-golang/chap3/model"
)

func main() {
	print("Run golang-api")

	db, err := bootDB(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}

	defer db.Close()

	repo := &model.UserRepository{db}
	handler := handler.Handler{UserRepository: repo}

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	http.HandleFunc("/wep/users", handler.UsersHandler)
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
