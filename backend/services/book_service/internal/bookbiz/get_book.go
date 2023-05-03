package bookbiz

import (
	"context"

	"supertruyen/entity"
	"supertruyen/plugins/tracing"
)

const tracerID = "book"

type GetBookRepo interface {
	GetBookById(ctx context.Context, id int) (*entity.Book, error)
}

type getBookBiz struct {
	repo GetBookRepo
}

func NewGetBookBiz(repo GetBookRepo) *getBookBiz {
	return &getBookBiz{repo: repo}
}

func (b *getBookBiz) Response(ctx context.Context, id int) (*entity.Book, error) {
	ctx, span := tracing.StartTrace(ctx, "book", "biz.get")
	defer span.End()

	return b.repo.GetBookById(ctx, id)
}
