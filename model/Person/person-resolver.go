package personModel
//Resolvers take care of the response fields for graphql schema

type PersonResolver struct {
	Person
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