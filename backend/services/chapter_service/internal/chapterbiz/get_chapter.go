package chapterbiz

import (
	"context"

	"supertruyen/entity"
)

type GetChapterRepo interface {
	GetChapterDetail(ctx context.Context, id int) (*entity.Chapter, error)
}

type getChapterBiz struct {
	repo GetChapterRepo
}

func NewGetChapterBiz(repo GetChapterRepo) *getChapterBiz {
	return &getChapterBiz{repo: repo}
}

func (b *getChapterBiz) Response(ctx context.Context, id int) (*entity.Chapter, error) {
	return b.repo.GetChapterDetail(ctx, id)
}
