---
title: "はじめに"
---

# はじめに

正規表現は文字列を文字列を評価・操作するために、広く使われています。ほとんどのプログラミング言語では、標準ライブラリや正規表現リテラルで、正規表現を利用可能でしょう。

正規表現は文字列のマッチに広く使われます。そのため、正規表現は、文字列の集合を表わす方法と捉える事も出来ます。正規表現は、現在のプログラミングにおいて、 **文字列の集合を表わすのに最も使われている方法** と言っても過言では無いでしょう。

しかし、 **正規表現では論理演算が論理和のみが容易であり、否定や論理積を示す演算には工夫が必要です。 [^1]**
否定や論理積を確認を、正規表現の外側のプログラムで実行する事も多いでしょう。

[^1]: 正規表現でも、先読みと否定先読みを使えば、否定や論理積の演算を実行できます。
参考) https://zenn.dev/j5c8k6m8/scraps/6d8b5816d0f4ee#comment-50371cb2ef06f1

また、正規表現はパターンマッチング用途に使用されるため、その正規表現自体にマッチする要素(文字列)の存在を保証するものではありません。
(とはいっても、否定や論理積の演算が出来ないので、要素が存在しない正規表現は限られます。)

もし、 **正規表現で論理演算できるとしたら、どうでしょうか？**
例えば、デッドコードの検出です。
複数の正規表現を用いた条件分岐においても、簡単にデッドコードの検出ができるでしょう。
例えば、文字列クラスの拡張です。
正規表現を利用した文字列クラスを定義し、かつ、クラス間の論理的整合性も担保できる型システムが実現可能でしょう。
**もっと堅牢なプログラミング** が実現可能ではないでしょうか？

しかしながら、正規表現での論理演算は工夫が必要で、一般的ではありません。筆者は、正規表現同士の論理演算が直接できるライブラリを見たことがありません。[^2]
正規表現同士の論理演算が可能な正規表現エンジンは作れないのでしょうか？

[^2]: 本書執筆後、twitterで論理演算ができるライブラリとして、[brics](https://www.brics.dk/automaton/doc/index.html?dk/brics/automaton/RegExp.html) の情報をいただきました。
https://twitter.com/Mi_Sawa/status/1404347444297490438?ref_src=twsrc%5Etfw

そんなことはありません。
本書では、 **正規表現の見方を少し変えて、論理演算可能な正規表現が実装できる** ことを示します。


# 前提

本書の執筆において、正規表現についての理解を深める際に、以下の記事を参考にしています。タイトルも以下の記事のオマージュとなります。

**[正規表現エンジンを作ろう - CodeZine](https://codezine.jp/article/corner/237)**

本書の一番の目的は、論理演算が可能な正規表現エンジンの実装を示すことです。そのため、正規表現の基本的な知識、とりわけ、上記記事の知識を前提とします。


# 完成品

本書で作成する正規表現エンジンは、Pythonパッケージ `spexm8p` としてPyPIに登録済みです。使用するだけであれば、以下のように `pip install spexm8p` 後、 `from spexm8p.builder import spex` を記述してimportすることで、 正規表現を利用した文字列集合クラス `Spex` を利用することが可能です。

``` console:sample
# pip install spexm8p
Collecting spexm8p
  Downloading spexm8p-0.0.4-py3-none-any.whl (8.2 kB)
Installing collected packages: spexm8p
Successfully installed spexm8p-0.0.4
# python3
Python 3.9.2 (default, Mar  4 2021, 15:30:45)
[GCC 8.3.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> from spexm8p.builder import spex
>>> # 文字列'a'の補集合は、'a'以外の1文字、もしくは、2文字以上の文字列であることの確認
>>> spex('!a') == spex('[^a]|..+')
True
>>>
```

なお、 `spexm8p` は、論理演算可能な正規表現エンジンのプロトタイプになります。最適化等が不充分であることや、今後、破壊的な更新を行う可能性もあるため、プロダクトでの利用は行わないでください。
