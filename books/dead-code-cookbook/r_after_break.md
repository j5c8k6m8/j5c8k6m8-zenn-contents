---
title: "🧪break後のコード"
---


|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` python:after_break.py:./projects/python/src/after_break.py
while True:
    break
    print("Am I dead?")

```

`break` は、多くのプログラミング言語で、 **ループ文から抜け出すために使用** する。プログラミング言語によっては、 `switch文` から抜け出すためにも使う。

`continue` などのループ文で用いられるジャンプ文を利用した場合も、本レシピに含める。

[`🔖中断を利用するパターン`](./p_after) である。`break` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `break` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `break` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `break` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)

# 言語毎

|🔧言語|🔩ツール|判定|備考|
|:--|:--|:--|:--|
|[`Python`](#🔧python)|-|`🆗実行可`|-|
||`flake8`|`🆗検知無`|-|
|[`Ruby`](#🔧ruby)|-|`⚠警告有`|`warning: statement not reached`|
||`rubocop`|`⚠検知有`|`UnreachableCode`|
|[`JavaScript`](#🔧javascript)|-|`🆗実行可`|-|
||`eslint`|`⚠検知有`|`no-unreachable`|
|[`Java`](#🔧java)|-|`🚫実行不可`|`unreachable statement`|
|[`Go`](#🔧go)|-|`⚠検知有`|`unreachable code`|

## 🔧Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`

`switch文` は提供していない。

``` python:after_break.py:./projects/python/src/after_break.py
while True:
    break
    print("Am I dead?")

```

``` console
$ # コード実行
$ python src/after_break.py 
$ # flake8
$ flake8 src/after_break.py 
$ 
```

## 🔧Ruby

`🔧Ruby` -> `⚠警告有`,  `🔩rubocop` -> `⚠検知有`

Rubyでは、`continue` ではなく、 `next` を用いる。また、 `break` , `next` の他に、 `redo` `retry` がジャンプできる予約語として用意されている。

`switch文` である `case` では `break` は使わない。

``` ruby:after_break.rb:./projects/ruby/src/after_break.rb
while true
  break
  puts 'Am I dead?'
end

```

``` console
$ # コード実行
$ ruby src/after_break.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_break.rb 
src/after_break.rb:3: warning: statement not reached
Syntax OK
$ # rubocop
$ rubocop src/after_break.rb 
Inspecting 1 file
W

Offenses:

src/after_break.rb:1:1: W: Lint/UnreachableLoop: This loop will have at most one iteration.
while true ...
^^^^^^^^^^
src/after_break.rb:1:1: C: Style/InfiniteLoop: Use Kernel#loop for infinite loops.
while true
^^^^^
src/after_break.rb:3:3: W: Lint/UnreachableCode: Unreachable code detected.
  puts 'Am I dead?'
  ^^^^^^^^^^^^^^^^^

1 file inspected, 3 offenses detected, 1 offense auto-correctable
$ 
```

## 🔧JavaScript

`🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `⚠検知有`

`switch` 文中で `break` が使用できる。

``` js:after_break.js:./projects/javascript/src/after_break.js
while(true) {
  break;
  console.log("Am I dead?");
}

```

``` console
$ # コード実行
$ node src/after_break.js 
$ # eslint
$ eslint src/after_break.js 

/app/javascript/src/after_break.js
  1:7  error  Unexpected constant condition  no-constant-condition
  3:3  error  Unreachable code               no-unreachable

✖ 2 problems (2 errors, 0 warnings)

$ 
```

## 🔧Java

`🔧Java` -> `🚫実行不可`

`switch` 文中で `break` が使用できる。

``` java:AfterBreak.java:./projects/java/src/main/java/AfterBreak.java
public class AfterBreak {
    public static void main(String[] args) {
        while(true) {
            break;
            System.out.println("Am I dead?");
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterBreak.java 
src/main/java/AfterBreak.java:5: error: unreachable statement
            System.out.println("Am I dead?");
            ^
1 error
error: compilation failed
$ 
```

また、　`label付break` が使用できる。この場合 `break` の直後ではなくても、以下の様なコードはデッドコードの理由により、コンパイルに失敗する。

``` java:AfterBreakLabel.java:./projects/java/src/main/java/AfterBreakLabel.java
public class AfterBreakLabel {
    public static void main(String[] args) {
        l:{
            while(true) {
                while(true) {
                    break l;
                }
                System.out.println("Am I dead?");
            }
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterBreakLabel.java 
src/main/java/AfterBreakLabel.java:8: error: unreachable statement
                System.out.println("Am I dead?");
                ^
1 error
error: compilation failed
$ 
```


## 🔧Go

`🔧Go` -> `⚠検知有`

``` go:after_break.go:./projects/golang/src/after_break.go
package main

import "fmt"

func main() {
	for {
		break
		fmt.Println("Am I dead?")
	}
}

```

``` console
$ # コード実行
$ go run src/after_break.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_break.go 
# command-line-arguments
src/after_break.go:8:3: unreachable code
$ 
```
