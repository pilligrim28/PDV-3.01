package main

import (
	"dr600-server/internal/config"
	"dr600-server/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status":  "ok",
            "version": cfg.DR600.BaseURL,
        })
    })
    
    r.Run(":" + cfg.Server.Port)
}

func setupRoutes(r *gin.Engine, storage *storage.Storage) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "DR600 Server")
	})
}