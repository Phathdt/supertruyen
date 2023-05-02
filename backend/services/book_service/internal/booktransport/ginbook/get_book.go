package ginbook

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"supertruyen/common"
	"supertruyen/plugins/appgrpc"
	"supertruyen/services/book_service/internal/bookbiz"
	"supertruyen/services/book_service/internal/bookrepo/postgresql"
)

func GetBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := otel.Tracer("book").Start(c.Request.Context(), "Transport/Get")
		defer span.End()
		ctx := trace.ContextWithSpan(c.Request.Context(), span)

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

		chapter, err := client.GetTotalChapter(ctx, id)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		fmt.Println(chapter)
		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
