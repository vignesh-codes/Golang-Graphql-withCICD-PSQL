# GraphQL Based API Using Golang and PostgreSQL
 The Project involves backend for storing audio file information.

## Features
- Signup Sigin Feature using JWT Auth
- Typical CRUD Operations
- Simple Role Based Access to create, update and delete operations (validating JWT token for matching user specific rows)
- CI - Github action workflows for lint tests and multiple platform compatability tests
- Jobs for automated build and publish docker images

## Tech
- [Golang] 
- [PostgreSQL]
- [Graphql-go] 
- [jwt-go] - to provide JWT Auth feature
- [govalidator] - To validate the users input
- [godotenv] - for reading .env configs
- [thoth] - Error logging for go
- [sqlx] - a library which provides a set of extensions on go's standard database/sql library


## Installation
The project requires postgresql to be installed and running in the machine
```sh
Git clone https://github.com/vignesh-codes/Golang-Graphql-withCICD-PSQL
cd Golang-Graphql-withCICD-PSQL
```

Create a .env file and add the following
```
PORT=8000
DB_HOST=localhost
DB_PORT=5432
DB_NAME=test
DB_USER=postgres
DB_PASS=root
```
And then to the run app
```sh
go clean
go build -o main .
go run main.go
```
Check http://localhost:8000/graphiql for graphiql playground

## Github Actions - CI
The .github/workflows contains scripts for CI pipeline in github actions. Make sure you add the necessary secrets in your github repo settings.

The Build and publish job will pass only if the golangci-lint test gets passed

## MOCK TESTS
Go to  http://localhost:8000/graphiql for graphiql playground

To Create a New User/Person: 
```
mutation {
	signup(person: {Username:  "AUser1",
		Password:  "abcd123",
		EmailID:   "a@a.com",
		FirstName: "NewMan",
		LastName:  "Nameless"})
  {
    Message
    Status
    FirstName
    LastName
    EmailID
    Username
  }
}
```
To Login:
```
query {
	login(username:  "AUser1",
		password:  "abcd123",
		)
		{
		    Token
		    Message
		    Status
		}
}
```
To Create a new Audio Information:
```
mutation {
	create(audio: {Token:  "jwttoken",
		Title:  "abcd123",
		Description:   "A Cool Podcast",
		Category: "ASD",
		})
  {
    Title
    Description
    Status
    Message
    CreatorName
	  CreatorEmail
    Destination
    CreatedBy
  }
}
```
To update the audio info:
```
mutation {
  
	update(id: 46, audio: {Token:  "jwttoken",
		Title:  "abcd1",
		Description:   "A cooler podcost v2.0",
		Category: "changed",
    
		})
  {
    Status
    Message
    Title
    Description
    CreatorName
	  CreatorEmail
    Destination
    CreatedBy
  }
}
```
To get Audio info by ID:

```
query {
  getbyid(id: enterurIDinInt)
  {
    Status
    Message
    Description
    Destination
    CreatedBy
    CreatorName
    CreatorEmail
  }
}
```
Get All Audio list:
```
{
  getall(limit:"youlimit",
		offset: "youroffset")
}
```
To get user info
```
query {
	viewer(token:"yourtoken",		
  ){
    user(id:"yourusername") {
      FirstName,
      LastName,
      Username
    }
  }
}
```
To Delete an Audio:
```
mutation {
 deletebyid(Token:"yourtoken",
  id:audioId_INT)
  {
    Message
    Status
  }
}
```
