package main

import (
	"fmt"
	// "basicMethodTest"
	hello "GoLangStudy/Hello"
	"GoLangStudy/basicMethodTest"
)

func main() {
	fmt.Print("Hello world\n")

	basicMethodTest.Aaa()

	// hello pkg
	// var name string = hello.Input("hello PKG method")
	name := hello.Input("PKG method")
	fmt.Println("Goodnight: " + name)
}
