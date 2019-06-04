// Go の標準ライブラリは、`net/http` パッケージに
// HTTP クライアントとサーバーに対する素晴らしいサポートを含んでいます。
// この例では、簡単な HTTP リクエストを発行するためにそれを使います。
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// サーバーに HTTP GET リクエストを発行します。
	// `http.Get` は、`http.Client` オブジェクトを作り、
	// `Get` メソッドを呼ぶ便利なショートカットです。
	// それは、有用なデフォルト設定をもつ `http.DefaultClient`
	// オブジェクトを使います。
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// HTTP レスポンスステータスを出力します。
	fmt.Println("Response status:", resp.Status)

	// レスポンスボディの最初の 5 行を出力します。
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
