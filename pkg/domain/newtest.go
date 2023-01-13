package domain

type Teacher struct {
	ID   int
	Name string
}

type Student struct {
	ID      int `sql:"primary key"`
	Title   string
	Body    string
	UserID  int
	Teacher Teacher `gorm:"foreignkey:UserID"`
}
