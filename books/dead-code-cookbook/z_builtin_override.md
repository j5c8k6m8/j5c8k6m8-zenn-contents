---
title: "🧟組込モジュールの上書き"
---

::: message
本チャプターは、今後デッドコードレシピの加筆に伴い修正予定です。
:::

**`組込モジュール`** とは、`組込メソッド` や `標準ライブラリ`を含んだ、 `分別のあるプログラミング` では挙動が変らないと推測される機能を指す言葉とする。なお、`演算子のオーバロード` は含まない。

[`🧪throw後のコード`](./r_after_throw) 等で利用できるが、 **プログラミング言語に大きく依存する** ため、言語毎にサンプルを示す。


# 言語毎

- `Python (調査中)`
- [`Ruby`](#🔧ruby)
- `JavaScript (調査中)`
- `Java (調査中)`
- `Go (調査中)`


## 🔧Ruby

`オープンクラス` による定義の上書きが可能。

::: message
オープンクラス以外の定義の上書き方法を追記したい
:::

例として、[`🧪throw後のコード`](./r_after_throw) の `raise` の定義を上書きするサンプルを示す。

``` ruby:after_override.rb:./projects/ruby/src/zombie/after_throw_override.rb
module Kernel
  def raise
    puts 'override raise!!'
  end
end

begin
  raise
  puts 'Am I dead?'
rescue RuntimeError
end

```

実行結果は、以下の通りゾンビ化できている。

``` console
$ # コード実行
$ ruby src/zombie/after_throw_override.rb 
override raise!!
Am I dead?
$ 
```

なお、`🔩rubocop` では、ゾンビ化されているにも関わらず、 `UnreachableCode` が `⚠検知有` であることに注意してほしい。

```
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/zombie/after_throw_override.rb 
Syntax OK
$ # rubocop
$ rubocop src/zombie/after_throw_override.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
W

Offenses:

src/zombie/after_throw_override.rb:9:3: W: Lint/UnreachableCode: Unreachable code detected.
  puts 'Am I dead?'
  ^^^^^^^^^^^^^^^^^
src/zombie/after_throw_override.rb:10:1: W: Lint/SuppressedException: Do not suppress exceptions.
rescue RuntimeError
^^^^^^^^^^^^^^^^^^^

1 file inspected, 2 offenses detected
$ 
```
