package main

import "github.com/raihankhan/golab/functionset"

func main() {
	functionset.NewFunctionSet([]func(...interface{}){
		func() {

		},
	}...)
}
