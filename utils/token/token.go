package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lemon/utils/errno"
	"time"
)

type Context struct {
	UserId   int
	Mobile   string
	NickName string
}

/*{
"iss": "admin",          //该JWT的签发者
"iat": 1535967430,        //签发时间
"exp": 1535974630,        //过期时间
"nbf": 1535967430,         //该时间之前不接收处理该Token
"sub": "www.admin.com",   //面向的用户
"jti": "9f10e796726e332cec401c569969e13e"   //该Token唯一标识
}*/
func ParseToken(tokenString, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	})

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.UserId = int(claims["UserId"].(float64))
		ctx.Mobile = claims["Mobile"].(string)
		ctx.NickName = claims["NickName"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, errno.ErrMissingHeader
	}

	return ParseToken(header, secret)
}

func CreateToken(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":   c.UserId,
		"Mobile":   c.Mobile,
		"NickName": c.NickName,
		"nbf":      time.Now().Unix(),          // 如果当前时间在nbf时间之前，则Token不被接受；一般都会留一些余地，比如几分钟。
		"exp":      time.Now().Unix() + 3600*2, // token过期时间
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}
