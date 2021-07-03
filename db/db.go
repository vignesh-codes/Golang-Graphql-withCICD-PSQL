package db

import (
	"Golang-Graphql-withCICD-PSQL/config"
	"fmt"
	"os"
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/ichtrojan/thoth"
)

type DB *sqlx.DB

var DBConnection DB

func Connect() (DB, error) {
	logger, _ := thoth.Init("log")
	db_user, ok := os.LookupEnv(("DB_USER"))

	if !ok {
		logger.Log(fmt.Errorf("DB_USER not found in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	db_password, ok := os.LookupEnv("DB_PASS")

	if !ok {
		logger.Log(fmt.Errorf("DB_PASS not found in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	db_host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(fmt.Errorf("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	db_port, exist := os.LookupEnv("DB_PORT")

	if !exist {
		logger.Log(fmt.Errorf("DB_PORT not set in .env"))
		log.Fatal("DB_PORT not set in .env")
	}

	db_name, exist := os.LookupEnv("DB_NAME")

	if !exist {
		logger.Log(fmt.Errorf("DB_NAME not set in .env"))
		log.Fatal("DB_NAME not set in .env")
	}
	connect_db := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", db_host, db_port, db_user, db_name, db_password)
	dbConn, err := sqlx.Connect("postgres", connect_db)
	if err != nil {
		fmt.Println(config.DBName)
		return nil, err
	}
	schema1 := `CREATE TABLE IF NOT EXISTS USERS (id SERIAL, first_name varchar(40), last_name varchar(40), email varchar(140) NOT NULL, username varchar(40) PRIMARY KEY, password varchar(60) NOT NULL, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW())`
	_, err = dbConn.Exec(schema1)
	if err != nil {
		return nil, err
	}
	schema2 := `CREATE TABLE IF NOT EXISTS AUDIO (id SERIAL, title varchar(40), description varchar(400), category varchar(40) NOT NULL, creatorname varchar(40), creatoremail varchar(40), createdby varchar(50), created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW())`
	_, err = dbConn.Exec(schema2)
	if err != nil {
		return nil, err
	}
	DBConnection = dbConn
	return dbConn, nil
}

func GetConnection() *sqlx.DB {
	return DBConnection
}
