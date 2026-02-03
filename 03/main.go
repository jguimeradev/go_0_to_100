package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type env struct {
	db_host            string
	db_port            string
	db_name            string
	db_user            string
	db_password        string
	db_max_connections string
	db_ssl_mode        string
}

func readEnvFile(line string) {

	fmt.Println("line")

}

func assignEnvVars(v env) {

	db := make(map[string]string)

	db["db_host"] = v.db_host
	db["db_port"] = v.db_port
	db["db_name"] = v.db_name
	db["db_password"] = v.db_password
	db["db_max_connections"] = v.db_max_connections
	db["db_ssl_mode"] = v.db_ssl_mode
}

func main() {
	//read file

	path := filepath.Ext(".env")
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f) //scan lines by default

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") {
			continue
		}
		readEnvFile(line)
	}

}
