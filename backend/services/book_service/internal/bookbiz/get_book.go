package bookbiz

import (
	"context"

	"supertruyen/entity/bookentity"
)

type GetBookRepo interface {
	GetBookDetail(ctx context.Context, id int) (*bookentity.Book, error)
}

type getBookBiz struct {
	repo GetBookRepo
}

func NewGetBookBiz(repo GetBookRepo) *getBookBiz {
	return &getBookBiz{repo: repo}
}

func (b *getBookBiz) Response(ctx context.Context, id int) (*bookentity.Book, error) {
	return b.repo.GetBookDetail(ctx, id)
}
