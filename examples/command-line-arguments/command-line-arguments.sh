# コマンドライン引数を試すには、`go build`
# で先にバイナリをビルドするのがよいです。
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]
[a b c d]
c

# 次は、フラグを使ったさらに進んだコマンドライン処理を
# 見ていきましょう。
