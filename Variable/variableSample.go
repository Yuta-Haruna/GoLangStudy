package variable

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 下記記載方法も許容している
// import "fmt"
// import "strconv"

func Samplevariable() {
	a, b, c := 100, 200, 300
	// 下記は結合できない
	// fmt.Println(a + b + c + " : tatal Int")

	fmt.Println(a + b + c)

	samplevariable()

	sampleConst()

	sampleSwitch()

}

/*
下記メソッドはメソッド名の先頭が小文字のため内部でしか使えない
*/
func samplevariable() {

	fmt.Println("samplevariable を利用している")

	// 文字列⇒整数
	var text string = "123"
	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("variable ErrorMSG")
		return
	}

	// 整数⇒文字列
	text2 := strconv.Itoa(num)

	fmt.Println("実行結果 ：" + text2)
}

/*
下記メソッドはメソッド名の先頭が小文字のため内部でしか使えない
*/
func sampleConst() {

	fmt.Println("sampleConst を利用している")

	// 型宣言はまだしていない
	const max_size = 999

	// 型宣言済み
	const min_size int = 1

	fmt.Println("max_size 実行結果 ：", max_size)
	// 型宣言を今実行している
	fmt.Println("max_size*1.1 実行結果 ：", max_size*1.1)
	fmt.Println("min_size 実行結果 ：", min_size)
	fmt.Println("min_size*2 実行結果 ：", min_size*2)
}

/*
下記メソッドはメソッド名の先頭が小文字のため内部でしか使えない
*/
func sampleSwitch() {

	fmt.Println("sampleSwitch を利用している")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var text = scanner.Text()

	//　goの仕様上、1つのみしか動かない。
	switch text {
	case "000":
		fmt.Println("数字文字列だな")
	case "001":
		fmt.Println("しかも1以上")
	case "011":
		fmt.Println("さらに11以上")
	default:
		fmt.Println(text, "何を入力したんだ貴様？")
	}

}

/*
下記メソッドはメソッド名の先頭が小文字のため内部でしか使えない
*/
func variable02() {

	fmt.Println("variable02 を利用している")

	text := "variable02"
	fmt.Println(text)

	var i = 0
	fmt.Println(i)

	var j int = 0
	fmt.Println(j)

	// 最後に実行
	defer fmt.Println("hello world")
	// 配列
	// [3]int = {1,2,3}

}
