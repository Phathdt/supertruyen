package bookbiz

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"supertruyen/entity"
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
	_, span := otel.Tracer(tracerID).Start(ctx, "Biz/Get")

	defer span.End()
	ctx = trace.ContextWithSpan(ctx, span)
	return b.repo.GetBookById(ctx, id)
}
