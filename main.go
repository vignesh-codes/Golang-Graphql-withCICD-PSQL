package main

import (
	"fmt"
	"Golang-Graphql-withCICD-PSQL/db"
	"Golang-Graphql-withCICD-PSQL/schema"
	"log"
	"net/http"
	"os"
	"github.com/neelance/graphql-go/relay"

	"github.com/neelance/graphql-go"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

func main() {
	//initialize logger
	logger, _ := thoth.Init("log")
	if err := godotenv.Load(); err != nil {
		logger.Log(fmt.Errorf("NO .env file Found"))
		log.Fatal("No env file found")
	}
	//get port from .env
	port, ok := os.LookupEnv("PORT")
	
	if !ok {
		logger.Log(fmt.Errorf("PORT not Found in .env"))
		log.Fatal("PORT not set in .env")
	}
	// Connect to DB
	_, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to Database, Exiting...", err)
	}
	//initalize the defined graphql schema
	parsedSchema, err := graphql.ParseSchema(schema.Schema, &schema.RootResolver{})
	if err != nil {
		fmt.Println("Error parsing schema", err)
	}
	//our routes
	http.Handle("/api", &relay.Handler{Schema: parsedSchema})
	http.HandleFunc("/graphiql", schema.GraphiQLHandler)
	
	//connect to port
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Cannot Listen to the port 8000")
	}
}
