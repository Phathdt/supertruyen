package bookbiz

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity/bookentity"
	"supertruyen/services/book_service/internal/bookmodel"
)

type ListBookRepo interface {
	ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]bookentity.Book, error)
}

type listBookBiz struct {
	repo ListBookRepo
}

func NewListBookBiz(repo ListBookRepo) *listBookBiz {
	return &listBookBiz{repo: repo}
}

func (b *listBookBiz) Response(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]bookentity.Book, error) {
	books, err := b.repo.ListBook(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return books, nil
}
