package binding

import (
	"github.com/changmin888/protoc-gen-go-gin/pkg/errors"
	"github.com/gin-gonic/gin"
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
