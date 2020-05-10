package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "lemon/app/controller"
	"lemon/models/user"
	"lemon/utils/crypto"
	"lemon/utils/errno"
	"lemon/utils/logging"
	"lemon/utils/token"
	"lemon/utils/utils"
	"strconv"
	"time"
)

/**
 * @api {post} /user/login 登陆获取token
 * @apiName Login
 * @apiGroup User
 *
 * @apiParam {String} mobile 手机号码
 * @apiParam {String} password 登陆密码
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 */
func Login(c *gin.Context) {
	time.Sleep(5 * time.Second)
	var u user.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	d, err := user.GetUserByMobile(u.Mobile)
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
	logging.Info(t)
	SendResponse(c, nil, user.Token{Token: t})
}

func GetUserById(c *gin.Context) {
	fmt.Println(c.Get("UserId"))
	fmt.Println(c.Get("Mobile"))
	fmt.Println(c.Get("NickName"))
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	user, err := user.GetUserById(userId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

func GetUserList(c *gin.Context) {
	userList, err := user.GetUserList()
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}
	for k, v := range userList {
		fmt.Println(v.Id)
		fmt.Println(v.Mobile)
		fmt.Println(k, v)
	}
	SendResponse(c, nil, userList)
}

func UpdateUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	fmt.Println(utils.TimeStandar())
	user.UpdateUserById(userId, map[string]interface{}{
		"update_at": time.Now(),
	})

	SendResponse(c, nil, nil)
}

func DeleteUserById(c *gin.Context) {

}
