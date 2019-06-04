// _関数 (Functions)_ は、Go の中核です。
// いくつかの例で関数について学んでいきます。

package main

import "fmt"

// これは 2 つの `int` を受け取って、
// その合計を `int` で返す関数です。
func plus(a int, b int) int {

	// Go では、明示的な return が必要です。つまり、
	// 最後に評価した式を自動的には return してくれません。
	return a + b
}

// 同じ型の引数が連続する場合は、途中の引数の型名を省略し、
// 最後の引数にのみ型を宣言することもできます。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// 期待通り、 `name(args)` で関数を呼び出せます。
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
