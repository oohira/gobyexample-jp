// [環境変数](https://en.wikipedia.org/wiki/Environment_variable)
// は、[設定情報を Unix プログラムに渡す](https://www.12factor.net/config)
// ための、一般的な仕組みです。
// 環境変数を設定、取得、一覧する方法を見ていきましょう。

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// キーと値のペアを設定するには、`os.Setenv` を使います。
	// キーに対する値を取得するには、`os.Getenv` を使います。
	// 環境にキーが存在しなければ、空文字列が返されます。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 環境に定義されたすべてのキーと値のペアを列挙するには、
	// `os.Environ` を使います。これは、`KEY=value`
	// という形式の文字列のスライスを返します。
	// キーと値をそれぞれ取得するために、`strings.SplitN`
	// を使えます。次の例は、すべてのキーを出力します。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
