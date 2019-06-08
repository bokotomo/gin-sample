//
// デザインを扱うドメイン
//
package domain

type Design struct {
	id        uint
	title     string
	text      string
	good      int
	comments  int
	date      string
	thumbnail string
	exists    bool
}

func (this *Design) Set(id uint, title string, text string, good int, comments int, date string, thumbnail string) {
	this.id = id
	this.title = title
	this.text = text
	this.good = good
	this.comments = comments
	this.date = date
	this.thumbnail = thumbnail
	this.exists = true
}

func (this *Design) SetPickup(id uint, title string, thumbnail string) {
	this.id = id
	this.title = title
	this.text = ""
	this.good = 0
	this.comments = 0
	this.date = ""
	this.thumbnail = thumbnail
	this.exists = true
}

func (this *Design) NotExists() bool {
	return !this.exists
}

func (this *Design) Id() uint {
	return this.id
}

func (this *Design) Title() string {
	return this.title
}

func (this *Design) Text() string {
	return this.text
}

func (this *Design) Good() int {
	return this.good
}

func (this *Design) Comments() int {
	return this.comments
}

func (this *Design) Date() string {
	return this.date
}

func (this *Design) Thumbnail() string {
	return this.thumbnail
}
