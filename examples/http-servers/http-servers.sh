# サーバーをバックグラウンドで実行します。
$ go run http-servers.go &

# `/hello` にアクセスします。
$ curl localhost:8090/hello
hello
