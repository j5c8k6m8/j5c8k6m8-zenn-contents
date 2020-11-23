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

[`🔖中断を利用するパターン`](./p_after) である。`return` 後に書かれているコードは、基本的にはデッドコードとなる。

 - デバッグのために残した `return` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `return` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `return` 追加による修正後、不要なコードを消さなかった。 -> `🛐`
 - goto文のラベルを記載している。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - `🛐` -> [`🧟ホイスティング`](./z_hoisting) に注意

# 言語一覧

|言語|実行可否|ツール|検知可否|
|:--|:--|:--|:--|
|Python|実行可|-|検知不可|
|||flake8|検知不可|
|Ruby|実行可|-|検知不可|
|||-w(turn warnings)|検知可能|
|||rubocop|検知可能|

# 言語毎の詳細

## Python

トップレベル(モジュール) では `return` は使えないことに注意

``` python:after_return.rb:./projects/python/src/after_return.py
def f():
    return
    print("Am I dead?")


f()

```

``` console
$ python /app/python/src/after_return.py
$ flake8 /app/python/src/after_return.py
$ 
```

## Ruby

``` ruby:after_return.rb:./projects/ruby/src/after_return.rb
return
puts 'Am I dead?'
```
