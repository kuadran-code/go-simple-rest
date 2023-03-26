package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlDBStruct struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	database string
}

type SqlDB interface {
	DSN() string
	ORM() *gorm.DB
}

func NewSqlDB(driver string, host string, port string, user string, password string, database string) SqlDB {
	return &sqlDBStruct{
		driver:   driver,
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
}

func (s *sqlDBStruct) DSN() string {
	dsn := ""
	if s.driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.host, s.port, s.user, s.password, s.database)
	} else if s.driver == "mysql" {
		dsn = fmt.Sprintf("mysql://%s:%s@%s:%s/%s", s.user, s.password, s.host, s.port, s.database)
	} else {
		log.Panic("Driver is not supported")
	}

	return dsn
}

func (s *sqlDBStruct) ORM() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // ambang Slow SQL
			LogLevel:                  logger.Silent, // tingkat Log
			IgnoreRecordNotFoundError: true,          // mengabaikan kesalahan ErrRecordNotFound  untuk logger
			Colorful:                  false,         // nonaktifkan warna
		},
	)

	dsnConn := s.DSN()
	var db *gorm.DB
	var err error

	if s.driver == "postgres" {
		db, err = gorm.Open(postgres.Open(dsnConn), &gorm.Config{
			Logger: newLogger,
		})
	} else if s.driver == "mysql" {
		db, err = gorm.Open(mysql.Open(dsnConn), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		log.Fatal("Invalid DSN or driver")
	}

	if err != nil {
		log.Fatal("ORM failed to connect to DB")
	}

	return db
}
