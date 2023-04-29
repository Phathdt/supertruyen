package bookbiz

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity"
	"supertruyen/services/book_service/internal/bookmodel"
)

type ListBookRepo interface {
	ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]entity.Book, error)
}

type listBookBiz struct {
	repo ListBookRepo
}

func NewListBookBiz(repo ListBookRepo) *listBookBiz {
	return &listBookBiz{repo: repo}
}

func (b *listBookBiz) Response(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]entity.Book, error) {
	books, err := b.repo.ListBook(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return books, nil
}
