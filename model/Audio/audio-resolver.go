//nolint:unparam
package audioModel
//Resolvers take care of the response fields for graphql schema



type AudioResolver struct {
	Audio
}

type DeleteResolver struct {
	DeleteHandler
}

type Audio1Resolver struct {
	Audio1
}
//





func (a AudioResolver) Title() string {
	return a.Audio.Title
}

func (a AudioResolver) Description() string {
	return a.Audio.Description
}
func (a AudioResolver) CreatorEmail() string {
	return a.Audio.CreatorEmail
}

func (a AudioResolver) CreatorName() string {
	return a.Audio.CreatorName
}

func (a AudioResolver) Category() string {
	return a.Audio.Category
}

func (a AudioResolver) CreatedBy() string {
	return a.Audio.CreatedBy
}

func (a AudioResolver) Token() string {
	return a.Audio.Token
}

func (a AudioResolver) Message() string {
	return a.Audio.Message
}

func (a AudioResolver) Status() string {
	return a.Audio.Status
}

func (a AudioResolver) Destination() string {
	return a.Audio.Destination
}



func (a1 Audio1Resolver) Title() string {
	return a1.Audio1.Title
}

func (a1 Audio1Resolver) Description() string {
	return a1.Audio1.Description
}
func (a1 Audio1Resolver) CreatorEmail() string {
	return a1.Audio1.CreatorEmail
}

func (a1 Audio1Resolver) CreatorName() string {
	return a1.Audio1.CreatorName
}

func (a1 Audio1Resolver) Category() string {
	return a1.Audio1.Category
}

func (a1 Audio1Resolver) Id() string {
	return a1.Audio1.Id
}
func (a1 Audio1Resolver) Destination() string {
	return a1.Audio1.Destination
}
func (a1 Audio1Resolver) CreatedBy() string {
	return a1.Audio1.CreatedBy
}


func (d DeleteResolver) Message() string {
	return d.DeleteHandler.Message
}

func (d DeleteResolver) Status() string {
	return d.DeleteHandler.Status
}