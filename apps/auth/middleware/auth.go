package authmiddleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/apps/auth/model"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

func AuthMiddleware(db *driver.DB, must bool, needUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		print("auth middleware")
		token := c.GetHeader("Token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			if must {
				response.Error(response.ErrNotAuthorized, "token 不存在").Done(c)
				c.Abort()
				return
			}
			c.Next()
			return
		}
		var userToken model.UserToken
		if err := db.Where("token = ?", token).First(&userToken).Error; err != nil {
			response.Error(response.ErrNotAuthorized, "token 不可用").Done(c)
			c.Abort()
			return
		}
		if userToken.ExpireTime.Before(time.Now()) {
			response.Error(response.ErrNotAuthorized, "token 已过期").Done(c)
			c.Abort()
			return
		}

		c.Set("userId", userToken.UserId)

		// 不需要用户具体的信息
		if !needUser {
			c.Next()
			return
		}

		var user model.User
		if err := db.Where("id = ?", userToken.UserId).First(&user).Error; err != nil {
			response.Error(response.ErrNotAuthorized, "用户不存在").Done(c)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
