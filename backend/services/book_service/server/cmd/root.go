package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/ginc"
	smdlw "github.com/viettranx/service-context/component/ginc/middleware"
	"github.com/viettranx/service-context/component/gormc"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/middleware"
	"supertruyen/plugins/clerkc"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("book service"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompPostgres, "")),
		sctx.WithComponent(clerkc.NewClerkComponent(common.KeyClerk)),
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

		clerkComp := serviceCtx.MustGet(common.KeyClerk).(common.ClerkComponent)

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx))

		router.GET("/ping", middleware.RequireAuth(clerkComp.GetClient()), func(c *gin.Context) {
			fmt.Println(1111111)
			requester := c.MustGet(core.KeyRequester).(*clerk.User)
			c.JSON(http.StatusOK, gin.H{"data": "pong", "user": requester})
		})

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
