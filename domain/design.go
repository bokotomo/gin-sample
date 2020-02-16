package domain

// Design デザインを扱うドメイン
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

// Set is
func (d *Design) Set(id uint, title, text string, good, comments int, date, thumbnail string) {
	d.id = id
	d.title = title
	d.text = text
	d.good = good
	d.comments = comments
	d.date = date
	d.thumbnail = thumbnail
	d.exists = true
}

// SetPickup is
func (d *Design) SetPickup(id uint, title string, good int, thumbnail string) {
	d.id = id
	d.title = title
	d.text = ""
	d.good = good
	d.comments = 0
	d.date = ""
	d.thumbnail = thumbnail
	d.exists = true
}

// NotExists is
func (d *Design) NotExists() bool {
	return !d.exists
}

// Id is
func (d *Design) Id() uint {
	return d.id
}

// Title is
func (d *Design) Title() string {
	return d.title
}

// Text is
func (d *Design) Text() string {
	return d.text
}

// Good is
func (d *Design) Good() int {
	return d.good
}

// Comments is
func (d *Design) Comments() int {
	return d.comments
}

// Date is
func (d *Design) Date() string {
	return d.date
}

// Thumbnail is
func (d *Design) Thumbnail() string {
	return d.thumbnail
}
