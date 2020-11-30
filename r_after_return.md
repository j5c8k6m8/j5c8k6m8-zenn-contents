---
title: "🧪return後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` ruby:after_return.rb:./projects/ruby/src/after_return.rb
return
puts 'Am I dead?'

```

`return` は、多くのプログラミング言語で、処理を呼出元に戻すために使われる。関数において使われることが多いが、スクリプト言語ではトップレベルのスクリプト環境でも実行できるものもある。

[`🔖中断を利用するパターン`](./p_after) である。`return` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `return` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `return` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `return` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto` 文のラベルが記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象にホイスティング対象が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)

# 言語毎

|🔧言語|🔩ツール|判定|備考|
|:--|:--|:--|:--|
|[`Python`](#🔧Python)|-|`🆗実行可`|-|
||`flake8`|`🆗検知無`|-|
|[`Ruby`](#🔧Ruby)|-|`⚠警告有`|-|
||`rubocop`|`⚠検知有`|-|
|[`JavaScript`](#🔧JavaScript)|-|`🆗実行可`|-|
||`eslint`|`⚠検知有`|-|
|[`Java`](#🔧Java)|-|`🚫実行不可`|-|
|[`Go`](#🔧Go)|-|`⚠検知有`|-|

## 🔧Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`

トップレベル(モジュール) では `return` を使えない。

``` python:after_return.py:./projects/python/src/after_return.py
def f():
    return
    print("Am I dead?")


f()

```

``` console
$ # コード実行
$ python src/after_return.py
$ # flake8
$ flake8 src/after_return.py
$ 
```

## 🔧Ruby

`🔧Ruby` -> `⚠警告有`,  `🔩rubocop` -> `⚠検知有`

トップレベルでも `return` が可能。

``` ruby:after_return.rb:./projects/ruby/src/after_return.rb
return
puts 'Am I dead?'

```

``` console
$ # コード実行
$ ruby src/after_return.rb
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_return.rb
src/after_return.rb:2: warning: statement not reached
Syntax OK
$ # rubocop
$ rubocop src/after_return.rb
Inspecting 1 file
W

Offenses:

src/after_return.rb:2:1: W: Lint/UnreachableCode: Unreachable code detected.
puts 'Am I dead?'
^^^^^^^^^^^^^^^^^

1 file inspected, 1 offense detected
$ 
```

## 🔧JavaScript

`🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `⚠検知有`

`🔧node.js` でトップレベルの `return` は実行可能。しかし、`🔩eslint` で、 `error  Parsing error: 'return' outside of function` が発生するため、関数内で `return` するコードを示す。

``` js:after_return.js:./projects/javascript/src/after_return.js
(() => {
    return;
    console.log("Am I dead?");
})()
```

``` console
$ # コード実行
$ node src/after_return.js
$ # eslint
$ eslint src/after_return.js

/app/javascript/src/after_return.js
  3:5  error  Unreachable code  no-unreachable

✖ 1 problem (1 error, 0 warnings)

$ 
```

## 🔧Java

`🔧Java` -> `🚫実行不可`

``` java:AfterReturn.java:./projects/java/src/main/java/AfterReturn.java
public class AfterReturn {
    public static void main(String[] args) {
        return;
        System.out.println("Am I dead?");
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterReturn.java 
src/main/java/AfterReturn.java:5: error: unreachable statement
        System.out.println("Am I dead?");
        ^
1 error
error: compilation failed
$ 
```

## 🔧Go

`🔧Go` -> `⚠検知有`

``` go:after_return.go:./projects/golang/src/after_return.go
package main

import "fmt"

func main() {
	return
	fmt.Println("Am I dead?")
}

```

``` console
$ # コード実行
$ go run src/after_return.go
$ # 静的解析(標準ツール)実行
$ go vet src/after_return.go 
# command-line-arguments
src/after_return.go:7:2: unreachable code
$ 
```