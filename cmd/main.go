package main

import (
	"os"

	"github.com/nezutero/shinjiru/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
