# 実行したプログラムは、5 つのジョブが様々なワーカーで
# 実行されていることを示します。このプログラムは、
# 合計で 5 秒分のタスクを実行するにもかかわらす、
# 2 秒しかかかりません。
# というのも、3 つのワーカーが並行実行されるからです。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
