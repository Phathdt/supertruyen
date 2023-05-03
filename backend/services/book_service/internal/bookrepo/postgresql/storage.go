package postgresql

import (
	"context"

	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
	"supertruyen/entity"
	"supertruyen/plugins/tracing"
	"supertruyen/services/book_service/internal/bookmodel"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

const tracerID = "book-repository-postgres"

func (r *repo) ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]entity.Book, error) {
	var books []entity.Book

	db := r.db.Table(entity.Book{}.TableName())

	//TODO update with filter

	// Count total records match conditions
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&books).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return books, nil
}

func (r *repo) GetBookById(ctx context.Context, id int) (*entity.Book, error) {
	_, span := tracing.StartTrace(ctx, "book", "repository.get")
	defer span.End()

	var book entity.Book

	if err := r.db.
		Table(book.TableName()).
		Where("id = ?", id).
		First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &book, nil
}
