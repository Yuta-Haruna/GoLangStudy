package forMethod

import (
	"fmt"
	"strconv"
)

// GoにWhileは存在しない
func WhileSample() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(strconv.Itoa(sum) + " : GoにWhileは存在しない")
}

func ForRealSample() {
	max := 10
	ans := 0
	for i := 0; i < max; i++ {
		ans += i
	}
	fmt.Println(strconv.Itoa(ans) + " for int i = 0 ; i < sum; i++ {ans += i} ")
}
