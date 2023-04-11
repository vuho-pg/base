package repository

import "go.uber.org/dig"

func Provide(c *dig.Container) error {
	repos := []any{
		NewDemoRepository,
	}
	for _, repo := range repos {
		if err := c.Provide(repo); err != nil {
			return err
		}
	}
	return nil
}
