package app

import (
	"net/http"

	"github.com/christmas-fire/fruit-app/internal/controller"
	"github.com/sirupsen/logrus"
)

func Run() error {
	s := http.Server{
		Addr: ":8080",
	}

	client := http.Client{}
	handler := controller.NewHandler(client)
	router := handler.InitRoutes()

	logrus.Println("Server is starting..")
	if err := http.ListenAndServe(s.Addr, router); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
