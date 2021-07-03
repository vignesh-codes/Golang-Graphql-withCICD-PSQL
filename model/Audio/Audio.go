package audioModel

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	
	// "go-graphql-start/person/PersonModel"
	// Postgres driver
	_ "github.com/lib/pq"
	
)





func (a Audio) Save(db *sqlx.DB) (Audio, error) {
	tx := db.MustBegin()
	
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	fmt.Println(a.CreatorName, a.CreatorEmail)
	_, err := tx.NamedExec("INSERT INTO audio (title, description, category, creatorname, creatoremail, createdby) VALUES (:title, :description, :category, :creatorname, :creatoremail, :createdby)", a)
	if err != nil {
		return Audio{}, err
	}
	fmt.Println(a.CreatorName, a.CreatorEmail)
	err = tx.Commit()
	return a, err
}

func (a Audio) Update(ID int32, user1 string, db *sqlx.DB) (Audio, error) {
	tx := db.MustBegin()
	
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	fmt.Println(a.CreatorName, a.CreatorEmail)
	_, err := tx.Exec("UPDATE audio SET title=$1, description=$2, category=$3, creatorname=$4, creatoremail=$5 WHERE id=$6 and createdby=$7;", a.Title, a.Description, a.Category, a.CreatorName, a.CreatorEmail, ID, user1)
	if err != nil {
		return Audio{}, err
	}
	fmt.Println(a.CreatorName, a.CreatorEmail)
	err = tx.Commit()
	return a, err
}
// Save the person object to DB


// ComparePassword :Compares person password




func (a Audio) GetByID(ID int32, db *sqlx.DB) (Audio, error) {
	var item Audio
	row := db.QueryRowx("SELECT * FROM audio WHERE id=$1", ID)

	err := row.Scan(&item.Token, &item.Title,  &item.Description, &item.Category, &item.CreatorName, &item.CreatorEmail, &item.CreatedBy )
	if err != nil {
		return Audio{}, err
	}
	return item, nil
	
}


func (a Audio1) DeleteByID(ID int32, username string, db *sqlx.DB) (Audio1, error) {
	var audio Audio1

	row, err := db.Exec("Delete FROM audio WHERE id=$1 and createdby=$2", int(ID), username)
	fmt.Println("CHECK THE STATUS", row)
	if err != nil {
		return Audio1{}, err
	}
	fmt.Println(row)
	return audio, nil
}

func GetAllNew(username string, db *sqlx.DB) ([]Audio1, error) {
	var audio []Audio1

	row := db.QueryRowx("SELECT * FROM audio")

	err := row.StructScan(&audio)
	if err != nil {
		fmt.Println(err)
	}
	return audio, nil

}

