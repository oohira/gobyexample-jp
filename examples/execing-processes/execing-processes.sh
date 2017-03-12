# プログラムを実行すると、`ls` に置き換わります。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Go は、古典的な Unix の `fork` 関数を提供していない点に
# 注意してください。ゴルーチンの開始や、プロセスの生成、
# プロセスの `exec` で `fork` のほとんどのユースケースを
# 網羅できるため、普通は問題になりません。
