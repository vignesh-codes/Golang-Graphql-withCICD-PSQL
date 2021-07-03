package audioModel
// // Person data type


type Audio struct {
	Token string
	Title string `db:"title"`
	Description  string `db:"description"`
	Category  string `db:"category"`
	CreatorName  string `db:"creatorname"`
	CreatorEmail  string `db:"creatoremail"`
	CreatedBy string `db:"createdby"`
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
	CreatedBy string
}



type AudioInput struct {
	Token string
	Title string
	Description  string
	Category  string
	CreatorName  string
	CreatorEmail  string
}

type DeleteHandler struct {
	Message string
	Status string
}