# プログラムを実行すると、ハッシュが計算され、
# 人間が読める 16 進フォーマットで表示されます。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 先に説明したのと同様のパターンを使って、他のハッシュを計算することができます。
# 例えば、MD5 ハッシュを計算するには、`crypto/md5` をインポートして
# `md5.New()` を使います。

# 暗号論的に安全なハッシュが必要な場合には、
# [ハッシュの強度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)
# を注意深く調査すべきである点に注意してください！
