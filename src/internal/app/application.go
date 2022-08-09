package app

import (
	"musement/src/internal/infrastructure/entrypoints/cmd"
)

func StartApp() error {
	cmd.Execute()
	return nil
}
