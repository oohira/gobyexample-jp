// Go の標準ライブラリは、Go プログラムからログを出力するための
// 簡単なツールを提供します。自由形式のログには
// [log](https://pkg.go.dev/log) パッケージを、構造化ログには
// [log/slog](https://pkg.go.dev/log/slog) パッケージを使えます。
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// `log` パッケージの `Println` などの関数を呼び出すと、
	// すでに設定済みの _標準_ ロガーを使って、`os.Stderr`
	// に適切なログを出力できます。`Fatal*` や `Panic*`
	// などのメソッドは、ログ出力の後にプログラムを終了します。
	log.Println("standard logger")

	// ロガーは、_フラグ (flag)_ を使用して出力形式を設定できます。
	// デフォルトでは、標準ロガーには `log.Ldate` と `log.Ltime`
	// フラグが設定されていて、これは `log.LstdFlags` と同等です。
	// 例えば、ミリ秒精度で時間を出力するようにフラグを変更できます。
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// 呼び出し元のファイル名と行番号を出力する設定もできます。
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// カスタムロガーを作成して利用することも有用です。
	// 新しいロガーを作成するときに、他のロガーの出力と区別しやすいよう
	// _接頭辞 (prefix)_ を設定できます。
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// 既存のロガー（標準ロガーを含む）の接頭辞は、`SetPrefix`
	// メソッドで設定できます。
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// ロガーは、カスタムの出力先を設定できます。
	// 任意の `io.Writer` が使えます。
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// この呼び出しは、ログ出力を `buf` に書き込みます。
	buflog.Println("hello")

	// これにより、標準出力に実際の出力が表示されます。
	fmt.Print("from buflog:", buf.String())

	// `slog` パッケージは、_構造化された_ ログ出力を提供します。
	// 例えば、JSON 形式のログ出力が簡単に実現できます。
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// メッセージに加えて、`slog` の出力には任意のキーと値のペアを
	// 含めることができます。
	myslog.Info("hello again", "key", "val", "age", 25)
}
