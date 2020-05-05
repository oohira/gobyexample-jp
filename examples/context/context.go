// 前回は、簡単な [HTTP サーバー](http-servers)
// をセットアップする例を見ました。HTTP サーバーは、
// キャンセルを制御するための `context.Context`
// をデモするのに便利です。`Context` は、API
// 境界やゴルーチンをまたいで、期限やキャンセルシグナル、
// その他のリクエストスコープの値を受け渡します。
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `context.Context` は `net/http` によってリクエスト毎に
	// 作られ、`Context()` メソッドで取得できます。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// クライアントに応答を返す前に数秒間待ちます。これは、
	// 実際にサーバーがする何らかのタスクをシミュレートしています。
	// タスクを実行する間、コンテキストの `Done()` チャネルを監視し、
	// シグナルがあれば可能な限り早くタスクをキャンセルして制御を戻します。
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// コンテキストの `Err()` メソッドは、`Done()`
		// チャネルがクローズされた理由を説明するエラーを返します。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// 前回と同様に、"/hello" ルートに対するハンドラを登録し、
	// サーバーを開始します。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
