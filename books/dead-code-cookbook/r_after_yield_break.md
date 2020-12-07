---
title: "🧪breakするコールバック関数呼出後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move) [`関数修正`](./a_func_fix) [`関数呼出追加`](./a_func_add)|
|🧟|-|

``` ruby:🚩 after_yield_break.rb:./projects/ruby/src/after_yield_break.rb
def f
  yield
  puts 'Am I dead?'
end

f do
  break
end

```

|🔧言語|🔩ツール|🚩|note|
|:--|:--|:--|:--|
|Python|-|-|対象外|
|[`Ruby`](#🔧ruby)|-|🆗|-|
||`rubocop`|🆗|-|
|JavaScript|-|-|対象外|
|Java|-|-|対象外|
|Go|-|-|対象外|

`🧪breakするコールバック関数呼出後のコード` は、`break` を使用するという点では、[`🧪break後のコード`](./r_after_break) と同じだが、 `break` 後のコードを対象とはしていない。`break` を含む関数の呼出元の後続の処理がデッドコードとなる。

関数の呼出を行い、処理を中断するという点においては、[`🧪throw後のコード`](./r_after_throw) と同じである。しかし、[`🧪throw後のコード`](./r_after_throw) とは中断させる方法が異なる。

`イテレータ` を利用する時、コールバック関数側で途中で処理を終了させたいことがあるだろう。このような事を目的として、 `break` を用いてコールバック関数内で呼出元関数の後続の処理を終了することができるプログラミング言語がある。本レシピは、そのような機能を保持しているプログラミング言語で実行可能なレシピである。


 - 呼出対象の関数、関数の引数、呼出先の関数の処理が誤っている。 -> [`👼関数修正`](./a_func_fix)
 - 別のコールバック関数を使用する関数呼出追加 -> [`👼関数呼出追加`](./a_func_add)
 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で 関数呼出 が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - 関数呼出 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)


# 🔧Ruby

`🔧Ruby` -> `🆗警告無`,  `🔩rubocop` -> `🆗検知有`

``` ruby:🚩 after_yield_break.rb:./projects/ruby/src/after_yield_break.rb
def f
  yield
  puts 'Am I dead?'
end

f do
  break
end

```

``` console
$ # コード実行
$ ruby src/after_yield_break.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_yield_break.rb 
Syntax OK
$ # rubocop
$ rubocop src/after_yield_break.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
.

1 file inspected, no offenses detected
$ 
```
