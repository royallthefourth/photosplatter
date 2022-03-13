package main

import (
	"errors"
	"os"
)

func initPath(p string) error {
	if p == "" {
		return errors.New("p flag must be nonempty")
	}

	return os.Chdir(p)
}
