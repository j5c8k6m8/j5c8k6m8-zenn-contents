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

`return` は、多くのプログラミング言語で、処理を呼出元に戻すために使われる。関数において使われることが多いが、スクリプト言語ではトップレベルのスクリプト環境でも実行できるものもある。例えば、pythonはモジュールで `return` を使用できないため、トップレベルのスクリプト環境では使えないが、rubyは使える。

[`🔖中断を利用するパターン`](./p_after) である。`return` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `return` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `return` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `return` 追加による修正後、不要なコードを消さなかった。 -> `🛐供養`
 - 削除対象にgoto文のラベルが記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象にホイスティング対象が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)

# 言語毎

|言語|実行可否|ツール|検知可否|
|:--|:--|:--|:--|
|Python|実行可|-|-|
|||flake8|検知不可|
|Ruby|実行可|-|-|
|||-w(turn warnings)|検知可能|
|||rubocop|検知可能|
|JavaScript|実行可|-|-|
|||eslint|検知可能|
|Java|実行不可|-|-|
|Go|実行可|-|-|

## Python

実行可能で、 `flake8` を使っても検知はできない。
トップレベル(モジュール) では `return` は使えないことに注意

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

## Ruby

トップレベルでも `return` が可能。
`-w` および `rubocop` で検知可能

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

## JavaScript

nodeにて、トップレベルの `return` は事項できるが、`eslint`にて、 `error  Parsing error: 'return' outside of function` が発生するため、本書では関数内でreturnするコードを示す。

``` js:after_return.js:./projects/javascript/src/after_return.js
(() => {
    return
    console.log("Am I dead?")
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

## Java

javaでは、実行不可となる。以下のコードは実行できない。

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

## Go

Goでは、実行可能である。

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
$ 
```