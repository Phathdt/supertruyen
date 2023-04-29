package bookrepo

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity/bookentity"
	"supertruyen/services/book_service/internal/bookmodel"
)

type BookStorage interface {
	GetBookById(ctx context.Context, id int) (*bookentity.Book, error)
	ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]bookentity.Book, error)
}

type repo struct {
	store BookStorage
}

func NewRepo(store BookStorage) *repo {
	return &repo{store: store}
}

func (r *repo) ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]bookentity.Book, error) {
	return r.store.ListBook(ctx, filter, paging)
}

func (r *repo) GetBookDetail(ctx context.Context, id int) (*bookentity.Book, error) {
	return r.store.GetBookById(ctx, id)
}
