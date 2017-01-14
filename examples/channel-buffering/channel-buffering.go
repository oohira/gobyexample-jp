// デフォルトでは、チャネルは _バッファリングされない (unbuffered)_
// ので、受信 (`<- chan`) の準備ができていないと送信 (`chan <-`)
// できません。
// しかし、_バッファリングされたチャネル (Buffered channels)_
// は、対応する受信側がいなくても決められた量までなら
// 値を送信することができます。

package main

import "fmt"

func main() {

    // この例では、`string` を 2 つまでバッファリングするチャネルを
    // `make` しています。
    messages := make(chan string, 2)

    // このチャネルはバッファリングされるので、対応する受信側が
    // いなくても値をチャネルに送信することができます。
    messages <- "buffered"
    messages <- "channel"

    // そして、あとで通常通り 2 つの値を受信することができます。
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
