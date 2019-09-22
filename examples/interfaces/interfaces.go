// _インターフェース (Interfaces)_
// は、メソッドシグネチャの集まりに名前を付けたものです。

package main

import (
	"fmt"
	"math"
)

// これは図形に対する基本的なインターフェースです。
type geometry interface {
	area() float64
	perim() float64
}

// 例として、 `rect` 型と `circle` 型で
// このインターフェースを実装してみましょう。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Go でインターフェースを実装するには、
// そのインターフェースのすべてのメソッドを実装すればよいだけです。
// ここでは、`rect` に対して `geometry` を実装しています。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` に対する実装です。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 変数がインターフェース型をもつなら、
// そのインターフェースにあるメソッドを呼ぶことができます。
// これを利用すると、任意の `geometry` に対して動作する
// `measure` 関数は次のようになります。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// `circle` および `rect` 構造体型はいずれも
	// `geometry` インターフェースを実装するので、
	// これらの構造体のインスタンスを `measure`
	// 関数の引数として使うことができます。
	measure(r)
	measure(c)
}
