---
title: "🧪exit後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting) [`組込モジュールの上書き`](./z_builtin_override) [`組込モジュールの隠蔽`](./z_builtin_hide)|

``` python:after_exit.py:./projects/python/src/after_exit.py
exit()
print("Am I dead?")

```

`exit` は、多くのプログラミング言語で、処理(プロセス)を終了させるために使われる。

`exit` は１つのプログラミング言語内に複数の関数が用意されていることが多い。
これは、一概に処理(プロセス)を終了指せるといっても、以下の様な観点で違いがあるからである。

- リターンコード
- 後処理の有無
- 他スレッドの待受

::: message
`🧪exit後のコード` においては、上記の違いは意識しなくてもよい。
しかし、本書は **デッドコードのパターンを通して 、プログラミングのパターンやスタイルについての考察を深める** ことも目的としているため、これらの機能的な差異については、今後記載を拡充していく予定。
:::

大域脱出という観点では、`throw` と同じである。`exit` は正常の場合にも使われる事が多く、 `throw` は例外処理で多く使われるという違いはある。しかし、後処理を伴う `exit` の場合、 `throw` との厳密な区別は難しいだろう。

例外処理を伴うプログラミングで、リターンコードを意識して、`throw` や 複数の `exit` を使い分けることは難しい。これらの使い分けについては、本書の対象外とする。

なお、多くの言語で、 [`🧪return後のコード`](./r_after_return), [`🧪throw後のコード`](./r_after_throw), [`🧪goto後のコード`](./r_after_goto) は予約語を使用していた事もあり、`⚠検知有` や `🚫実行不可` だったものが、`🧪exit後のコード` では `🆗実行可` で `🆗検知無` であることには注意してほしい。

 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `exit` が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `exit` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)
 - `exit` の定義が上書きされている ->  [`🧟組込モジュールの上書き`](./z_builtin_override)
 - `exit` が内側のスコープで定義されている -> [`🧟組込モジュールの隠蔽`](./z_builtin_hide)


# 言語毎

|🔧言語|🔩ツール|判定|備考|
|:--|:--|:--|:--|
|[`Python`](#🔧python)|-|`🆗実行可`|-|
||`flake8`|`🆗検知無`|-|
|[`Ruby`](#🔧ruby)|-|`⚠警告有`|`warning: statement not reached`|
||`rubocop`|`⚠検知有`|`UnreachableCode`|
|[`JavaScript`](#🔧javascript)|-|`🆗実行可`|-|
||`eslint`|`🆗検知無`|`no-unreachable`|
|[`Java`](#🔧java)|-|`🆗実行可`|`unreachable statement`|
|[`Go`](#🔧go)|-|`🆗検知無`|`unreachable code`|

## Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`

Pythonには3つの `exit` 関数がある[^1]。以下では組み込み関数の `exit` の例を示す。

[^1]: https://docs.pyq.jp/python/library/exit.html

``` python:after_exit.py:./projects/python/src/after_exit.py
exit()
print("Am I dead?")

```

``` console
$ # コード実行
$ python src/after_exit.py 
$ # flake8
$ flake8 src/after_exit.py 
$ 
```

## Ruby

`🔧Ruby` -> `🆗警告無`,  `🔩rubocop` -> `⚠検知有`

rubyでは、Kernelモジュールに、`exit`, `exit!`, `abort` の3つの関数がある[^2]。以下では`exit` の例を示す。

[^2]: https://docs.pyq.jp/python/library/exit.html


``` ruby:after_exit.rb:./projects/ruby/src/after_exit.rb
exit
puts 'Am I dead?'

```

``` console
$ # コード実行
$ ruby src/after_exit.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_exit.rb 
Syntax OK
$ # rubocop
$ rubocop src/after_exit.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
W

Offenses:

src/after_exit.rb:2:1: W: Lint/UnreachableCode: Unreachable code detected.
puts 'Am I dead?'
^^^^^^^^^^^^^^^^^

1 file inspected, 1 offense detected
$ 
```

## JavaScript

`🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `🆗検知無`

node.jsでは、processオブジェクトに、`exit`, `abort` のの関数がある[^3]。以下では組み込み関数の `exit` の例を示す。

[^3]: https://nodejs.org/api/process.html#process_process

``` js:after_exit.js:./projects/javascript/src/after_exit.js
process.exit()
console.log("Am I dead?");

```

``` console
$ # コード実行
$ node src/after_exit.js 
$ # eslint
$ eslint src/after_exit.js 
$ 
```

## Java

`🔧Java` -> `🆗実行可`

クラス `Runtime`[^4] に `exit`, `halt` メソッドが用意されている。

[^4]: https://docs.oracle.com/javase/jp/7/api/java/lang/Runtime.html

`System.exit` は `Runtime.getRuntime().exit` の呼出と同じである[^5]。

[^5]: https://docs.oracle.com/javase/jp/7/api/java/lang/System.html#exit(int)

``` java:AfterExit.java:./projects/java/src/main/java/AfterExit.java
public class AfterExit {
    public static void main(String[] args) {
        try {
            System.exit(0);
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}
```

``` console
# コード実行
$ java src/main/java/AfterExit.java 
$ 
```

## Go

`🔧Go` -> `🆗検知無`

`os.Exit`, `runtime.Goexit` が使用できる[^6]。

[^6]: https://sharpknock.com/posts/programming/golang-exit.html

``` go:after_exit.go:./projects/golang/src/after_exit.go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(0)
	fmt.Println("Am I dead?")
}

```

``` console
$ # コード実行
$ go run src/after_exit.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_exit.go 
$ 
```
