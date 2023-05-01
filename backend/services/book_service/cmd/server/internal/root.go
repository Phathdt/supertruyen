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
	"supertruyen/common"
	"supertruyen/plugins/appgrpc"
	"supertruyen/plugins/discovery/consul"
	"supertruyen/services/book_service/internal/booktransport/ginbook"
)

const (
	serviceName = "book-service"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("book service"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(consul.NewConsulComponent(common.KeyCompConsul, serviceName, version, 3000)),
		sctx.WithComponent(appgrpc.NewChapterClient(common.KeyCompChapterClient)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start book service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		publicRouter := router.Group("/api/books")
		{
			publicRouter.GET("", ginbook.ListBook(serviceCtx))
			publicRouter.GET("/:id", ginbook.GetBook(serviceCtx))
		}

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
