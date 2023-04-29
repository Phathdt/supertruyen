package entity

type Chapter struct {
	SQLModel
	BookId  int    `json:"book_id"`
	Content string `json:"content"`
}

func (Chapter) TableName() string {
	return "chapters"
}
