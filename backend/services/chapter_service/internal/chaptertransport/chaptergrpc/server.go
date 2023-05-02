package chaptergrpc

import (
	"context"
	"fmt"

	sctx "github.com/viettranx/service-context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"supertruyen/common"
	protos "supertruyen/proto/out/proto"
	"supertruyen/services/chapter_service/internal/chapterrepo/postgresql"
)

type chapterGrpcServer struct {
	sc sctx.ServiceContext
}

func NewChapterGrpcServer(sc sctx.ServiceContext) *chapterGrpcServer {
	return &chapterGrpcServer{sc: sc}
}

func (s *chapterGrpcServer) GetTotalChapter(ctx context.Context, request *protos.GetTotalChapterRequest) (*protos.GetTotalChapterResponse, error) {
	_, span := otel.Tracer("chapter-grpc").Start(ctx, "grpc/Get")
	defer span.End()
	ctx = trace.ContextWithSpan(ctx, span)
	db := s.sc.MustGet(common.KeyCompGorm).(common.GormComponent)

	repo := postgresql.NewRepo(db.GetDB())

	chapter, err := repo.GetChapterById(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	fmt.Println(chapter)

	return &protos.GetTotalChapterResponse{Total: 1}, nil
}
