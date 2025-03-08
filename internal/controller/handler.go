package controller

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Client http.Client
}

func NewHandler(client http.Client) *Handler {
	return &Handler{
		Client: client,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Разрешаем запросы только с http://localhost:3000
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		fruits := api.Group("/fruits")
		{
			fruits.GET("/all", h.getAllFruits)
			fruits.GET("/name", h.getFruitsByName)
			fruits.GET("/family", h.getFruitsByFamily)
			fruits.GET("/genus", h.getFruitsByGenus)
			fruits.GET("/order", h.getFruitsByOrder)
		}
	}

	return router
}
