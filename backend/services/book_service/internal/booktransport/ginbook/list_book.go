package ginbook

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/services/book_service/internal/bookbiz"
	"supertruyen/services/book_service/internal/bookmodel"
	"supertruyen/services/book_service/internal/bookrepo"
	"supertruyen/services/book_service/internal/bookstorage"
)

func ListBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type reqParam struct {
			bookmodel.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		rp.Paging.Process()

		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		storage := bookstorage.NewStorage(db.GetDB())
		repo := bookrepo.NewRepo(storage)
		biz := bookbiz.NewListBookBiz(repo)

		books, err := biz.Response(c.Request.Context(), &rp.Filter, &rp.Paging)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(books, rp.Paging, rp.Filter))
	}
}
