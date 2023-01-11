package domain
type Teacher struct {
    ID int `sql:"primary key"`
    Name string
}

type Student struct {
    ID int `sql:"primary key"`
    Title string
    Body string
    UserID int `sql:"type:integer REFERENCES users(id)"`
}
