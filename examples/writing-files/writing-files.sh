# ファイル書き込みのコードを実行します。
$ go run writing-files.go
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# そして、書き出されたファイルの中身を確認します。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 次は、`stdin` や `stdout` ストリームに見られる、
# いくつかのファイル I/O の考え方を適用する例を見ていきます。
