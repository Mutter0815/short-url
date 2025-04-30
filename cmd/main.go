package main

import (
	"log"
	"net/http"

	"github.com/Mutter0815/short-url-go/configs"
	"github.com/Mutter0815/short-url-go/internal/handlers"
	"github.com/Mutter0815/short-url-go/internal/repository"
	"github.com/Mutter0815/short-url-go/internal/service"
	"github.com/Mutter0815/short-url-go/internal/storage/database"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := configs.Load()
	db := database.Connect(&cfg.DB)

	err := db.Ping()
	if err != nil {
		log.Fatalf("Ошибка соеднинениия с сервером %v", err)
	}

	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(*linkRepo)
	LinkHandler := handlers.NewLinkHandler(linkService)

	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	r.GET("/orginalurl", LinkHandler.GetOriginalURLByShortLink)
	r.GET("/shortlink", LinkHandler.GetShortLinkByOriginalLink)
	r.POST("link", LinkHandler.CreateLink)
	r.GET("/:short_code", LinkHandler.RedirectLink)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("ОШибка запуска сервера %v", err)
	}

}
