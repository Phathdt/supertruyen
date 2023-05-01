package appgrpc

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	sctx "github.com/viettranx/service-context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	protos "supertruyen/proto/out/proto"
)

type ChapterClient interface {
	GetTotalChapter(ctx context.Context, id int) (int, error)
}

type chapterClient struct {
	id         string
	prefix     string
	consulHost string
	url        string
	logger     sctx.Logger
	client     protos.ChapterServiceClient
}

func NewChapterClient(id string) *chapterClient {
	return &chapterClient{
		id:     id,
		prefix: id,
	}
}

func (c *chapterClient) ID() string {
	return c.id
}

func (c *chapterClient) InitFlags() {
	flag.StringVar(&c.consulHost, "grpc_consul_host", "localhost:8500", "consult host, should be localhost:8500")
}

func (c *chapterClient) Activate(sc sctx.ServiceContext) error {
	c.logger = sc.Logger(c.id)

	c.logger.Infoln("Setup chapter client service:", c.prefix)

	target := fmt.Sprintf("consul://%s/%s?healthy=true", c.consulHost, "chapter-server")
	conn, err := grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	c.client = protos.NewChapterServiceClient(conn)

	c.logger.Infof("Setup chapter client service success with url %s", target)

	return nil
}

func (c *chapterClient) Stop() error {
	c.logger.Infoln("chapterClient grpc service stopped")
	return nil
}

func (c *chapterClient) GetTotalChapter(ctx context.Context, id int) (int, error) {
	rs, err := c.client.GetTotalChapter(ctx, &protos.GetTotalChapterRequest{Id: int32(id)})
	if err != nil {
		return 0, err
	}

	return int(rs.Total), nil
}