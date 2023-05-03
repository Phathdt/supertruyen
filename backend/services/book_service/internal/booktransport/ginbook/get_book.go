package ginbook

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/plugins/appgrpc"
	"supertruyen/plugins/tracing"
	"supertruyen/services/book_service/internal/bookbiz"
	"supertruyen/services/book_service/internal/bookrepo/postgresql"
)

func GetBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracing.StartTrace(ctx, "book", "transport.get")
		defer span.End()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		client := sc.MustGet(common.KeyCompChapterClient).(appgrpc.ChapterClient)
		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		repo := postgresql.NewRepo(db.GetDB())
		biz := bookbiz.NewGetBookBiz(repo)

		book, err := biz.Response(ctx, id)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		_, err = client.GetTotalChapter(ctx, id)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
