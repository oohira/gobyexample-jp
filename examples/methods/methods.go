// Go は構造体型に定義する _メソッド (methods)_ をサポートします。

package main

import "fmt"

type rect struct {
    width, height int
}

// この `area` メソッドは、 `*rect` を _レシーバ型 (receiver type)_
// としてもちます。
func (r *rect) area() int {
    return r.width * r.height
}

// メソッドは、ポインタまたは値のどちらのレシーバ型に対しても
// 定義することができます。これは値レシーバの例です。
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // 構造体に定義された 2 つのメソッドを呼び出します。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Go は、メソッド呼び出しに対する値・ポインタ間の変換を
    // 自動的に扱います。
    // メソッドを呼び出す時の構造体の不要なコピーを避けたり、
    // 元の構造体自体を変更したりするためには、
    // ポインタレシーバ型を使ってください。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
