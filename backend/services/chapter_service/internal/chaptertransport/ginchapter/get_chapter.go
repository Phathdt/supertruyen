package ginchapter

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/services/chapter_service/internal/chapterbiz"
	"supertruyen/services/chapter_service/internal/chapterrepo"
	"supertruyen/services/chapter_service/internal/chapterstorage"
)

func GetBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		storage := chapterstorage.NewStorage(db.GetDB())
		repo := chapterrepo.NewRepo(storage)
		biz := chapterbiz.NewGetChapterBiz(repo)

		book, err := biz.Response(c.Request.Context(), id)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
