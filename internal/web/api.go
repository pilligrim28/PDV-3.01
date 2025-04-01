package web

import (
	"dr600-server/internal/models"
"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/mod/sumdb/storage"
)

func SetupAPI(r *gin.Engine, db storage.Storage) {
    r.GET("/repeaters", func(c *gin.Context) {
        // Получение списка ретрансляторов
    })

    r.POST("/repeaters", func(c *gin.Context) {
        // Добавление нового ретранслятора
        var repeater models.Repeater
        c.BindJSON(&repeater)
        db.Create(&repeater)
    })
}