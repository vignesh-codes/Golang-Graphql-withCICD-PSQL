package schema

import (
	"context"
	"Golang-Graphql-withCICD-PSQL/db"
	"Golang-Graphql-withCICD-PSQL/model/Person"
)

type ViewerResolver struct {
	viewer personModel.Person
	
}
type AudioResolver struct {
	aa []personModel.Audio1
	
}
type NewResolver struct {
	audio personModel.Audio
	
}
func (resolver *ViewerResolver) User(ctx context.Context, args struct {
	ID string
}) (personModel.PersonResolver, error) {
	result, err := personModel.GetByUsername(args.ID, db.GetConnection())
	return personModel.PersonResolver{result}, err
}



// func (resolver *ViewerResolver) GetAllAudio(ctx context.Context) (personModel.AudioResolver, error) {
// 	result, err := personModel.GetAll(db.GetConnection())
// 	return personModel.AudioResolver{result}, err
// }