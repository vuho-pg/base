package service

import "go.uber.org/dig"

func Provide(container *dig.Container) error {
	services := []any{
		NewDemoService,
	}
	for _, service := range services {
		if err := container.Provide(service); err != nil {
			return err
		}
	}
	return nil
}
