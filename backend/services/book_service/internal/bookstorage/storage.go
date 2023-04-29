package bookstorage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
	"supertruyen/entity/bookentity"
	"supertruyen/services/book_service/internal/bookmodel"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{db: db}
}

func (s *storage) ListBook(ctx context.Context, filter *bookmodel.Filter, paging *core.Paging) ([]bookentity.Book, error) {
	var books []bookentity.Book

	db := s.db.Table(bookentity.Book{}.TableName())

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

func (s *storage) GetBookById(ctx context.Context, id int) (*bookentity.Book, error) {
	var book bookentity.Book

	if err := s.db.
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
