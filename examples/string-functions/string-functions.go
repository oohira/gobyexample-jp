// 標準ライブラリの `strings` パッケージは、
// 文字列関連の役に立つ関数を数多く提供しています。
// 使い方を把握するための例をいくつか挙げます。

package main

import (
	"fmt"
	s "strings"
)

// あとで何回も使うので、`fmt.Println`
// に短い名前のエイリアスを定義します。
var p = fmt.Println

func main() {

	// `strings` パッケージで使える関数の例は次の通りです。
	// これらはパッケージで定義される関数であって、
	// 文字列オブジェクトそのもののメソッドではないため、
	// 関数への第 1 引数として対象の文字列を渡す必要があります。
	// さらに多くの関数を [`strings`](https://pkg.go.dev/strings)
	// パッケージのドキュメントで確認できます。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
