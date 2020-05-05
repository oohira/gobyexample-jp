// ときには、Go プログラムで
// [Unix シグナル](http://en.wikipedia.org/wiki/Unix_signal)
// を扱いたいこともあります。例えば、`SIGTERM`
// を受け取ったらサーバーを正常に終了させたり、`SIGINT`
// を受け取ったらコマンドラインツールの入力処理を止めたり、
// といった具合です。
// 以下は、Go でチャネルを使ってシグナルを扱う方法です。

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go のシグナル通知は、チャネルに `os.Signal`
	// 値を送信することで行います。
	// これらの通知を受信するためのチャネル
	// (と、プログラムが終了できることを通知するためのチャネル)
	// を作ります。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` は、指定されたシグナル通知を受信するために、
	// 与えられたチャネルを登録します。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// このゴルーチンは、シグナルを同期的に受信します。
	// シグナルを受信したら、それを表示して、
	// プログラムに終了できることを通知します。
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// プログラムはシグナルを受信するまで
	// (前述の `done` に値を送信するゴルーチンで知らされる)
	// 待機した後、終了します。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
