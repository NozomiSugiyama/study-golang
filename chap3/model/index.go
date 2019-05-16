package model

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Session = gorm.DB

var session *gorm.DB
var once sync.Once
var openErr error
var pingErr error

func GetSession(host, user, pass string) (*Session, error) {
	once.Do(func() {
		driver := "mysql"
		protocol := "tcp"
		port := 3306
		name := "study_golang"
		args := "?charset=utf8&parseTime=True&loc=Local"

		con, openErr := gorm.Open(driver,
			fmt.Sprintf("%s:%s@%s([%s]:%d)/%s%s", user, pass, protocol, host, port, name, args),
		)

		if openErr != nil {
			return
		}

		con.DB().SetConnMaxLifetime(time.Second * 10)

		pingErr := con.DB().Ping()

		if pingErr != nil {
			return
		}

		session = con
	})

	if openErr != nil {
		return nil, openErr
	}

	if pingErr != nil {
		return nil, openErr
	}

	return session, nil
}

type Model struct {
	ID        uint      `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
