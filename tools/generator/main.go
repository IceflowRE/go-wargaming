package main

import (
	"fmt"
	"os"

	"github.com/IceflowRE/go-wargaming/v4/tools/generator/internal"
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
