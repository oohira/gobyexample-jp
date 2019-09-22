// 文字列から数値をパースすることは基本的な、
// しかし多くのプログラムで共通のタスクです。以下は、Go での方法です。

package main

// 組み込みパッケージの `strconv` が数値のパース機能を提供します。
import (
	"fmt"
	"strconv"
)

func main() {

	// `ParseFloat` で、`64`
	// は何ビットの精度でパースするかを指定しています。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt` に対しては、`0`
	// は文字列から基数を推測することを意味します。
	// `64` は結果が 64 ビット長に合うことを要求します。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` は、16 進文字列も認識します。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` も使うことができます。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` は、10 進数の `int` をパースする便利関数です。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// パース関数は不正な入力に対してはエラーを返します。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
