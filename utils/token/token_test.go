package token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCreateToken(t *testing.T) {
	ctx := Context{
		UserId:   1,
		Mobile:   "18800001111",
		NickName: "1",
	}
	fmt.Println(CreateToken(&gin.Context{}, ctx, ""))
}
