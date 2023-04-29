package chapterbiz

import (
	"context"

	"github.com/viettranx/service-context/core"
	"supertruyen/entity"
	"supertruyen/services/chapter_service/internal/chaptermodel"
)

type ListChapterRepo interface {
	ListChapter(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error)
}

type listChapterBiz struct {
	repo ListChapterRepo
}

func NewListChapterBiz(repo ListChapterRepo) *listChapterBiz {
	return &listChapterBiz{repo: repo}
}

func (b *listChapterBiz) Response(ctx context.Context, filter *chaptermodel.Filter, paging *core.Paging) ([]entity.Chapter, error) {
	return b.repo.ListChapter(ctx, filter, paging)
}
