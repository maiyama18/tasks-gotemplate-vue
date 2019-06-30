package main

import (
	"fmt"
	"os"

	"github.com/maiyama18/tasks-gotemplate-vue/infra/router"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := router.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		return exitCodeErr
	}
	return exitCodeOK
}
