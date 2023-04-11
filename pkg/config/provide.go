package config

import "go.uber.org/dig"

func Provide(c *dig.Container) error {
	configs := []any{
		ParseMySQL,
	}
	for _, config := range configs {
		if err := c.Provide(config); err != nil {
			return err
		}
	}

	return nil
}
