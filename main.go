package main

import (
	"os"
	"strings"

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

	lines := strings.Split(file, "\n")

	p := parser.Parser{
		Lines: lines,
	}
	p.Parse()
}
