// Go では、エラーを明示的な、別の戻り値として扱うのが特徴です。
// これは、Java や Ruby のような言語で使われる例外や、
// C で時々使われる結果/エラーを多重定義した単一の値とは
// 対照的です。
// Go のアプローチは、どの関数がエラーを返したかを調べやすくし、
// エラー以外のほかのタスクに使うのと同じ言語機能でエラーも
// 扱えるようにします。
//
// 詳細については、[errors パッケージ](https://pkg.go.dev/errors)
// のドキュメントや、[このブログ記事](https://go.dev/blog/go1.13-errors)
// を参照してください。

package main

import (
	"errors"
	"fmt"
)

// 慣例的に、エラーは戻り値の最後にし、
// 組み込みインターフェースである `error` 型をもちます。
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` は、与えられたエラーメッセージをもつ
		// 基本的な `error` 値を作ります。
		return -1, errors.New("can't work with 42")
	}

	// エラーに対する `nil` 値は、エラーがなかったことを示します。
	return arg + 3, nil
}

// センチネルエラー (sentinel error) とは、特定のエラー状態を
// 表すために使用される、あらかじめ宣言された変数のことです。
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// エラーをより高レベルなエラーでラップして、文脈を追加する
		// ことができます。最も簡単な方法は、`fmt.Errorf` の `%w`
		// を使用することです。ラップされたエラーは、論理的な連鎖
		// (A が B をラップし、B が C をラップする、など) を形成し、
		// `errors.Is` や `errors.As` のような関数で問い合わせできます。
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// `if` 文の中でインラインでエラーチェックするのは一般的です。
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` は、指定されたエラー（またはその連鎖内の任意のエラー）
			// が特定のエラー値と一致するかを確認します。これは、ラップされたエラーや
			// ネストされたエラーの場合に特に有用で、エラーの連鎖内で特定のエラータイプや
			// センチネルエラーを識別することができます。
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
