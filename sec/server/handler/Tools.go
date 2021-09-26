package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meowalien/rabbitgather-article/sec/logger"
	"net/http"
)

type ExpectErrorTypes string

const (
	WrongFormat = "WrongFormat"
)


// ErrorCheck Check if the error is the given errorType, if so return
// the error message and abort, if not, return false and abort.
// will return true if the error is nil
func ErrorCheck(err error, c *gin.Context, errorType ExpectErrorTypes , extraInfo ...interface{}) bool {
	extraInfoString := fmt.Sprint(extraInfo...)
	if err == nil {
		return true
	} else {
		switch errorType {
		case WrongFormat:
			c.AbortWithStatusJSON(http.StatusBadRequest, StandardJsonResponse{
				Type: ErrorResponse,
				Msg:  fmt.Sprint("wrong format, please check the input.",extraInfoString),
			})
			logger.Logger.Info("wrong format input: ",err)
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
