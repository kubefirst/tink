package main

import (
	"os"

	"github.com/kubefirst/tink/internal/cli"
)

func main() {
	if err := cli.NewAgent().Execute(); err != nil {
		os.Exit(-1)
	}
}
