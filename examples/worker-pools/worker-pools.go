// ここでは、ゴルーチンとチャネルを使って、
// _ワーカープール (worker pool)_ を実装する例を見ていきます。

package main

import (
	"fmt"
	"time"
)

// これは、複数インスタンスが並行実行されるワーカーです。
// これらのワーカーは、`jobs` チャネルからタスクを受信し、
// 結果を `results` チャネルに送信します。
// 実行コストの高いジョブをシミュレートするため、
// 各タスクは 1 秒スリープします。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ワーカープールを使うためには、タスクをワーカーに送信し、
	// それらの結果を集める必要があります。
	// そのために、2 つのチャネルを作成します。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ここで、3 つのワーカーを開始しますが、
	// 最初はまだジョブがないためブロックします。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 次に、5 つのジョブを送信し、それがすべてであることを
	// 伝えるためにチャネルを `close` します。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 最後に、すべてのタスクの結果を集めます。
	// これにより、ワーカーのゴルーチンが完了することも保証します。
	// 複数のゴルーチンを待つ別の方法としては、
	// [WaitGroup](waitgroups) が使えます。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
