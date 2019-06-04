// Go は、組み込みで [正規表現](http://en.wikipedia.org/wiki/Regular_expression)
// をサポートしています。以下は、Go で正規表現を使う一般的な例です。

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// これは、パターンが文字列にマッチするか判定します。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上の例では文字列パターンを直接使いましたが、
	// 他の正規表現タスクでも使うためには、`Compile`
	// して最適化された `Regexp` 構造体を作る必要があります。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// この構造体では、多くのメソッドが利用できます。
	// 最初に見たような、マッチの判定は次の通りです。
	fmt.Println(r.MatchString("peach"))

	// 正規表現にマッチする部分の検索は次の通りです。
	fmt.Println(r.FindString("peach punch"))

	// これも最初にマッチする部分を検索しますが、文字列の代わりに、
	// 開始および終了インデックスを返します。
	fmt.Println(r.FindStringIndex("peach punch"))

	// `Submatch` は、全体のパターンにマッチした部分と、
	// その中でサブマッチした部分の情報を含みます。
	// 例えば次の例では、`p([a-z]+)ch` と `([a-z]+)`
	// にマッチした部分を返します。
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 同様に、次の例はサブマッチした部分のインデックスも返します。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// `All` がつくバリエーションは、最初のマッチだけでなく、
	// 入力の全てのマッチ部分を対象とします。
	// 例えば、正規表現にマッチする全ての文字列を検索する例は、
	// 次の通りです。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 同様に、前述した他の関数にも `All` がつく関数があります。
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// これらの関数の第 2 引数に非負の整数を与えると、
	// マッチする数を制限することができます。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// これまでの例では、文字列型の引数をもち、`MatchString`
	// といった名前の関数を使いました。名前から `String`
	// を除いた関数に、`[]byte` 型の引数を渡すこともできます。
	fmt.Println(r.Match([]byte("peach")))

	// 正規表現の定数を作る場合は、`Compile` の
	// バリエーションとして `MustCompile` を使えます。
	// ただの `Compile` は、2 つの戻り値を返すため、
	// 定数には使えません。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// `regexp` パッケージは、文字列の一部を他の文字列で
	// 置換するためにも使えます。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` がつくバリエーションを使うと、マッチした文字列を
	// 指定した関数で変換できます。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
