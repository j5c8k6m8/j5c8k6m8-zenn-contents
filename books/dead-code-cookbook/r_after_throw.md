---
title: "🧪throw後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` python:after_throw.py:./projects/python/src/after_throw.py
try:
    raise Exception()
    print("Am I dead?")
except Exception:
    pass

```

`throw` は、多くのプログラミング言語で **例外処理のために大域脱出を行う機能** として用意されている。
Pythonでは `raise`、Goでは `panic`の予約後を用いる。RubyでもPythonと同じ `raise` だが、Kernelモジュールのメソッドとして提供されている。
例外処理と大域脱出を分離する(別のステートメントを使う)プログラミング言語もある。Rubyにおいては、例外処理と似た大域脱出として、`throw` が `raise` と同様Kernelモジュールのメソッドとして提供されている。Rubyの `throw` も本パターンには含めるが、例外処理以外の大域脱出は使われる機会が少ないため、コードは記載しない。
`assert` も例外処理の大域脱出を行う機能であるが、デバッグ用途であり通常の例外処理とは大きく目的が異なるため、 [`🧪常に条件がfalseとなるassert後のコード`](./r_after_assert) として別レシピとする

[`🔖中断を利用するパターン`](./p_after) である。`throw` 後に書かれているコードは、基本的にはデッドコードとなる。

 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `throw` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `throw` 追加による修正後、不要なコードを消さなかった。 -> `🛐`
 - 削除対象にgoto文のラベルが記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に代入が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)
 - `throw` の定義が上書きされている ->  [`🧟定義の上書き`](./z_override_def)
 - `throw` が宣言されている -> [🧟内側のスコープによる隠蔽](./z_override_scope)


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

実行可能で、 `flake8` を使っても検知はできない。

``` python:after_throw.py:./projects/python/src/after_throw.py
try:
    raise Exception()
    print("Am I dead?")
except Exception:
    pass

```

``` console
$ python src/after_throw.py 
$ flake8 src/after_throw.py
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
