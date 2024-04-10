package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB = nil

type DSN struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
	TimeZone string `json:"TimeZone"`
}

func (dsn *DSN) ConnectionString() string {
	str := "host=" + dsn.Host + " user=" + dsn.User + " password=" + dsn.Password + " dbname=" + dsn.DBName + " port=" + dsn.Port + " sslmode=" + dsn.SSLMode + " TimeZone=" + dsn.TimeZone
	return str
}

func Connect() error {

	if Conn != nil {
		return nil
	}

	dsn := DSN{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
	// Abre uma conexão com o banco de dados PostgreSQL
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn.ConnectionString()), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("Conn,", Conn)
	log.Println("Conexão com o PostgreSQL estabelecida com sucesso!")
	return nil
}

func Disconnect() {
	postgres, err := Conn.DB()
	if err != nil {
		return
	}
	postgres.Close()
	Conn = nil
	log.Println("Conexão com o PostgreSQL encerrada com sucesso!")
}
