package main

import (
	"net/http"

	"github.com/christmas-fire/fruit-app/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	client := http.Client{}
	handler := controller.NewHandler(client)
	router := handler.InitRoutes()

	router.Static("/web/static", "./web/static")

	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	logrus.Println("Server is starting..")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logrus.Fatal(err)
	}
}
