package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lemon/app/router"
	"lemon/config"
	"lemon/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Init("")
	gin.SetMode(viper.GetString("runmode"))
	models.DB.Init()
	defer models.DB.Close()

	engine := gin.New()

	srv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: router.Load(engine),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
