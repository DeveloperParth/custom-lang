package main

import (
	"os"

	"github.com/developerparth/my-own-lang/parser"
)

func getFile() string {
	data, err := os.ReadFile("./my.lang")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	file := getFile()

	p := parser.Parser{}
	p.Parse(file)
}
