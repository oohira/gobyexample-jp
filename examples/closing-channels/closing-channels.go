// チャネルを _クローズする (Closing)_ ことは、
// もう値を送信しないことを意味します。これは、
// チャネルの受け手に完了を伝えるのに便利です。

package main

import "fmt"

// この例では、`main()` ゴルーチンからワーカーのゴルーチンへ
// タスクの完了を伝えるために、`jobs` チャネルを使います。
// ワーカータスクがなくなれば、`jobs` チャネルを `close` します。
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // ワーカーのゴルーチンは次の通りです。`j, more := <-jobs`
    // で `jobs` チャネルから繰り返し受信します。
    // この 2 値の形式の受信では、`jobs` が `close` され、
    // チャネルのすべての値がすでに受信されていれば、
    // `more` の値が `false` になります。
    // ここでは、すべてのタスクが完了したときに、
    // `done` チャネルへ通知するために使っています。
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // これは、`jobs` チャネルを通して 3 つのジョブを
    // ワーカーへ送信し、その後チャネルをクローズします。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // すでに学んだチャネルの [同期](channel-synchronization)
    // 手法を使って、ワーカーの完了を待ちます。
    <-done
}
