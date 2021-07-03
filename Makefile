info:
	echo "A Golang App with GraphQL and Postgres. Has JWT Auth Feature and CI. Visit https://github.com/vignesh-codes/Golang-Graphql-withCICD-PSQL"

build:
	go build -o bin/main main.go

run:
	go run main.go