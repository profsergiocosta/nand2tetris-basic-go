package main

import (
	"nand2tetris-go/compiler/parser"
	"nand2tetris-go/vm"
	"os"
	"path"
)

func main() {

	fileName := os.Args[1]
	//parser.Interpret(fileName)
	ext := path.Ext(fileName)
	if ext == ".jack" {
		parser.Compile(fileName)
	} else {
		vm.Interpret(fileName)
	}

}
