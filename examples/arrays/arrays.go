// Go では、 _配列 (array)_ は特定の長さをもつ
// 番号付けられた要素の列です。

package main

import "fmt"

func main() {

    // ここでは、ちょうど 5 つの `int` をもつ配列
    // `a` を作成しています。要素の型と長さはいずれも
    // 配列の型の一部です。デフォルトでは、配列は
    // ゼロ値で初期化されるので、 `int` の場合は
    // `0` になります。
    var a [5]int
    fmt.Println("emp:", a)

    // `array[index] = value` で指定位置に値を設定し、
    // `array[index]` で値を取得することができます。
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // `len` ビルトイン関数は配列の長さを返します。
    fmt.Println("len:", len(a))

    // 配列を 1 行で宣言かつ初期化するには、
    // 次の記法を使います。
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 配列型は一次元ですが、多次元のデータ構造を
    // 表す型を構成することもできます。
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
