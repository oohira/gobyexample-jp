# 作成したフィルタプログラムを試すために、
# まず小文字の行を含むファイルを作成します。
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# それから大文字の行を得るためにフィルタを実行します。
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
