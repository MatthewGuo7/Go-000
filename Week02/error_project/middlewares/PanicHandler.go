package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PanicHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e:= recover(); e != nil {
				fmt.Printf("panic: %+v", e )
				context.JSON(http.StatusBadRequest, gin.H{"msg":"请重试"})
			}
		}()
		context.Next()
	}
}
