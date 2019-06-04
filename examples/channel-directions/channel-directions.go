// チャネルを関数の引数として使う場合、そのチャネルが
// 送信または受信だけを意図しているか指定することができます。
// これは、プログラムをより型安全にします。

package main

import "fmt"

// この `ping` 関数は、チャネルを送信専用で受け取ります。
// このチャネルで受信しようとすると、コンパイルエラーになります。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// この `pong` 関数は、1 つ目のチャネルを受信専用で (`pings`)、
// 2 つ目のチャネルを送信専用で (`pongs`) 受け取ります。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
