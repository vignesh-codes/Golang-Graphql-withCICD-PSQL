package db

import (
	"fmt"
	"Golang-Graphql-withCICD-PSQL/config"

	"github.com/jmoiron/sqlx"
)

type DB *sqlx.DB

var DBConnection DB

func Connect() (DB, error) {
	// connString := fmt.Sprintf("host=localhost port=5432 user=postgres password=root dbname=test sslmode=disable")
	dbConn, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=test password=root sslmode=disable")
	if err != nil {
		fmt.Println(config.DBName)
		return nil, err
	}
	schema1 := `CREATE TABLE IF NOT EXISTS USERS (id SERIAL, first_name varchar(40), last_name varchar(40), email varchar(140) NOT NULL, username varchar(40) PRIMARY KEY, password varchar(60) NOT NULL)`
	_, err = dbConn.Exec(schema1)
	if err != nil {
		return nil, err
	}
	schema2 := `CREATE TABLE IF NOT EXISTS AUDIO (id SERIAL, title varchar(40), descriptionription varchar(400), category varchar(40) NOT NULL, creatorname varchar(40), creatoremail varchar(40) )`
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
