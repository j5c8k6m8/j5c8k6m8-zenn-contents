---
title: "Zennと連携するGitHubリポジトリ名は何がいいのか？"
emoji: "🤔"
type: "idea" # tech: 技術記事 / idea: アイデア
topics: ["zenn"]
published: true
---

# Zenn.devはじめました

記念すべき、初投稿です。


Zenn.devの魅力は数多あると思いますが、私は特に
[GitHubリポジトリでZennのコンテンツを管理する](https://zenn.dev/zenn/articles/connect-to-github)
魅力を感じ、早速使ってみようと思いました。

早速、GitHubリポジトリを作ろうと思ったのですが、ふと

`リポジトリ名は、何にするのがいいのか🤔`

と疑問に思いました。

なお、本記事の結論は、以下となります。

**「zenn-contents」が多いけど、「ユーザ名-zenn-contents」がいいのではないか**

# 他の人のを見てみよう

ということで、早速トレンドに載っているユーザの情報を参考にさせていただきました。

GitHub連携は公開リポジトリではなく、プライベートリポジトリでも実施できるため、

 - GitHub連携(公開リポジトリ)
 - GitHub連携(プライベートリポジトリ)
 - GitHub連携を使わない

 という、3つの編集方法がありますが、

  - アカウントにGitHubユーザ名を登録している
    - ⇒GitHub連携を使っている
  - アカウントのGitHubに「articles」ディレクトリのあるリポジトリがある
    - ⇒GitHub連携(公開リポジトリ)を利用している

と、判断すると、9/19段階での体感では、それぞれ1/3ずつ位なのかなぁと思いました。


さて、本記事公開時は、トレンド -> ユーザページ -> GitHub -> Repositories
で確認していましたが、 **Zenn CLIでは、 `npx zenn init` でREADME.mdを作成する**
ことを考慮すると、GitHubの検索で対象が絞れるのではないかと考えました。

そこで、npx zenn initで作られるREADME内の、zenn-cli-guideのURLを用いて、
以下の検索テキストでGitHubを検索してみました。

`https://zenn.dev/zenn/articles/zenn-cli-guide in:readme`

![](https://github.com/j5c8k6m8/j5c8k6m8-zenn-contents/raw/master/articles/images/zenn-github-repository-name___img-github-search.png)

9/20日時点で、22件のリポジトリがHITしました。
以下、リポジトリ名の多い順のリストです。

 - zenn-contents (6件)
   - [heriet/zenn-content](https://github.com/heriet/zenn-content)
   - [kateinoigakukun/zenn-contents](https://github.com/kateinoigakukun/zenn-contents)
   - [koher/zenn-contents](https://github.com/koher/zenn-contents)
   - [proudust/zenn-contents](https://github.com/proudust/zenn-contents)
   - [ria3100/zenn-contents](https://github.com/ria3100/zenn-contents)
   - [turara/zenn-contents](https://github.com/turara/zenn-contents)
 - zenn (6件)
   - [e-jigsaw/zenn](https://github.com/e-jigsaw/zenn)
   - [mcasashi-aso/zenn](https://github.com/mcasashi-aso/zenn)
   - [nikukyugamer/zenn](https://github.com/nikukyugamer/zenn)
   - [steelydylan/zenn](https://github.com/steelydylan/zenn)
   - [thanaism/zenn](https://github.com/thanaism/zenn)
   - [y-takagi/zenn](https://github.com/y-takagi/zenn)
 - zenn-dev (2件)
   - [koji/zenn-dev](https://github.com/koji/zenn-dev)
   - [skanehira/zenn-dev](https://github.com/skanehira/zenn-dev)
 - zenn-docs
   - [MakotoUwaya/zenn-docs](https://github.com/MakotoUwaya/zenn-docs)
 - zenn-doc
   - [tennashi/zenn-doc](https://github.com/tennashi/zenn-doc)
 - zenn.dev
   - [kou029w/zenn.dev](https://github.com/kou029w/zenn.dev)
 - zenn-articles
   - [regonn/zenn-articles](https://github.com/regonn/zenn-articles)
 - zenn-articles-and-books
   - [tsuyoshicho/zenn-articles-and-books](https://github.com/tsuyoshicho/zenn-articles-and-books)
 - zennlog
   - [d6rkaiz/zennlog](https://github.com/d6rkaiz/zennlog)
 - Blog
   - [Shougo/Blog](https://github.com/Shougo/Blog)
 - alpaca-notes
   - [hhiroshell/alpaca-notes](https://github.com/hhiroshell/alpaca-notes)

実は、Zenn公式の記事のリポジトリ名は **zenn-docs**([zenn-dev/zenn-docs](https://github.com/zenn-dev/zenn-docs))
なのですが、こちらのREADMEには、zenn-cli-guideへのリンクが貼られていなかったので、
上記検索方法ではHITしませんでした。

なお、 **`zenn-contents`は [GitHubリポジトリでZennのコンテンツを管理する](https://zenn.dev/zenn/articles/connect-to-github) 記事**
の説明画像中で使われていました。


# 結局どうしたのか。

私は、zenn-contentsに決めようと思ったのですが、ふと、GitHubで記事を管理する利点の一つに、
**GitHubの機能が使える** つまり、 **記事のプルリクができる** ことに気づきました。

Qiitaでは、編集リクエストの機能があり、これが記事の質を高めるのにある程度貢献していると思っていますが、
今のところ、Zennには、編集リクエストの機能がありません。

しかし、GitHubで記事を管理するれば、プルリクによる編集リクエストが可能です。

**その際に、`zenn-contents`というリポジトリ名だと、、リポジトリ名がかぶってしまいます！！**

つまり、Zennの記事用のGitHubのリポジトリ名は、ユーザが分かるものがよいのではないか？
と考え、私は「ユーザ名-zenn-contents」というリポジトリ名にしました。
(zenn-docsだと、zenn自体のドキュメントというニュアンスが出るので。)

↓ 筆者の作成したzenn投稿記事用GitHubリポジトリ

https://github.com/j5c8k6m8/j5c8k6m8-zenn-contents


# 公開リポジトリ管理のメリット/デメリット

**公開リポジトリで記事を書くと、プルリクがもらえるかもしれない** というメリットがあるため、
公開リポジトリを用いること自体は悪くはないのですが、現在の仕様では１点懸念を感じました。

それは、 [Zenn CLIを使ってコンテンツを作成する](https://zenn.dev/zenn/articles/zenn-cli-guide) の
コメントに気づきがあったのですが、

> yt4
> 有料で本を出す場合は、管理するリポジトリはプライベートにしておいたほうが良い感じですかね？
> 
> Zenn公式
> そうですね。有料販売する場合はプライベートリポジトリにされることをおすすめします。

Zennの差別化の一つである、有料販売では用いにくいことです。
**連携リポジトリは現状1つに限られているため、** 有料販売を考えている人は、注意が必要そうです。

# さいごに

初投稿をしたくて、書き殴りました。

GitHubでの管理には、無限の可能性を感じているので、そこら辺の記事をしばらく投稿したいと思ってます。

私のリポジトリは、Dockerでの構築手順を追記してるので、次はここらへんを。

(あと本記事のリストを作成できるscriptの作成とか)
