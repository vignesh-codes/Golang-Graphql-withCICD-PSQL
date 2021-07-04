package audioModel
//Handling all AUDIO related controls here -> interacting with DB

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	
	// "go-graphql-start/person/PersonModel"
	// Postgres driver
	_ "github.com/lib/pq"
	
)




//save the audio to db
func (a Audio) Save(db *sqlx.DB) (Audio, error) {
	tx := db.MustBegin()
	
	
	_, err := tx.NamedExec("INSERT INTO audio (title, description, category, creatorname, creatoremail, createdby, destination) VALUES (:title, :description, :category, :creatorname, :creatoremail, :createdby, :destination)", a)
	if err != nil {
		return Audio{}, err
	}
	
	err = tx.Commit()
	return a, err
}

//Update the audio to DB only if the given creator username matches
func (a Audio) Update(ID int32, user1 string, db *sqlx.DB) (Audio, error) {
	tx := db.MustBegin()
	
	
	_, err := tx.Exec("UPDATE audio SET title=$1, description=$2, category=$3, creatorname=$4, creatoremail=$5, destination=$6 WHERE id=$7 and createdby=$8;", a.Title, a.Description, a.Category, a.CreatorName, a.CreatorEmail, a.Destination, ID, user1)
	if err != nil {
		return Audio{}, err
	}
	
	err = tx.Commit()
	return a, err
}




//get audio by id
func (a Audio) GetByID(ID int32, db *sqlx.DB) (Audio, error) {
	var item Audio
	row := db.QueryRowx("SELECT * FROM audio WHERE id=$1", ID)

	err := row.Scan(&item.Token, &item.Title,  &item.Description, &item.Category, &item.CreatorName, &item.CreatorEmail, &item.CreatedBy, &item.Destination )
	if err != nil {
		return Audio{}, err
	}
	return item, nil
	
}

//Delete audio by id happens only if the username matches the given id. Otherwise it means not the owner triyng to delete
func (a Audio1) DeleteByID(ID int32, username string, db *sqlx.DB) (Audio1, error) {
	var audio Audio1

	row, err := db.Exec("Delete FROM audio WHERE id=$1 and createdby=$2", int(ID), username)
	fmt.Println("CHECK THE STATUS", row)
	if err != nil {
		return Audio1{}, err
	}
	
	return audio, nil
}

//Get all audio
func GetAllNew(username string, db *sqlx.DB) ([]Audio1, error) {
	var audio []Audio1

	row := db.QueryRowx("SELECT * FROM audio")

	err := row.StructScan(&audio)
	if err != nil {
		fmt.Println(err)
	}
	return audio, nil

}

