package ginbook

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/plugins/appgrpc"
	"supertruyen/services/book_service/internal/bookbiz"
	"supertruyen/services/book_service/internal/bookrepo/postgresql"
)

func GetBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		client := sc.MustGet(common.KeyCompChapterClient).(appgrpc.ChapterClient)
		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		repo := postgresql.NewRepo(db.GetDB())
		biz := bookbiz.NewGetBookBiz(repo)

		book, err := biz.Response(c.Request.Context(), id)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		chapter, err := client.GetTotalChapter(c.Request.Context(), id)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		fmt.Println(chapter)
		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
