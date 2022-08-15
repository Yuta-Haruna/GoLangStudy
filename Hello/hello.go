package hello

import (
	"bufio"
	"fmt"
	"os"
)

// 入力されたテキストに　追加ワードを入れる
// デフォルト文字「 aaaa 」
func Input(msg string) string {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(msg + " aaaa ")
	scanner.Scan()
	return scanner.Text()
}
