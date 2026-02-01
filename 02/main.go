package main

import (
	"fmt"
	"os"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main ls <args>")
		os.Exit(1)
	}

	var path string

	//TODO: fix number of args
	path = os.Args[2]

	switch os.Args[1] {

	case "ls":
		dir, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("Error %s\n", err)
			os.Exit(1)
		}
		for _, file := range dir {

			info, _ := file.Info()

			//TODO: format ModTime

			fmt.Println(info.Mode().Perm().String(), info.Size(), info.ModTime(), info.Name())
		}

	default:
		fmt.Println("expected 'ls' command. bye")
		os.Exit(1)
	}

}
