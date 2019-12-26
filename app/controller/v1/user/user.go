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

func Test(c *gin.Context) {
	// Query			PostForm			获取key对应的值，不存在为空字符串
	// GetQuery			GetPostForm			多返回一个key是否存在的结果
	// QueryArray		PostFormArray		获取key对应的数组，不存在返回一个空数组
	// GetQueryArray	GetPostFormArray	多返回一个key是否存在的结果
	// QueryMap			PostFormMap			获取key对应的map，不存在返回空map
	// GetQueryMap		GetPostFormMap		多返回一个key是否存在的结果
	// DefaultQuery		DefaultPostForm		key不存在的话，可以指定返回的默认值

	// queryArray := c.QueryArray("media")
	queryMap := c.QueryMap("ids")
	SendResponse(c, nil, queryMap)
}

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
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	user, err := user.GetUserById(uint64(userId))
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

	fmt.Println(utils.GetTimeStandar())
	user.UpdateUserById(uint64(userId), map[string]interface{}{
		"update_at": time.Now(),
	})

	SendResponse(c, nil, nil)
}

func DeleteUserById(c *gin.Context) {

}
