package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: clean-migrate-core <arg>")
		return
	}
	fmt.Println(os.Args[1])
}
