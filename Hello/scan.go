package hello

import "fmt"

func scanDefault() {

	var int1 int
	fmt.Scan(&int1)

	var int2 int
	var int3 int
	fmt.Scan(&int2, &int3)

	fmt.Println(int1)
}
