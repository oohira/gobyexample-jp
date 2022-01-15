// Go は、`recover` ビルトイン関数を使うことで、
// panic から _recover_ できるようになっています。
// `recover` は、`panic` がプログラムを異常終了するのを止め、
// 実行を継続させます。

// これが役に立つ例として、サーバーがあります。
// サーバーは、クライアント接続の1つが致命的なエラーを出したとしても
// クラッシュさせたくありません。代わりに、サーバーはその接続を閉じ、
// ほかのクライアントの処理を継続したいでしょう。実際、Go の
// `net/http` は、HTTP サーバーに同じことをしています。

package main

import "fmt"

// この関数は panic を起こします。
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` は、defer 関数の中で実行する必要があります。
	// 外側の関数が panic を起こした場合、defer が起動し、
	// その中の `recover` 呼び出しが panic をキャッチします。
	defer func() {
		if r := recover(); r != nil {
			// `recover` の戻り値は、`panic`
			// の呼び出しに渡された error です。
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// `mayPanic` が panic を起こすので、このコードは実行されません。
	// `main` の実行は panic の時点でストップし、defer クロージャの
	// 中で再開されます。
	fmt.Println("After mayPanic()")
}
