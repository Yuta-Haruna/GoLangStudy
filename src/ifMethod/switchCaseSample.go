package ifmethod

import (
	"fmt"
)

func switchCase01(a string) {
	fmt.Print("a = " + a)
	switch {
	case a == "a":
		fmt.Println("a")
	case a == "Hello":
		fmt.Println("こんにちは")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", a)
	}
}

func SwitchCaseExample(a string) {
	switchCase01("a")
	switchCase01("Hello")
	switchCase01("abc")
}
