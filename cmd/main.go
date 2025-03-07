package main

import (
	"github.com/christmas-fire/fruit-app/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
