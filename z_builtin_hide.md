---
title: "🧟組込モジュールの隠蔽"
---

::: message
本チャプターは、今後デッドコードレシピの加筆に伴い修正予定です。
:::

**`組込モジュール`** という単語については [`🧟組込モジュールの上書き`](./z_builtin_override) 参照。

[`🧪throw後のコード`](./r_after_throw) 等で利用できるが、 **プログラミング言語に大きく依存する** ため、言語毎にサンプルを示す。


# 言語毎

- `Python (調査中)`
- [`Ruby`](#🔧ruby)
- `JavaScript (調査中)`
- `Java (調査中)`
- `Go (調査中)`


## 🔧Ruby

内側のスコープから名前解決が行われるため、予約語以外であれば、内側のスコープで定義をすることで、 `組込モジュール` を隠蔽できる。

::: message
説明の表現が正確ではないため、今後修正予定
:::

例として、[`🧪throw後のコード`](./r_after_throw) の `raise` の定義を隠蔽するサンプルを示す。

``` ruby:after_override.rb:./projects/ruby/src/zombie/after_throw_hide.rb
def raise
  puts 'hide raise!!'
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
$ ruby src/zombie/after_throw_hide.rb 
hide raise!!
Am I dead?
$ 
```

なお、`🔩rubocop` では、ゾンビ化されているにも関わらず、 `UnreachableCode` が `⚠検知有` であることに注意してほしい。

```
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/zombie/after_throw_hide.rb 
Syntax OK
$ # rubocop
$ rubocop src/zombie/after_throw_hide.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
W

Offenses:

src/zombie/after_throw_hide.rb:7:3: W: Lint/UnreachableCode: Unreachable code detected.
  puts 'Am I dead?'
  ^^^^^^^^^^^^^^^^^
src/zombie/after_throw_hide.rb:8:1: W: Lint/SuppressedException: Do not suppress exceptions.
rescue RuntimeError
^^^^^^^^^^^^^^^^^^^

1 file inspected, 2 offenses detected
$ 
```
