package bookrepo

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity"
	"supertruyen/services/book_service/internal/bookmodel"
)

type BookStorage interface {
	GetBookById(ctx context.Context, id int) (*entity.Book, error)
	ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]entity.Book, error)
}

type repo struct {
	store BookStorage
}

func NewRepo(store BookStorage) *repo {
	return &repo{store: store}
}

func (r *repo) ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]entity.Book, error) {
	return r.store.ListBook(ctx, filter, paging)
}

func (r *repo) GetBookDetail(ctx context.Context, id int) (*entity.Book, error) {
	return r.store.GetBookById(ctx, id)
}
