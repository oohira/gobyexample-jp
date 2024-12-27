// [Unix エポック](https://en.wikipedia.org/wiki/Unix_time)
// からの秒数やミリ秒数、ナノ秒数を取得することは、
// プログラムの一般的な要件です。Go での方法は、次の通りです。

package main

import (
	"fmt"
	"time"
)

func main() {

	// Unix エポックからの経過時間を秒、ミリ秒、ナノ秒で取得するには、
	// `time.Now` を `Unix` や `UnixMilli`、`UnixNano` と使います。
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// エポックからの秒数やナノ秒数を表す数値を、
	// 対応する `time` に変換できます。
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
