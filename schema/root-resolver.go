package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"Golang-Graphql-withCICD-PSQL/auth"
	"Golang-Graphql-withCICD-PSQL/config"
	db "Golang-Graphql-withCICD-PSQL/db"
	"Golang-Graphql-withCICD-PSQL/model/Audio"
	"Golang-Graphql-withCICD-PSQL/model/Person"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//RootResolver -> Has the methods to handle the operations defined in schema.grapql
type RootResolver struct{}

//To view the users info based ont he given token after login
func (resolver *RootResolver) Viewer(ctx context.Context, args struct{ Token string }) (*ViewerResolver, error) {
	user, err := auth.ValidateToken(args.Token)
	if err != nil {
		return nil, err
	}
	user.Message = "SUCCESS"
	user.Status = "200"
	return &ViewerResolver{user}, nil
}


//To get the list of audios from the db with pagination
func (resolver *RootResolver) Getall(ctx context.Context, args struct{
	Limit string
	Offset string}) (string, error){
	db := db.GetConnection()
	
	// var item1 interface{}
	var audio []audioModel.Audio1
	limit := args.Limit
	offset := args.Offset
	fmt.Println(limit, offset)
	row, err := db.Query("SELECT * FROM audio LIMIT $1 OFFSET $2;", limit, offset)
		if err != nil {
			fmt.Println(err)
		}
	for row.Next(){
		var a audioModel.Audio1
		err := row.Scan(&a.Id, &a.Title, &a.Description, &a.Category, &a.CreatorName, &a.CreatorEmail, &a.CreatedBy, &a.Destination)
		if err != nil {
			fmt.Println(err)
		}
		audio = append(audio, a)
	}
	out, err := json.Marshal(audio)

	fmt.Println(audio, err)
	return string(out), nil
	
}



// func (resolver *RootResolver) Newgetall(ctx context.Context, args struct{
// 	Limit string
// 	Offset string}) ([]personModel.Audio1Resolver){
// 	db := db.GetConnection()
	
// 	// var item1 interface{}
// 	var audio []interface{}
// 	// outy := make([]personModel.Audio1,0,50)
// 	limit := args.Limit
// 	offset := args.Offset
// 	fmt.Println(limit, offset)
// 	row, err := db.Query("SELECT * FROM audio LIMIT $1 OFFSET $2;", limit, offset)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	for row.Next(){
// 		var p personModel.Audio1
// 		err := row.Scan(&p.Id, &p.Title, &p.Description, &p.Category, &p.CreatorName, &p.CreatorEmail)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(p)
// 		audio = append(audio, p)
// 	}
// 	// out, err := json.Marshal(audio)
// 	// if err != nil{
// 	// 	fmt.Println("ERROR")
// 	// }
// 	fmt.Println("1 Audio", audio)
// 	newout, err := json.Marshal(audio)
// 	if err != nil{
// 		fmt.Println("ERROR", newout)
// 	}
// 	tob, err := json.Marshal(audio)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var mapping []personModel.Audio1
// 	if err := json.Unmarshal(tob, &mapping);err != nil {
// 			fmt.Println("ERROR")
// 		}
// 	fmt.Println("mapping", mapping)
	
// 	return []personModel.Audio1Resolver{}
	
// }


// To login and get the access token
func (resolver *RootResolver) Login(ctx context.Context, args struct {
	Username string
	Password string
}) (personModel.LoginResolver) {
	db := db.GetConnection()

	var user personModel.Person
	var login personModel.LoginHandler
	user = personModel.Person{Username: args.Username, Password: args.Password}
	personData, err := user.ComparePassword(db)
	fmt.Println(personData)
	if err != nil {
		login.Message = "Incorrect Passwrod"
		login.Status = "400"
		return personModel.LoginResolver{login}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "Golang-Graphql-withCICD-PSQL",
		"iot":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 48).Unix(),
		"data": personData,
	})

	tokenString, err := token.SignedString(config.GetJWTSecret())
	if err != nil {
		login.Message = "Token Expired - Incorrect Token"
		login.Status = "400"
		return personModel.LoginResolver{login}
	}
	login.Token = tokenString
	login.Message = "Login Success"
	login.Status = "200"
	return personModel.LoginResolver{login}

}


//To signup -> create a new person
func (resolver *RootResolver) Signup(ctx context.Context, args struct {
	Person personModel.PersonInput
}) (personModel.PersonResolver) {
	db := db.GetConnection()

	var user personModel.Person
	var responser personModel.Response
	user = personModel.Person{
		Username:  args.Person.Username,
		Password:  args.Person.Password,
		EmailID:   args.Person.EmailID,
		FirstName: args.Person.FirstName,
		LastName:  args.Person.LastName,
	}

	err := user.Validate()
	if err != nil {
		fmt.Println(err)
		user.Message = "Validation Error - Make sure Username, Password, FirstName is >2 and <50 char"
		user.Status = "422"
		return personModel.PersonResolver{user}
	}
	savedPerson, err := user.Save(db)
	if err != nil {
		fmt.Println(err)
		user.Message = "Username Already Exist. Try Other"
		user.Status = "400"
		return personModel.PersonResolver{user}
	}
	fmt.Println(savedPerson, responser)
	user.Message = "SUCCESS"
	user.Status = "201"
	return personModel.PersonResolver{Person: user}

}


