package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/gittduse/protoc-gen-go-gin/pkg/errors"
	"net/http"
)

func Response(ctx *gin.Context, out interface{}, err error) {
	if err != nil {
		iErr := errors.FromError(err)
		ctx.JSON(int(iErr.Code), iErr)
	} else {
		ctx.JSON(http.StatusOK, out)
	}
}
