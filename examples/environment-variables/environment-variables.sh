# プログラムを実行すると、プログラム中で設定した `FOO`
# の値は取得できますが、`BAR` の値は空になることが分かります。
$ go run environment-variables.go
FOO: 1
BAR:

# 環境に定義されたキーの一覧は、あなたのマシンに依存します。
TERM_PROGRAM
PATH
SHELL
...

# もし環境変数 `BAR` を先に設定すると、
# 実行したプログラムはその値を取得できます。
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
