package chaptergrpc

import (
	"context"

	sctx "github.com/viettranx/service-context"
	"supertruyen/common"
	"supertruyen/plugins/tracing"
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
	ctx, span := tracing.WrapTraceIdFromIncomingContext(ctx, "chapter", "grpc.get")
	defer span.End()

	db := s.sc.MustGet(common.KeyCompGorm).(common.GormComponent)

	repo := postgresql.NewRepo(db.GetDB())

	_, err := repo.GetChapterById(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &protos.GetTotalChapterResponse{Total: 1}, nil
}
