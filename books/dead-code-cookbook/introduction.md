---
title: "はじめに"
---

# 本書の目的

デッドコードは、 **プログラムの一部として存在するが、決して実行されないコード** のことである。一時的な変更等でデッドコードを作る場合もあるが、多くの場合は **バグ（プログラム自体の誤り）の可能性が極めて高い** だろう。例えば 「**goto fail bug**[^gotofailbug]」 がデッドコードを含んだ有名なバグの実例だ。

[^gotofailbug]: https://ja.wikipedia.org/wiki/%E5%88%B0%E9%81%94%E4%B8%8D%E8%83%BD%E3%82%B3%E3%83%BC%E3%83%89#%E8%84%86%E5%BC%B1%E6%80%A7%E3%81%A8%E5%88%B0%E9%81%94%E4%B8%8D%E8%83%BD%E6%80%A7

そのため、理想的にはプログラミング言語レベルでデッドコードを検知し、警告出力や実行不可などの対応をとることが望ましい。実際に、特定のパターンのデッドコードは言語仕様上許されていない場合もある。

しかし、任意のコードが到達不能(デッドコード)かどうかを判断することは停止性問題を解くことと等価であり、 **あらゆるデッドコードを正しく把握することは不可能**[^deadcodeinspect] である。デッドコードを完全に排除している汎用プログラミング言語は、今のところ存在していないだろう。

[^deadcodeinspect]: 出典) https://ja.wikipedia.org/wiki/%E5%88%B0%E9%81%94%E4%B8%8D%E8%83%BD%E3%82%B3%E3%83%BC%E3%83%89#%E5%88%B0%E9%81%94%E4%B8%8D%E8%83%BD%E6%80%A7%E3%81%AE%E6%A4%9C%E8%A8%BC

**では、実際のプログラミングのコードには、どのようなデッドコードのパターンがあるのだろうか？**

この疑問の答えを求めるのが本書である。
しかしながら、 **デッドコードのパターンから何が得られるのだろうか？**

デッドコードは除去すべきものであり、デザインパターンのように、再利用を促すものではない。しかし、デッドコードのパターンから得られるものもある。

デッドコードのパターンを整理するためには、プログラミングのパターンを整理する必要がある。それは、単なるデッドコードを回避するという直接的な理由以外に、 **プログラミングのパターンやスタイルについての考察を深める** ことにも繋がる。

また、デッドコードが実行可能かどうかは、プログラミング言語によっても異なる。また、あるプログラミング言語ではデッドコードとなるパターンであっても、高度な機能を用いることで、別のプログラミング言語ではコードを通すことが出来る場合もある。そのため、デッドコードのパターンを整理し、プログラミング言語毎の具体例を確認することで、 **各プログラミング言語への理解を深める** ことにも繋がる。

あるいは、実際のプロジェクトで、前任者が書いたデッドコードに出くわした経験がある人も多いだろう。デッドコードの除去はソースの可読性を向上させ、保守のしやすさにつながる。しかし、デッドコードの除去で動作が変わらない（影響がない）と言えるだろうか？リスク回避という観点からデッドコードの除去を断念することもあるだろう。もし、該当のデッドコードに関する考察がすでになされていて、注意点がまとまっていればどうだろうか？ **正確なデッドコードの除去を後押しする** ことができるだろう。

これらが、本書の目的である。

つまり、 **デッドコードのパターンを通して** 、プログラミングのパターンやスタイルについての考察を深め、各プログラム言語への理解を深め、正確なデッドコードの除去を後押しし、 **より良いコードを書く** ことである。

本書のタイトルは「**DEAD CODE COOKBOOK**」、デッドコードのレシピ集である。

しかし、本書は **意図的にデッドコードを作るためのレシピ集ではない。**

見つけ難い不具合を意図的に埋め込むためではなく、 **より良いコードを書く** ためにのみ活用してほしい。

# 対象範囲と用語の定義

## ☠️デッドコード☠️

デッドコード(dead coad)は到達不能なコードである。類語の到達不能コード(unreachable code)との使い分けは行わず、本書ではデッドコードで統一する。

本書のデッドコードは、 **高級言語の、処理が伴うソースコード** を対象とする。

実行時のコード(機械語、あるいは元のプログラムよりも低い水準のコード)は対象外として考慮しない。大半のオプティマイザで削除されるであろう、未参照変数に対する代入は、実行時の議論であり、処理が伴わないコードのため、本書のデッドコードには含めない。

デッドコードは厳密に到達不能のコードのみではなく、後述するデッドコードレシピに基づいたコードであれば、デッドコードとする。

## 🧪デッドコードレシピ🧪

デッドコードの作り方、つまり、デッドコードを実現できるプログラミング言語の機能の組み合わせを、デッドコードレシピと定義する。

