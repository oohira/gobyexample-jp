# 期待通り、まず `"one"`、次に `"two"` を受信しています。
$ time go run select.go 
received one
received two

# 全体の実行時間が高々 2 秒程度であることに注意してください。
# というのも、1 秒と 2 秒の `Sleeps` は並行実行されるからです。
real	0m2.245s
