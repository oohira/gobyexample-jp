// [タイマー](timers) は、未来に一度だけ何かしたいときに使いますが、
// _ティッカー (tickers)_ は定期的に何かしたいときに使います。
// ここでは、停止するまで定期的に動作するティッカーの例を見ます。

package main

import "time"
import "fmt"

func main() {

    // ティッカーは、タイマーと同様の仕組み、すなわち
    // 値を送信するチャネルを使います。
    // ここでは、500ミリ秒ごとに受信される値を繰り返し
    // 処理するために、`range` ビルトイン関数を使っています。
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // ティッカーは、タイマーと同じように停止できます。
    // ティッカーが停止されると、そのチャネルからはもう
    // 値を受信しなくなります。
    // この例では、1600ミリ秒後に停止します。
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