// To create a new audio row in DB -> can create only after loogin in (with access token for now)
func (resolver *RootResolver) Create(ctx context.Context, args struct {
	Audio audioModel.AudioInput
	// Person audioModel.PersonInput
}) (audioModel.AudioResolver, error) {
	var audio audioModel.Audio
	db := db.GetConnection()
	//validate the token and get info about user to put in audio row
	user1, err := auth.ValidateToken1(args.Audio.Token)
	if err != nil {
		audio.Message = "Unathorized - Access Token Not Valid"
		audio.Status = "401"
		return audioModel.AudioResolver{Audio: audio}, nil
	}
	fmt.Println("u1:", user1)
	// handle storing in S3 of the uploaded audio file
	
	
	fmt.Println("user is: ",user1)
	audio = audioModel.Audio{
		Title:  args.Audio.Title,
		Description:  args.Audio.Description,
		Category:   args.Audio.Category,
		CreatorName: user1.FirstName,
		CreatorEmail:  user1.EmailID,
		CreatedBy: user1.Username,
		Destination: "s3SampleURL",
	}
	fmt.Println(audio.Title, audio.Description)
	
	fmt.Println(user1)
	
	savedAudio, err := audio.Save(db)
	savedAudio.Message = "Success"
	savedAudio.Status = "201"
 	if err != nil {
		 fmt.Println(err)
		savedAudio.Message = "Error"
		savedAudio.Status = "400"
		return audioModel.AudioResolver{Audio: savedAudio}, nil
	}
	return audioModel.AudioResolver{Audio: savedAudio}, nil
	
}


// To update the audio info in DB (with access token only)
func (resolver *RootResolver) Update(ctx context.Context, args struct {
	ID int32
	Audio audioModel.AudioInput
	// Person audioModel.PersonInput
}) (audioModel.AudioResolver, error) {
	var audio audioModel.Audio
	db := db.GetConnection()
	user1, err := auth.ValidateToken1(args.Audio.Token)
	
	if err != nil {
		audio.Message = "Unathorized - Access Token Not Valid"
		audio.Status = "401"
		return audioModel.AudioResolver{Audio: audio}, nil
	}
	fmt.Println(user1)
	// When the uploaded audio file has to be replaced, the following logic will be performed
	// Delete the file stored in s3 using the url derived from the input
	// Upload the new file and store the url of the new file to the updated schema


	
	audio = audioModel.Audio{
		Title:  args.Audio.Title,
		Description:  args.Audio.Description,
		Category:   args.Audio.Category,
		CreatorName: args.Audio.CreatorName,
		CreatorEmail:  user1.EmailID,
		Destination: "sampleUpdatedS3URL",
	}
	
	//save the updated audio only if the username is the owener of the audio content
	savedAudio, err := audio.Update(args.ID, user1.Username, db)
	savedAudio.Message = "Success"
	savedAudio.Status = "201"
 	if err != nil {
		 fmt.Println(err)
		savedAudio.Message = "Error Updating the Audio"
		savedAudio.Status = "400"
		return audioModel.AudioResolver{Audio: savedAudio}, nil
	}
	return audioModel.AudioResolver{Audio: savedAudio}, nil
	
}

//To get info of a specific audio based on the Audio ID given 
func (resolver *RootResolver) Getbyid(ctx context.Context, args struct{ ID int32}) (audioModel.AudioResolver, error){
	db := db.GetConnection()
	var audio audioModel.Audio
	item := audio
	fmt.Println(args.ID, db)
	
	derivedAudio, err := audio.GetByID(args.ID, db)
	fmt.Println(derivedAudio.Title)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Id is: ", item.Token)
	//derivedAudio.Token = Id of the audio-> if empty means no row with this id is found. using .token just to match the schema without creating new one
	if (derivedAudio.Token == ""){
		
		derivedAudio.Message = "NOT FOUND"
		derivedAudio.Status = "400"
		return audioModel.AudioResolver{Audio: derivedAudio}, nil
	}
	
	derivedAudio.Message = "Success"
	derivedAudio.Status = "200"
	
	return audioModel.AudioResolver{Audio: derivedAudio}, nil
	
}

// To delete a audio (happens only for signed in users (with access token))
func (resolver *RootResolver) Deletebyid(ctx context.Context, args struct{ 
	ID int32
	Token string
	}) (audioModel.DeleteResolver){
	db := db.GetConnection()
	var audio audioModel.Audio1
	var info audioModel.DeleteHandler
	
	user1, err := auth.ValidateToken1(args.Token)
	
	if err != nil {
		info.Message = "Unathorized - Access Token Not Valid"
		info.Status = "401"
		return audioModel.DeleteResolver{DeleteHandler: info}
	}
	fmt.Println(args.ID, db)
	//Delete the file in the s3 storage after deleting the row in the DB
	// This happens only after certain conditions map in the DB
	deleteRow, err := audio.DeleteByID(args.ID, user1.Username, db)
	if err != nil {
		fmt.Println(err)
		info.Message = "Error Deleting the Audio"
		info.Status = "400"
		return audioModel.DeleteResolver{DeleteHandler: info}
	}
	fmt.Println(deleteRow)
	info.Message = "Success"
	info.Status = "201"
	return audioModel.DeleteResolver{DeleteHandler: info}
}