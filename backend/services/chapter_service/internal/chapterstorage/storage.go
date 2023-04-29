package chapterstorage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
	"supertruyen/entity"
	"supertruyen/services/chapter_service/internal/chaptermodel"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{db: db}
}
func (s *storage) ListChapter(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	db := s.db.Table(entity.Chapter{}.TableName())

	if filter.BookId != nil {
		db = db.Where("book_id = ?", filter.BookId)
	}

	// Count total records match conditions
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&chapters).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return chapters, nil
}
