// 複数のゴルーチンの完了を待つために、*WaitGroup* を使うことができます。

package main

import (
	"fmt"
	"sync"
	"time"
)

// これは、ゴルーチンで実行する関数です。
// WaitGroup はポインタで関数に渡さなければならない点に注意してください。
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// 時間のかかるタスクをシミュレートするためにスリープします。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// このワーカーが完了したことを WaitGroup に通知します。
	wg.Done()
}

func main() {

	// この WaitGroup は、ここで起動したすべてのゴルーチンが完了するのを
	// 待つために使われます。
	var wg sync.WaitGroup

	// いくつかのゴルーチンを起動し、そのたびに WaitGroup のカウンターを
	// インクリメントします。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// WaitGroup のカウンターが 0 に戻るまで（つまり、すべてのワーカーが
	// 完了したと通知してくるまで）ブロックします。
	wg.Wait()
}
