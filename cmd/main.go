package main

import (
	"net/http"
	"study/lemon/config"
	"study/lemon/app/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"study/lemon/models"
)

func main() {
	if err := config.Init(""); err != nil{
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	models.DB.Init()
	defer models.DB.Close()

	engine := gin.New()

	router.Load(
		engine,
	)

	http.ListenAndServe(viper.GetString("addr"), engine)
}
