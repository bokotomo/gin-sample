package domain

type Design struct {
	Id    uint   `json:"id"`
	Title string `json:"name"`
}

func NewDesign(id uint, title string) *Design {
	return &Design{
		Id:    id,
		Title: title,
	}
}
