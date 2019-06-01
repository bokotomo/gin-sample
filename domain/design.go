package domain

type Design struct {
	Id    int `json:"id"`
	Title string `json:"name"`
}

func NewDesign(id int, title string) *Design {
	return &Design{
		Id:    id,
		Title: title,
	}
}
