package main

import (
	"context"
	"github.com/vuho-pg/base/pkg/config"
	"github.com/vuho-pg/base/pkg/connection"
	"github.com/vuho-pg/base/pkg/model"
	"github.com/vuho-pg/base/pkg/repository"
	"go.uber.org/dig"
	"gorm.io/gorm"
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
	ctx := context.Background()

	if err := container.Invoke(func(demoRepo repository.DemoRepository) {
		data := &model.Demo{
			AuthorModel: model.AuthorModel{
				CreatedBy: "vuhk",
				UpdatedBy: "vuhk",
			},
			Name:  "name",
			Value: "value",
		}
		if err := demoRepo.Create(ctx, data); err != nil {
			panic(err)
		}

		if err := demoRepo.Delete(ctx, &model.Demo{Model: gorm.Model{ID: 1}}); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}

}
