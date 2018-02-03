// プログラムでデータのコレクションに対して何か操作をしたい
// ことはよくあります。
// 与えられた述語を満たす全てのアイテムを選択したり、
// カスタム関数を使って全てのアイテムを新しいコレクションに
// マッピングしたりといった具合です。

// いくつかの言語では [ジェネリック (generic)](http://en.wikipedia.org/wiki/Generic_programming)
// なデータ構造とアルゴリズムを使うのが慣習です。
// しかし、Go はジェネリクスをサポートしていません。
// Go では、プログラムやデータ型が必要とする場合だけ
// コレクション関数を提供するのが一般的です。

// 以下は、`string` のスライスに対するコレクション関数の例です。
// あなた自身の関数を作るために、これらの例を使えるでしょう。
// ときには、ヘルパー関数を作って呼び出す代わりに、
// コレクションを操作するコードを単にインラインで定義した方が、
// 分かりやすいこともあるので覚えておきましょう。

package main

import "strings"
import "fmt"

// Index は、目的の文字列 `t` の最初のインデックスを返します。
// 見つからなかった場合は、`-1` になります。
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Include は、文字列 `t` がスライスに含まれる場合に
// `true` を返します。
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Any は、スライスの文字列が 1 つでも述語 `f` を満たす場合に
// `true` を返します。
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// All は、スライスの全ての文字列が述語 `f` を満たす場合に
// `true` を返します。
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Filter は、述語 `f` を満たす全ての文字列を含む、新しいスライスを返します。
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Map は、元のスライスの各文字列に関数 `f` を適用した結果を含む、
// 新しいスライスを返します。
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    // 色々なコレクション関数を試します。
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    // 上の例は全て無名関数を使っていますが、
    // 正しい型の名前付き関数を使うこともできます。
    fmt.Println(Map(strs, strings.ToUpper))

}
