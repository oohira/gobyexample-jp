// ゴルーチン間の実行を同期するために、チャネルを使うことができます。
// ここでの例は、ゴルーチンの完了を待つために、ブロッキング受信を使います。
// 複数のゴルーチンの完了を待つ場合は、
// [WaitGroup](waitgroups) を使うとよいでしょう。

package main

import "fmt"
import "time"

// ゴルーチンで実行する関数は次の通りです。
// この関数が完了したことを別のゴルーチンに通知するために、
// `done` チャネルが使われます。
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // 完了したことを通知するために値を送信します。
    done <- true
}

func main() {

    // 通知用のチャネルを渡して、`worker` ゴルーチンを開始します。
    done := make(chan bool, 1)
    go worker(done)

    // チャネルへの完了通知を受信するまでブロックします。
    <-done
}
