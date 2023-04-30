package chapterrepo

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity"
	"supertruyen/services/chapter_service/internal/chaptermodel"
)

type ChapterStorage interface {
	ListChapter(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error)
	GetChapterDetail(ctx context.Context, id int) (*entity.Chapter, error)
}

type repo struct {
	storage ChapterStorage
}

func NewRepo(storage ChapterStorage) *repo {
	return &repo{storage: storage}
}

func (r *repo) ListChapter(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error) {
	return r.storage.ListChapter(ctx, filter, paging)
}

func (r *repo) GetChapterDetail(ctx context.Context, id int) (*entity.Chapter, error) {
	return r.storage.GetChapterDetail(ctx, id)
}
