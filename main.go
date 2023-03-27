package main

import (
	"fmt"
	"os"

	"github.com/colbee1/stdcsv/cmd/stdcsv"
)

func main() {
	app := new(stdcsv.App)
	err := app.Setup()
	if err == nil {
		err = app.Run()
		app.Close() // Don't use defer because of os.Exit()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
