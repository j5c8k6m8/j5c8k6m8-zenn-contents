---
title: "🧪exit後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` python:after_exit.py:./projects/python/src/after_exit.py
exit()
print("Am I dead?")

```

`exit` は、多くのプログラミング言語で、処理(プロセス)を終了させるために使われる。呼出により直ちに処理を終了させる目的の関数という点は共通しているが、**後処理の有無



# 言語毎

|言語|実行可否|ツール|検知可否|
|:--|:--|:--|:--|
|Python|実行可|-|-|
|||flake8|検知不可|
|Ruby||-|-|
|||-w(turn warnings)||
|||rubocop||
|JavaScript||-|-|
|||eslint||
|Java||-|-|
|Go||-|-|

## Python

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

``` ruby:template.rb:./projects/ruby/src/template.rb
```

``` console
```

## JavaScript

``` js:template.js:./projects/javascript/src/template.js
```

``` console
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
