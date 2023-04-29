package entity

type Book struct {
	SQLModel
	Name string `json:"name"`
}

func (Book) TableName() string {
	return "books"
}
