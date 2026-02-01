package main

import (
	"fmt"
	"os"
)

func main() {

	//args := os.Args[1:]

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main ls <args>")
		os.Exit(1)
	}

	path := os.Args[2]

	switch os.Args[1] {

	case "ls":
		dir, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("Error %s\n", err)
			os.Exit(1)
		}
		for _, file := range dir {
			fmt.Println(file)
		}

	default:
		fmt.Println("expected 'ls' command. bye")
		os.Exit(1)
	}

}
