// Go は、文字や文字列、真偽値、数値の _定数 (constants)_
// をサポートします。

package main

import "fmt"
import "math"

// `const` は定数を宣言します。
const s string = "constant"

func main() {
    fmt.Println(s)

    // `const` ステートメントは、 `var` ステートメントが
    // 書ける場所ならどこにでも書けます。
    const n = 500000000

    // 定数式は、任意の精度の算術演算ができます。
    const d = 3e20 / n
    fmt.Println(d)

    // 数値の定数は、明示的に型変換されるなどして、
    // 初めて具体的な型が決まります。
    fmt.Println(int64(d))

    // 数値は、変数代入や関数呼び出しなど、
    // その状況で必要とされる型をとります。この例では、
    // `math.Sin` は `float64` を受け取ります。
    fmt.Println(math.Sin(n))
}
