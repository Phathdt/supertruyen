package ginbook

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"supertruyen/common"
	"supertruyen/services/book_service/internal/bookbiz"
	"supertruyen/services/book_service/internal/bookrepo"
	"supertruyen/services/book_service/internal/bookstorage"
)

func GetBook(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		db := sc.MustGet(common.KeyCompGorm).(common.GormComponent)

		storage := bookstorage.NewStorage(db.GetDB())
		repo := bookrepo.NewRepo(storage)
		biz := bookbiz.NewGetBookBiz(repo)

		book, err := biz.Response(c.Request.Context(), id)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
