---
title: "🧪break後のコード"
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

`break` は、多くのプログラミング言語で、ループ文から抜け出すために使用される。プログラミング言語によっては、switch文から抜け出すためにも使われる。

`continue` などのループ文で用いられるジャンプ文を利用した場合も、本パターンに含める。

[`🔖中断を利用するパターン`](./p_after) である。`break` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `return` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 文等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `return` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `return` 追加による修正後、不要なコードを消さなかった。 -> `🛐供養`
 - 削除対象にgoto文のラベルが記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に代入が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)


continue(rubyではnext)を含む

redo / retry

