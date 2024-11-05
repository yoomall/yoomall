package authmiddleware

import (
	"time"

	"yoomall/modules/auth/model"
	"yoomall/yoo/driver"
	"yoomall/yoo/helper/response"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareGroup struct {
	AuthMiddleware             gin.HandlerFunc
	MustAuthMiddleware         gin.HandlerFunc
	MustAuthMiddlewareWithUser gin.HandlerFunc
}

func NewAuthMiddlewareGroup(db *driver.DB) *AuthMiddlewareGroup {
	return &AuthMiddlewareGroup{
		AuthMiddleware:             NewAuthMiddleware(db, false, false),
		MustAuthMiddleware:         NewMustAuthMiddleware(db),
		MustAuthMiddlewareWithUser: NewMustAuthMiddlewareWithUser(db),
	}
}

func NewMustAuthMiddleware(db *driver.DB) gin.HandlerFunc {
	return NewAuthMiddleware(db, true, false)
}

func NewMustAuthMiddlewareWithUser(db *driver.DB) gin.HandlerFunc {
	return NewAuthMiddleware(db, true, true)
}

func NewAuthMiddleware(db *driver.DB, must bool, needUser bool) gin.HandlerFunc {
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

		c.Set("token", userToken)
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
