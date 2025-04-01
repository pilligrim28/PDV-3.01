package main

import (
	"dr600-server/internal/sip"
	"dr600-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация хранилища
	storage := storage.New()

	// Запуск SIP сервера
	sipServer := sip.NewServer(":5060")
	go sipServer.Start()

	// Настройка веб-сервера
	r := gin.Default()
	setupRoutes(r, storage)
	r.Run(":8080")
}

func setupRoutes(r *gin.Engine, storage *storage.Storage) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "DR600 Server")
	})
}