package schema

import (
	"context"
	"Golang-Graphql-withCICD-PSQL/db"
	"Golang-Graphql-withCICD-PSQL/model/Person"
)

type ViewerResolver struct {
	viewer personModel.Person
	
}

//to map the derived details from validatetoken method to personModel
func (resolver *ViewerResolver) User(ctx context.Context, args struct {
	ID string
}) (personModel.PersonResolver, error) {
	result, err := personModel.GetByUsername(args.ID, db.GetConnection())
	return personModel.PersonResolver{result}, err
}
