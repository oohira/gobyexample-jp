// [_コマンドライン引数_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// は、プログラムの実行をパラメーター化する一般的な方法です。
// 例えば、`go run hello.go` は、`run` と
// `hello.go` を `go` プログラムの引数として使います。

package main

import "os"
import "fmt"

func main() {

	// `os.Args` は、生のコマンドライン引数へのアクセスを提供します。
	// このスライスの 1 つ目の値は、プログラム自身のパスで、
	// `os.Args[1:]` がプログラムへの引数を保持することに
	// 注意してください。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 個々の引数はインデックスを使って通常通り取得できます。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
