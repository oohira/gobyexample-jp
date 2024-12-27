# 起動されたプログラムは、それらをコマンドラインから
# 直接実行したときと同じ出力を返します。
$ go run spawning-processes.go
> date
Thu 05 May 2022 10:10:12 PM PDT

# date コマンドは `-x` フラグをもたないため、
# エラーメッセージと非ゼロのエラーコードで終了します。
command exited with rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
