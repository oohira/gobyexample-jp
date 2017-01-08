// Go は、 _複数の戻り値 (multiple return values)_ をサポートします。
// この機能は、Go で慣用的によく使われます。例えば、
// 関数から結果とエラーの両方の値を返す場合などです。

package main

import "fmt"

// この関数シグネチャの `(int, int)` は、この関数が
// 2 つの `int` を返すことを示しています。
func vals() (int, int) {
    return 3, 7
}

func main() {

    // ここでは、 _多重代入 (multiple assignment)_
    // で関数呼び出しからの 2 つの戻り値を使っています。
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // もし戻り値の一部しか欲しくないのであれば、
    // ブランク識別子 `_` を使ってください。
    _, c := vals()
    fmt.Println(c)
}
