package controller

import "go.uber.org/dig"

func Provide(container *dig.Container) error {
	controllers := []any{
		NewDemoController,
	}
	for _, controller := range controllers {
		if err := container.Provide(controller); err != nil {
			return err
		}
	}
	return nil
}
