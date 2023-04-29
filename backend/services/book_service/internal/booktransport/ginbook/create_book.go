package ginbook

import (
	"fmt"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

func CreateBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(core.KeyRequester).(*clerk.User)
		fmt.Println(requester.ID)

		c.JSON(http.StatusOK, gin.H{"user": requester})
	}
}
