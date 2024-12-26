// カスタム型に `Error()` メソッドを実装することで、独自の型を
// `error` として使用できます。
// 以下は、引数エラーを明示的に表すカスタム型を使用した例です。

package main

import (
	"errors"
	"fmt"
)

// カスタムエラー型には通常 "Error" という接尾辞を付けます。
type argError struct {
	arg     int
	message string
}

// `argError` に `Error` メソッドを追加することで、
// `error` インターフェースを実装できます。
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// カスタムエラーを返します。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.As` は、`errors.Is` のより高度なバージョンです。
	// 指定されたエラー（またはその連鎖内の任意のエラー）が特定の
	// エラー型に一致するかを確認し、一致する場合はその型の値に変換して
	// `true` を返します。一致しない場合は `false` を返します。
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
