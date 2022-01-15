// 複数のゴルーチンの完了を待つために、*WaitGroup* を使うことができます。

package main

import (
	"fmt"
	"sync"
	"time"
)

// これは、すべてのゴルーチンで実行する関数です。
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 時間のかかるタスクをシミュレートするためにスリープします。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// この WaitGroup は、ここで起動したすべてのゴルーチンが完了するのを
	// 待つために使われます。WaitGroup を関数に引数として渡す場合は、
	// *ポインタで* 渡す必要がある点に注意してください。
	var wg sync.WaitGroup

	// いくつかのゴルーチンを起動し、そのたびに WaitGroup のカウンターを
	// インクリメントします。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// ゴルーチンのクロージャで同じ `i` の値が再利用されてしまうのを防ぎます。
		// 詳細は、[FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// を参照してください。
		i := i

		// クロージャでワーカー呼び出しをラップし、ワーカーが完了したことを
		// WaitGroup に伝えます。こうすると、ワーカー自身は並行処理のことを
		// 意識する必要がなくなります。
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// WaitGroup のカウンターが 0 に戻るまで（つまり、すべてのワーカーが
	// 完了したと通知してくるまで）ブロックします。
	wg.Wait()

	// このアプローチは、ワーカーからエラーを簡単に伝える方法がない点に
	// 注意してください。より発展的なユースケースのためには、
	// [errgroup パッケージ](https://pkg.go.dev/golang.org/x/sync/errgroup)
	// の利用を検討してください。
}
