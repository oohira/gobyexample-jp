// Go は、よりシームレスな型の _コンポジション (composition)_
// を表現できるよう、構造体やインターフェースの _埋め込み (embedding)_
// をサポートしています。

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` は、`base` を _埋め込み_ ます。
// 埋め込みは、名前のないフィールドのように見えます。
type container struct {
	base
	str string
}

func main() {

	// リテラルを使って構造体を作成する場合には、
	// 埋め込まれた型を明示的に初期化する必要があります。
	// このとき、埋め込まれた型がフィールド名として機能します。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// `co` に対して、`co.num` のように base のフィールドに
	// 直接アクセスできます。
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// 代わりに、埋め込まれた型の名前を使って、
	// フルパスで指定することもできます。
	fmt.Println("also num:", co.base.num)

	// `container` は `base` を埋め込んでいるので、`base`
	// のメソッドも `container` のメソッドになります。ここでは、
	// `co` に対して `base` から埋め込まれたメソッドを直接呼び出しています。
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// メソッドをもつ構造体の埋め込みは、ほかの構造体にインターフェース
	// 実装を与えるために利用されることがあります。
	// この例では、`container` が `base` を埋め込んでいるため、
	// `describer` インターフェースを実装していることになります。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
