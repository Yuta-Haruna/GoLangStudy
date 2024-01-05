package ifmethod

import (
	"fmt"
	"strconv"
)

func ifSample01(x, y int) string {
	var ans string = ""
	if x < 0 && y < 0 {
		return "両方とも小さいではないか"
	} else if x >= 0 {
		ans += strconv.Itoa(x) + " : は0以上"
		if y >= 0 {
			ans += "　かつ　" + strconv.Itoa(y) + " : も0以上やで"
		}
	} else if y >= 0 {
		return strconv.Itoa(y) + " : は0以上(yのみ)"
	}
	return "x = " + strconv.Itoa(x) + " : y = " + strconv.Itoa(x)
}

func IfExample() {
	fmt.Println(ifSample01(-4, 10))
}
