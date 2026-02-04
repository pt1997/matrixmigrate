package main

import (
	"os"

	"github.com/aligundogdu/matrixmigrate/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}



