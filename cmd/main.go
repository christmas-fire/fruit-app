package main

import (
	"net/http"

	"github.com/christmas-fire/fruit-app/internal/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	client := http.Client{}
	handler := controller.NewHandler(client)

	logrus.Println("Server is starting..")
	if err := http.ListenAndServe(":8080", handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
