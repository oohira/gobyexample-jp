# プログラムを実行すると、ハッシュが計算され、
# 人間が読める 16 進フォーマットで表示されます。
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# 先に説明したのと同様のパターンを使って、他のハッシュも計算できます。
# 例えば、SHA512 ハッシュを計算するには、`crypto/sha512`
# をインポートして `sha512.New()` を使います。

# 暗号論的に安全なハッシュが必要な場合には、
# [ハッシュの強度](https://en.wikipedia.org/wiki/Cryptographic_hash_function)
# を注意深く調査すべきという点に注意してください！
