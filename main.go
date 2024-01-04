package main

import (
	"fmt"
	// "basicMethodTest"
	hello "GoLangStudy/Hello"
	// variable "GoLangStudy/Variable"
	"GoLangStudy/basicMethodTest"
	forMethod "GoLangStudy/src/forMethod"
)

func main() {
	fmt.Print("Hello world\n")

	basicMethodTest.Aaa()

	// hello pkg
	// var name string = hello.Input("hello PKG method")
	name := hello.Input("PKG method")
	fmt.Println("Goodnight: " + name)

	// 変数　テスト用
	// variable.Samplevariable()

	// srcファイル直下の関数群
	forMethod.ForRealSample()
}
