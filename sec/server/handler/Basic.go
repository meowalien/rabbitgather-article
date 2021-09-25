package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meowalien/rabbitgather-article/global"
	"net/http"
)

// 文章基本的增刪查改
type Basic struct {
}

type ExpectErrorTypes string

const (
	WrongFormat = "WrongFormat"
)

type StandardResponseTypes string

const (
	ErrorResponse StandardResponseTypes = "err"
)

type StandardJsonResponse struct {
	Type StandardResponseTypes `json:"type"`
	Msg  string                `json:"msg,omitempty"`
	Obj  interface{}           `json:"obj,omitempty"`
}

// ErrorCheck Check if the error is the given errorType, if so return
// the error message and abort, if not, return false and abort.
// will return true if the error is nil
func ErrorCheck(err error, c *gin.Context, errorType ExpectErrorTypes) bool {
	if err == nil {
		return true
	} else {
		switch errorType {
		case WrongFormat:
			c.AbortWithStatusJSON(http.StatusBadRequest, StandardJsonResponse{
				Type: ErrorResponse,
				Msg:  "wrong format, please check the input.",
			})
			global.Logger.Info("wrong format input: ",err)
			return true
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, StandardJsonResponse{
				Type: ErrorResponse,
				Msg:  "unknown error",
			})
			panic(fmt.Sprintf("The handling method of errorType \"%s\" has not been defined",errorType))
		}

	}
}

func (receiver Basic) Get(ctx *gin.Context) {
	type request struct {
		ID int `form:"id" binding:"required"`
	}
	var req request
	err := ctx.ShouldBindQuery(&req)

	if !ErrorCheck(err,ctx,WrongFormat) {
		return
	}
	global.Logger.Debug("req: ", req)
}
func (receiver Basic) POST(ctx *gin.Context) {

}
func (receiver Basic) DELETE(ctx *gin.Context) {

}
func (receiver Basic) PATCH(ctx *gin.Context) {

}
