// `net/http` パッケージを使えば、基本的な HTTP サーバーを
// 書くのは簡単です。
package main

import (
	"fmt"
	"net/http"
)

// `net/http` サーバーの基本コンセプトは、*ハンドラ* です。
// ハンドラは、`http.Handler` インターフェースを実装した
// オブジェクトです。ハンドラを書く一般的な方法は、
// 適切なシグネチャをもつ関数に対して、`http.HandlerFunc`
// アダプタを使うことです。
func hello(w http.ResponseWriter, req *http.Request) {

	// ハンドラとして動作する関数は、引数として
	// `http.ResponseWriter` と `http.Request` を受け取ります。
	// `ResponseWriter` は、HTTP レスポンスを書き込むために
	// 使われます。ここでは単に "hello\n" とレスポンスします。
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// このハンドラはもう少しだけ高度です。
	// すべてのHTTPリクエストヘッダを読み込み、
	// それらをレスポンスボディにエコーバックします。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// 便利関数 `http.HandleFunc` を使ってハンドラをサーバーの
	// ルーティングに登録します。これは、`net/http` パッケージ内の
	// *デフォルトルーター* を設定し、関数を引数に取ります。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最後に、ポート番号とハンドラを指定して `ListenAndServe`
	// を呼び出します。`nil` は、上で設定したデフォルトルーターを
	// 使用することを意味します。
	http.ListenAndServe(":8090", nil)
}
