package middleware

import (
	"github.com/gin-gonic/gin"
	"lemon/app/controller"
	"lemon/models/user"
	"lemon/utils/errno"
	"lemon/utils/token"
	"lemon/utils/utils"
	"net/url"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 校验token合法性
		JWTpayload, err := token.ParseRequest(c)
		if err != nil {
			controller.SendResponse(c, errno.ErrToken, nil)
			c.Abort()
			return
		}

		// 校验sign合法性
		var params url.Values
		var sign string
		method := c.Request.Method
		debug := c.Query("debug")

		if method == "GET" {
			params = c.Request.URL.Query()
			debug = c.Query("debug")
			sign = c.Query("sign")
		} else if method == "POST" {
			c.Request.ParseForm()
			params = c.Request.PostForm
			debug = c.PostForm("debug")
			sign = c.PostForm("sign")
		}

		if debug != "1" && sign != utils.CreateSign(params) {
			controller.SendResponse(c, errno.ErrSign, nil)
			c.Abort()
			return
		}

		_, err = user.GetUserById(int(JWTpayload.UserId))
		if err != nil {
			controller.SendResponse(c, errno.ErrUserNotFound, err.Error())
			c.Abort()
			return
		}

		c.Set("UserId", JWTpayload.UserId)
		c.Set("Mobile", JWTpayload.Mobile)
		c.Set("NickName", JWTpayload.NickName)

		//// 更新最后操作时间
		//go model.UpdateLastSeen(JWTpayload.UserId)

		c.Next()
	}
}
