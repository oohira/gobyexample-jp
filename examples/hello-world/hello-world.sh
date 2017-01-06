# プログラムを実行するには、コードを `hello-world.go`
# ファイルに保存し、 `go run` を実行します。
$ go run hello-world.go
hello world

# 作成したプログラムをバイナリ形式にビルドしたいことも
# あるでしょう。 `go build` を使えば可能です。
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# そうすれば、ビルドしたバイナリを直接実行できます。
$ ./hello-world
hello world

# これで基本的な Go プログラムをビルド・実行できるように
# なりました。さらに詳しく学んでいきましょう。
