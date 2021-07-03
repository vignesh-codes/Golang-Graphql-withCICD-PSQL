package personModel



// // Person data type
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	EmailID   string `db:"email"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Message string
	Status string
}

type Response struct {
	Successmsg string
	Errmsg string
	Statuscode string
}
type Audio struct {
	Token string
	Title string `db:"title"`
	Description  string `db:"description"`
	Category  string `db:"category"`
	CreatorName  string `db:"creatorname"`
	CreatorEmail  string `db:"creatoremail"`
	Message string
	Status string
}

type DeleteHandler struct {
	Message string
	Status string
}

type LoginHandler struct {
	Token string
	Id string
	Message string
	Status string
}


type Audio1 struct {
	Id string
	Title string 
	Description  string 
	Category  string 
	CreatorName  string 
	CreatorEmail  string 
}

type PersonInput struct {
	FirstName string
	LastName  string
	EmailID   string
	Password  string
	Username  string
}

type AudioInput struct {
	Token string
	Title string
	Description  string
	Category  string
	CreatorName  string
	CreatorEmail  string
}

