package audioModel


type Audio struct {
	Token string
	Title string `db:"title"`
	Description  string `db:"description"`
	Category  string `db:"category"`
	CreatorName  string `db:"creatorname"`
	CreatorEmail  string `db:"creatoremail"`
	CreatedBy string `db:"createdby"`
	Destination string `db:"destination"`
	Message string
	Status string
}



//Defined a new struct to add ID field
type Audio1 struct {
	Id string
	Title string 
	Description  string 
	Category  string 
	CreatorName  string 
	CreatorEmail  string 
	CreatedBy string
	Destination string
	
}


//struct to take the input for creating audio
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