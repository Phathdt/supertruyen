package bookentity

import "supertruyen/entity"

type Book struct {
	entity.SQLModel
	Name string `json:"name"`
}

func (Book) TableName() string {
	return "books"
}
