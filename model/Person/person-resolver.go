package personModel


type PersonResolver struct {
	Person
}

type AudioResolver struct {
	Audio
}
type ResponseResolver struct {
	Response
}

type DeleteResolver struct {
	DeleteHandler
}

type LoginResolver struct {
	LoginHandler
}
type Audio1Resolver struct {
	Audio1
}
//
func (d DeleteResolver) Message() string {
	return d.DeleteHandler.Message
}

func (d DeleteResolver) Status() string {
	return d.DeleteHandler.Status
}

func (l LoginResolver) Token() string {
	return l.LoginHandler.Token
}

func (l LoginResolver) Id() string {
	return l.LoginHandler.Id
}

func (l LoginResolver) Message() string {
	return l.LoginHandler.Message
}

func (l LoginResolver) Status() string {
	return l.LoginHandler.Status
}

func (r ResponseResolver) Successmsg() string {
	return r.Response.Successmsg
}

func (r ResponseResolver) Errmsg() string {
	return r.Response.Errmsg
}

func (r ResponseResolver) Statuscode() string {
	return r.Response.Statuscode
}


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

func (a AudioResolver) Token() string {
	return a.Audio.Token
}

func (a AudioResolver) Message() string {
	return a.Audio.Message
}

func (a AudioResolver) Status() string {
	return a.Audio.Status
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

// func (a1 Audio1Resolver) Message() string {
// 	return a1.Audio1.Message
// }

// func (a1 Audio1Resolver) Status() string {
// 	return a.Audio1Resolver.Status
// }

func (p PersonResolver) FirstName() string {
	return p.Person.FirstName
}
func (p PersonResolver) LastName() string {
	return p.Person.LastName
}

func (p PersonResolver) Username() string {
	return p.Person.Username
}
func (p PersonResolver) EmailID() string {
	return p.Person.EmailID
}


func (p PersonResolver) Message() string {
	return p.Person.Message
}

func (p PersonResolver) Status() string {
	return p.Person.Status
}