package user

import (
	"github.com/gin-gonic/gin"
	"study/lemon/models"
	"study/lemon/utils/errno"
	. "study/lemon/app/controller"
	"study/lemon/utils/token"
	"study/lemon/utils/crypto"
	"strconv"
	"study/lemon/utils/logging"
)

// @Summary 登陆获取token
// @Description 登陆获取token
// @Tags user
// @Accept  json
// @Produce  json
// @Param mobile path string true "Mobile"
// @Param password path string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var u models.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	d, err := models.GetUserByMobile(u.Mobile)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := crypto.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	ctx := token.Context{UserId: d.Id, Mobile: d.Mobile, NickName: d.NickName}
	t, err := token.CreateToken(c, ctx, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	logging.Error(t)
	SendResponse(c, nil, models.Token{Token: t})
}

func GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	user, err := models.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

func GetUserList(c *gin.Context) {

}

func UpdateUserById(c *gin.Context) {

}

func DeleteUserById(c *gin.Context) {

}
