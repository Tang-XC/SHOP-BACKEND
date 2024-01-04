package middleware

import (
	"github.com/gin-gonic/gin"
	"shop/pkg/common"
	utils "shop/pkg/utils/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从header中获取token
		token := c.Request.Header.Get("Authorization")
		//验证token
		_, err := utils.ParseToken(token)
		if err != nil {
			common.SuccessResponse(c, "验证失败，请先登录")
			c.Abort()
			return
		}
		c.Next()
	}
}
