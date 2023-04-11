package main

import (
	"github.com/vuho-pg/base/pkg/config"
	"github.com/vuho-pg/base/pkg/connection"
	"github.com/vuho-pg/base/pkg/repository"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	if err := config.Provide(container); err != nil {
		panic(err)
	}
	if err := connection.Provide(container); err != nil {
		panic(err)
	}
	if err := repository.Provide(container); err != nil {
		panic(err)
	}

}
