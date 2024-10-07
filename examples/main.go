package main

import (
	"log"

	"github.com/techforge-lat/linkit"
	"github.com/techforge-lat/linkit/examples/user"
)

func main() {
	dependencyContainer, err := BuildDependencies()
	if err != nil {
		log.Fatal(err)
	}

	userHandler, _ := linkit.Get[user.Handler](dependencyContainer, linkit.DependencyName("user.handler"))

	userHandler.Create()
}
