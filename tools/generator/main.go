package main

import (
	"fmt"
	"github.com/IceflowRE/go-wargaming/v3/tools/generator/internal"
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
