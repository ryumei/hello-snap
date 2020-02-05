package main

import (
	"fmt"
	"os"
)

func greeting(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s\n", greeting("world"))
	} else {
		for _, arg := range os.Args[1:] {
			fmt.Printf("%s\n", greeting(arg))
		}
	}
}
