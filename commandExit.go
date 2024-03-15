package main

import (
	"os"
)

func commandExit(_ *config) error {
	os.Exit(0)
	return nil
}
