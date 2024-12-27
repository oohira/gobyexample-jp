# サーバーをバックグラウンドで実行します。
$ go run http-server.go &

# `/hello` にアクセスします。
$ curl localhost:8090/hello
hello
