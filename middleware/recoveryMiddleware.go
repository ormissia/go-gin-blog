// @File: recoveryMiddleware
// @Date: 2020/12/11 17:49
// @Author: 安红豆
// @Description: 统一错误处理

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/util"
)

func ServerError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Error(ctx, fmt.Sprint(err), nil)
				return
			}
		}()
		ctx.Next()
	}
}
