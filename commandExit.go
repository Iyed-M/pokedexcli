package main

import (
	"os"
)

func commandExit(_ *config, _ string) error {
	os.Exit(0)
	return nil
}
