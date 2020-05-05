// _チャネル (Channels)_ は、並行実行されるゴルーチン間を
// 接続するパイプです。あるゴルーチンからチャネルを通して
// 値を送信し、別のゴルーチンでその値を受信できます。

package main

import "fmt"

func main() {

	// `make(chan 値の型)` で新しいチャネルを作成します。
	// チャネルは送受信する値で型付けされます。
	messages := make(chan string)

	// `channel <-` 構文で、チャネルへ値を _送信_ します。
	// ここでは、新しいゴルーチンから `messages` チャネルへ
	// `"ping"` という値を送信しています。
	go func() { messages <- "ping" }()

	// `<-channel` 構文で、チャネルから値を _受信_ します。
	// ここでは、上で送信した `"ping"` メッセージを受信して
	// 画面に表示しています。
	msg := <-messages
	fmt.Println(msg)
}
