---
title: "🧪goto後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

`goto` は、指定したラベルにジャンプするために使用する。

ラベルを利用する点で、`ラベル付break` と共通点があるが、 `ラベル付break` を利用したデッドコードレシピは [`🧪break後のコード`](./r_after_break) に含める。

`goto` は自由度が高いため、可読性を損なう書き方ができる。そのため、 `goto` を採用している言語は限られる。

なお、本書の [`🔖中断を利用するパターン`](./p_after) は、 **中断された処理以降に記載されているコードを、デッドコード** としているため、 `goto` 後にラベルを記載しているケースは [`🧟goto文のラベルによるジャンプ`](./z_goto) となる。つまり **`goto` を用いると自然と本書のゾンビ化のパターンを利用してしまう** ことがある。このことは、 `goto` の自由度が高く、可読性を損なう書き方ができるということの一例といえるだろう。

`goto`は、その自由度の高さから、デバッグで使われることもあるだろう。

 - デバッグのために残した `goto` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `goto` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `goto` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)


# 言語毎

|🔧言語|🔩ツール|判定|備考|
|:--|:--|:--|:--|
|Python|-|対象外|-|
|Ruby|-|対象外|-|
|JavaScript|-|対象外|-|
|Java|-|対象外|-|
|[`Go`](#🔧go)|-|`⚠検知有`|`unreachable code`|

## 🔧Go

`🔧Go` -> `⚠検知有`

``` go:after_goto.go:./projects/golang/src/after_goto.go
package main

import "fmt"

func main() {
	goto L
	fmt.Println("Am I dead?")
L:
	return
}

```

``` console
$ # コード実行
$ go run src/after_goto.go
$ # 静的解析(標準ツール)実行
$ go vet src/after_goto.go 
# command-line-arguments
src/after_goto.go:7:2: unreachable code
$ 
```
