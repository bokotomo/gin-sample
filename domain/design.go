//
// デザインを扱うドメイン
//
package domain

type Design struct {
	id    uint
	title string
}

func NewDesign(id uint, title string) *Design {
	return &Design{
		id:    id,
		title: title,
	}
}

func (this *Design) Id() uint {
	return this.id
}

func (this *Design) Title() string {
	return this.title
}
