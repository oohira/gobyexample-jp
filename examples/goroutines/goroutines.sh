# このプログラムを実行すると、最初に同期呼び出しの出力、
# その次に 2 つのゴルーチンの混ざった出力を確認できます。
# これは、ゴルーチンが Go ランタイムによって
# 並行実行されていることを示しています。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# 次回は、Go の並行プログラムを補完するチャネルを見ていきます。
