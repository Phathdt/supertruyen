package internal

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/ginc"
	smdlw "github.com/viettranx/service-context/component/ginc/middleware"
	"github.com/viettranx/service-context/component/gormc"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"google.golang.org/grpc"
	"supertruyen/common"
	"supertruyen/plugins/clerkc"
	"supertruyen/plugins/discovery/consul"
	"supertruyen/plugins/gprc_server"
	"supertruyen/plugins/tracing"
	protos "supertruyen/proto/out/proto"
	"supertruyen/services/chapter_service/internal/chaptertransport/chaptergrpc"
	"supertruyen/services/chapter_service/internal/chaptertransport/ginchapter"
)

const (
	serviceName = "chapter-server"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("chapter service"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(clerkc.NewClerkComponent(common.KeyCompClerk)),
		sctx.WithComponent(consul.NewConsulComponent(common.KeyCompConsul, serviceName, version, 3000)),
		sctx.WithComponent(gprc_server.NewGprcServer(common.KeyCompGrpcServer)),
		sctx.WithComponent(tracing.NewTracingClient(common.KeyCompJaeger, serviceName)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start chapter service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		//clerkComp := serviceCtx.MustGet(common.KeyCompClerk).(common.ClerkComponent)
		grpcComp := serviceCtx.MustGet(common.KeyCompGrpcServer).(common.GrpcServer)
		grpcComp.SetRegisterHdl(func(server *grpc.Server) {
			protos.RegisterChapterServiceServer(server, chaptergrpc.NewChapterGrpcServer(serviceCtx))
		})

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx), otelgin.Middleware(serviceName))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		publicRouter := router.Group("/api/chapters")
		{
			publicRouter.GET("", ginchapter.ListChapter(serviceCtx))
			publicRouter.GET("/:id", ginchapter.GetBook(serviceCtx))
		}

		//protectedRoute := router.Group("/api/books", middleware.RequireAuth(clerkComp.GetClient()))

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
