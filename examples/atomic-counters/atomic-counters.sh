# ちょうど 50,000 オペレーションになるはずです。仮に
# アトミックでない `ops++` を使ってカウンターをインクリメントすると、
# ゴルーチン同士が互いに干渉するため、実行するたびに異なる数値が
# 出力されるはずです。さらに、`-race` フラグありで実行すると
# データ競合エラーになるでしょう。
$ go run atomic-counters.go
ops: 50000

# 次は、状態を管理するもう 1 つの方法である、
# ミューテックスを見ていきましょう。
