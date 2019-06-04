// チャネル上の基本的な送受信はブロックします。
// しかし、`default` 句をもった `select` を使えば、
// _ブロックしない (non-blocking)_ 送信、受信、
// そして多重 `select` さえ実装することができます。

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// これは、ブロックしない受信の例です。
	// もし `messages` チャネルで値が準備できていれば、
	// `select` は `<-messages` のケースを実行します。
	// そうでなければ、すぐに `default` ケースを実行します。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// ブロックしない送信も同様に動作します。`messages`
	// チャネルには、バッファもなければ受ける側もいないため、
	// `msg` を送信することはできません。
	// そのため、`default` ケースが実行されます。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// ブロックしない多重 `select` を実現するために、
	// `default` 句の前に複数の `case` を使うことができます。
	// 次の例では、`messages` と `signals` の両方に対して
	// ブロックしない受信を試みます。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
