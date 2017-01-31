# このプログラムを実行すると、panic を引き起こし、
# エラーメッセージとゴルーチンのトレースが表示され、
# 非ゼロのステータスで終了します。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 多くのエラー処理に例外を使う言語とは異なり、Go では可能な限り
# エラーを示す戻り値を使うのが慣習であることに注意してください。
