package main

import (
	"fmt"
	"os"
	"strings"
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

func formatModTime(modTime string) string {

	m := strings.Split(modTime, " ")

	b, _, _ := strings.Cut(m[1], ".")

	m[1] = b

	modified := strings.Join(m[:2], " ")

	return modified
}

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

			fmt.Println(info.Mode().Perm().String(), info.Size(), formatModTime(info.ModTime().String()), info.Name())
		}

	default:
		fmt.Println("expected 'ls' command. bye")
		os.Exit(1)
	}

}
