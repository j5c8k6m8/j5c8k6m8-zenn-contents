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

`break` は、多くのプログラミング言語で、ループ文から抜け出すために使用される。プログラミング言語によっては、switch文から抜け出すためにも使われる。

`continue` (rubyでは `next`)などのループ文で用いられるジャンプ文を利用した場合も、本パターンに含める。

[`🔖中断を利用するパターン`](./p_after) である。`break` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `break` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `break` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `return` 追加による修正後、不要なコードを消さなかった。 -> `🛐供養`
 - 削除対象にgoto文のラベルが記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に代入が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)

# 言語毎

|言語|実行可否|ツール|検知可否|
|:--|:--|:--|:--|
|Python|実行可|-|-|
|||flake8|検知不可|
|Ruby|実行可|-|-|
|||-w(turn warnings)|検知可|
|||rubocop|検知可|
|JavaScript||-|-|
|||eslint||
|Java||-|-|
|Go||-|-|

## Python

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

## Ruby

Rubyでは、 `break` の他に、 `next` `redo` `retry` がジャンプできる予約語として用意されている。

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

## JavaScript

JavaScriptでは、 `switch` 文中でも `break` が使用される。

``` js:after_break.js:./projects/javascript/src/after_break.js
while(true) {
  break;
  console.log("Am I dead?")
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

## Java

``` java:Template.java:./projects/java/src/main/java/Template.java
```

``` console
```

## Go

``` go:template.go:./projects/golang/src/template.go
```

``` console
```



