# このプログラムを実行すると、シグナルを待ってブロックします。
# `ctrl-C` (ターミナルは `^C` と表示する) を入力すると、
# `SIGINT` シグナルを送信することができ、
# プログラムは `interrupt` を表示して終了します。
$ go run signals.go
awaiting signal
^C
interrupt
exiting
