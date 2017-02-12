// [Unix エポック](http://en.wikipedia.org/wiki/Unix_time)
// からの秒数やミリ秒数、ナノ秒数を取得することは、
// プログラムの一般的な要件です。Go での方法は、次の通りです。

package main

import "fmt"
import "time"

func main() {

    // Unix エポックからの経過時間を秒やナノ秒で取得するには、
    // `time.Now` を `Unix` や `UnixNano` と使います。
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // なお、`UnixMillis` はない点に注意してください。
    // そのため、エポックからのミリ秒を取得するためには、
    // ナノ秒を手動で割って計算する必要があります。
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // エポックからの秒数やナノ秒数を表す数値を、
    // 対応する `time` に変換することもできます。
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
