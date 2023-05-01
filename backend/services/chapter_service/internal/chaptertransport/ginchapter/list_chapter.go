package ginchapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/services/chapter_service/internal/chapterbiz"
	"supertruyen/services/chapter_service/internal/chaptermodel"
	"supertruyen/services/chapter_service/internal/chapterrepo/postgresql"
)

func ListChapter(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type reqParam struct {
			chaptermodel.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		repo := postgresql.NewRepo(db.GetDB())
		biz := chapterbiz.NewListChapterBiz(repo)

		chapters, err := biz.Response(c.Request.Context(), &rp.Filter, &rp.Paging)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(chapters, rp.Paging, rp.Filter))
	}
}
