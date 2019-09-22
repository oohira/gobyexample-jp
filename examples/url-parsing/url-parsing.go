// URL は、[リソースを特定する統一された方法](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)
// を提供します。以下は、Go で URL をパースする方法です。

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// このサンプル URL をパースするとします。
	// これは、スキーム、認証情報、ホスト、ポート番号、パス、
	// クエリパラメーターとクエリフラグメントを含みます。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL をパースし、エラーがないことを確認します。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// スキームへのアクセスは簡単です。
	fmt.Println(u.Scheme)

	// `User` は認証情報全体を保持します。個々の値を取得するには、
	// このオブジェクトの `Username` と `Password` を呼び出します。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` は、指定されていればホスト名とポート番号の両方を保持します。
	// 個々の値を抽出するには、 `SplitHostPort` を使ってください。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 次の例は、`path` と `#` の後のフラグメントを抽出しています。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// `k=v` 形式の文字列でクエリパラメーターを取得するには、
	// `RawQuery` を使います。
	// クエリパラメーターをマップとしてパースすることもできます。
	// パースされたクエリパラメーターのマップは、
	// 文字列 → 文字列のスライスへのマップです。そのため、
	// 最初の値だけ取得したければインデックス `[0]` を指定します。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
