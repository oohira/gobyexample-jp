// _列挙型_ (enum) は、[直和型](https://en.wikipedia.org/wiki/Algebraic_data_type)
// の特殊なケースです。列挙型とは、決まった値のどれかをもつ型で、
// それぞれの値には固有の名前がつけられています。Go には言語機能としての
// 列挙型はありませんが、既存の言語の慣用表現を使えば簡単に実装できます。

package main

import "fmt"

// 列挙型 `ServerState` の実体は `int` 型です。
type ServerState int

// `ServerState` の取りうる値を定数として定義します。
// 特殊なキーワード [iota](https://go.dev/ref/spec#Iota)
// は、連続する定数値を自動生成します。
// この場合、0, 1, 2, ... となります。
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// インターフェースを実装することで、`ServerState` の値を
// 文字列で出力したり、文字列に変換したりできるようになります。
//
// 値が多い場合、この実装は面倒になることがあります。その場合は、
// [stringer ツール](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// を `go:generate` と組み合わせて自動化できます。
// [こちらの記事](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// に詳細な説明があります。
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// `int` 型の値を `transition` に渡そうとすると、コンパイラは
	// 型の不一致を指摘します。これにより、列挙型に対する型安全性が
	// コンパイル時にある程度は担保されます。

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition は、サーバーの状態遷移をエミュレートします。
// この関数は、現在の状態を受け取り、新しい状態を返します。
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// 次の状態を決定するための条件をここでチェックする想定です。
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
