package postgresql

import (
	"context"

	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
	"supertruyen/entity"
	"supertruyen/plugins/tracing"
	"supertruyen/services/chapter_service/internal/chaptermodel"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}
func (r *repo) ListChapter(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	db := r.db.Table(entity.Chapter{}.TableName())

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

func (r *repo) GetChapterById(ctx context.Context, id int) (*entity.Chapter, error) {
	_, span := tracing.StartTrace(ctx, "chapter", "repo.get")
	defer span.End()

	var data entity.Chapter

	if err := r.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}
