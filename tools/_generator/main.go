package main

import (
	"fmt"
	"github.com/IceflowRE/go-wargaming/v4/tools/_generator/internal"
	"os"
)

func main() {
	generator, err := internal.NewGenerator(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = generator.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
