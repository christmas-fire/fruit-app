package controller

import (
	"net/http"

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
