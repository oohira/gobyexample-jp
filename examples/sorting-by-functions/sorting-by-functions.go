// ときには、コレクションを自然順以外でソートしたいこともあります。
// 例えば、文字列をアルファベット順ではなく、
// 長さ順でソートしたいとしましょう。
// Go のカスタムソートの例は次の通りです。

package main

import "sort"
import "fmt"

// Go でカスタム関数を使ってソートするためには、
// 対応する型が必要です。ここでは、`ByLength`
// 型を作りました。これは、
// 組み込みの `[]string` 型のただのエイリアスです。
type ByLength []string

// `sort` パッケージの `Sort` 関数を使えるように、
// `sort.Interface` すなわち `Len`, `Less`, `Swap`
// 関数を実装します。
// `Len` と `Swap` はどの型でもだいたい同じになり、
// `Less` が実際のカスタムソートのロジックをもちます。
// この例では、文字列の長さの昇順でソートしたいので、
// `len(s[i])` と `len(s[j])` を使っています。
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// 元の `fruits` スライスを `ByLength` にキャストし、
// `sort.Sort` 関数を使うことでカスタムソートを実現できます。
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