デッドコードレシピに従ったパターンによるソースコードであれば、デッドコードとする。つまり、厳密には到達可能であっても、分別のあるプログラミングでは到達不能と認識されるソースコードはデッドコードとする。たとえば、プロトタイプ汚染のみにより到達可能なコードはデッドコードである。

## 🛐供養🛐

デッドコードの該当部分のみを削除することを、供養と定義する。

通常、デッドコードは存在しないほうが望ましい。デッドコードを削除するにあたり、挙動を変えない場合は該当部分のみを削除、つまり供養すればよい。

ただし、 **供養してはいけないケースがある** 点には注意が必要だ。そもそも、バグによりデッドコードになっている場合は、供養により挙動は変わらないが、バグが残る。バグを残さないためには、後述のレイズによりデッドコードが実行されるように修正する必要がある。

また、一見到達不能に見えるが、厳密には到達可能なデッドコードの場合は、供養により挙動が変わってしまう。後述のゾンビ化によりソースコードが実行されていないかを確認する必要がある。

供養の方法は簡単だが、供養の判断は難しく、時として勇気が必要となる。

本書では、供養の方法は逐一記載しない。しかし、後述する復活の呪文はなるべく詳細に記載する。
それは、復活の呪文の知識が、供養の判断の後押しをするためである。

しかし、供養を躊躇ってもよいが、怠ってはいけない。大半のデッドコードは供養するのが正解である。



## 🧙復活の呪文🧙

デッドコードを実行するための方法を、復活の呪文とする。

復活の呪文は、２種類に分かれる。ひとつは、該当箇所のソースコードの書き換えを伴うもので、これをレイズとする。もうひとつは、該当箇所のソースコードの書き換えを伴わないもので、これをゾンビ化とする。

なお、復活させずにソースコードから削除する方法は、前述の供養である。供養は復活の呪文ではない。


## 👼レイズ👼

ソースコードを書き換え、デッドコードを到達可能にする方法をレイズとする。レイズされたデッドコードは、デッドコードレシピから外れるため、デッドコードではなくなる。

レイズは、デッドコードを作り込んだ理由を想定し、ソースコードを正しい状態にする。

プログラムの挙動が変わるため、レイズの正しさは、個別のコード毎に判断する必要がある。そして、多くの場合は供養が正解であり、レイズが正解ではない点に注意してほしい。


## 🧟ゾンビ化🧟


## 🐕ケルベロス🐕



# サンプルコードと実行環境

未使用のデータ項目に対するデッドコードは対象外

I am already dead!!

## 対象言語と環境

|プログラミング言語|Dockerタグ|
|:--|:--|
|C言語||
|Python||
|JavaScript(Node.js)||
|Ruby|ruby:2.7.2-buster|
|Go|1.15.4-buster|
|Java|openjdk:15.0.1|
|TypeScript||
|Rust||

環境構築は別章とする。

``` console
docker run --rm -v $ZENN_HOME/books/dead-code-cookbook/src/ruby:/app/ruby ruby:2.7.2-buster ruby /app/ruby/hello_world.rb && echo success
```

``` console
docker run --rm -v $ZENN_HOME/books/dead-code-cookbook/src/golang:/app/golang golang:1.15.4-buster go run /app/golang/hello_world.go && echo success
```

# 読者へのおねがい

筆者はプログラミングの経験が不十分。（特に高度な機能）
パターンも他に有るだろうし、言語のサンプルが充分でないこともある。
プルリクをして欲しい。もしくは、Qiitaでもいいだろう。ハッシュタグで呟いてもいいし、

2020年度のアドベンドカレンダー形式の本とする試み。

これは、時間をかけてじっくり読んでもらう事で、いろんな意見がもらえるのではないか？という意図がある。

全てに目を通すことは難しいかもしれないが、時間をかけて可能な限り本書に反映していくつもりがある。

# レシピ(目次)

1.リテラルを利用した条件式
2.ローカル変数を利用した条件式
1.定数を利用した条件式
3.return後のコード
4.throw後のコード
5.break/continue後のコード
6.条件の重ねあわせ。分岐のネスト
7.条件の重ね合わせ。switch
8.空配列のeach
9.正規表現の利用
10.プライベートメソッドの活用
11.関数の上書き
12.メソッドとオーバーライド
13.メソッドチェーンの利用
14.整数値に対する演算の使用 範囲
15.整数値に対する演算の使用 2じょう
16.文字列の活用
17.マクロの利用
18.組み込み例外の利用
19.ソースコードの取得とevalの利用
20.catchの利用
21.gotoの利用
22.never型の利用？
23.ワンライナーの利用"goto fail bug"
24.三項演算子
25.プロミス？
26.exit。大域脱出？