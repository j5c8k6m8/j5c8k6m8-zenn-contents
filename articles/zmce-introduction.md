---
title: "Zennの執筆体験を向上させるnpm「zmce」を公開した"
emoji: "🔧"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: [ "Zenn", "md", "markdownit", "zmce" ]
published: true
---

# 「zmce」とは

**Zenn Markdown Code Embed** の頭文字を取った **npmパッケージ** です。

**コマンドひとつで、Zennの記事/本 のコードブロックを更新する** ためのツールです。

:::message
本記事は、 **v0.2.0以降** のモジュールに対する説明です。
バージョンは **`npm info zmce`** で確認し、v0.2.0以前のモジュールの場合、
**`npm install zmce@latest`** でアップデートしてください。
:::

# 前提環境

 1. **[Zenn](https://zenn.dev/)** で **記事** または **本** を執筆している。
 2. **[GitHubリポジトリでZennのコンテンツを管理](https://zenn.dev/zenn/articles/connect-to-github)** している。

# 想定ユーザ

 1. 記事または本の執筆時に、 **別ファイルのコードを参照しているが、都度コピペしたくない** と感じている。
 2. 記事または本の執筆時に、 **エディタのモードを切替えるため、コードは別ファイルに記載したい** と感じている。

# 導入手順(npmインストール)

 1. あらかじめ、 **[Zenn CLIを導入](https://zenn.dev/zenn/articles/install-zenn-cli)** します。
 2. Zennのルートディレクトリで、以下のコマンドを実行します。  
**$ `npm install zmce`**
 3. Zennのルートディレクトリに、「submodules」ディレクトリを作成します。  
**$ `mkdir submodules`**

これでzmceの導入は完了です🎉

:::message
submodulesのディレクトリの作成は必須ではなく **推奨(デフォルト)** です。
必要に応じて、config設定で名称を変更して利用することもできます。
:::

# 使い方

**`npx zmce`** コマンドで、 **記事/本 が更新** されます。

:::message
後述する **zmce拡張記法** を用いていなければ、 **コマンドを実行しても 記事/本 は更新されません** が、念のため、Gitのワークツリーに変更が無い状態で実行することを推奨します。
:::

``` shell:変更対象が存在しない / 変更がない場合
$ npx zmce
[zmce] 処理を開始します。
[zmce] 処理を終了します。(変更有 0, 変更無 2, エラー有 0, 対象無 1, スキップ 0)
```

``` shell:変更があった場合
$ npx zmce
[zmce] 処理を開始します。
[articles/zmce-introduction.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 1, 変更無 1, エラー有 0, 対象無 1, スキップ 0)
```

# zmce拡張記法

zmceでは、 **Zennのマークダウン記法で表示されない部分を利用** し、
**zmceコマンドで反復可能な更新[^repeat]** をするための拡張記法を使います。

[^repeat]: コマンドを複数回実施しても結果が変らないこと。参照先の変更をコマンド実行で、いつでも反映できることを担保する。コードブロックを置換する際は、 置換対象コードブロック区切り文字列 が参照先に含まれている場合は、警告を表示し、置換を行わない制御をしている。

## コードブロックに外部ファイルを埋め込む

**[コードブロック記法](https://zenn.dev/zenn/articles/markdown-guide#%E3%82%B3%E3%83%BC%E3%83%89%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF)** の、 **言語/ファイル名指定部分を拡張** します。

コードブロックに埋め込みたいファイルの参照パスを、
**`言語:ファイル名:参照パス` のように:区切りの三番目に記載** します。

:::message
ZennのMarkdown記法では、:区切りの三番目以降は無視されます(使われません)。
:::

ファイル名を指定したくない場合は、
**`言語::参照パス` のように省略して記載** することも可能です。

**参照パスは、用途に応じて以下の3種類の記法** ができます。
（単一の 記事/本 中で **用途に応じて複数の記法を混在** できます。）

 1. **記事/本 からの相対パスで記載する** (`./` もしくは `../`はじまりのパス)
 2. **参照用ディレクトリ(submodules) からの相対パスで記載する** (先頭が、`/` `./` `../`ではじまらないパス)
 3. **絶対パスで記載する** (先頭が、`/` ではじまるパス)

---

### １． 記事/本 からの相対パスで記載する

`./sample/helloWorld.js` や `./memo.txt` や `../package.json` など、
**`./` もしくは `../`はじまりのパス** は、
**記事/本 からの相対パス** として扱います。

:::message
記事/本 のために作成した、 **専用のスニペットやメモを参照** することを想定しています。
:::

:::details 使用例を見る
以下のような構成で 記事/本 を作成します。

``` console:フォルダ構成
|--articles
|  |--sample_article.md
|  |--article_memo.txt
|  |--sample
|  |  |--helloWorld.js
|--books
|  |--sample_book
|  |  |--config.yaml
|  |  |--example1.md
|  |  |--example2.md
|  |  |--fizzbuzz
|  |  |  |--fizzbuzz.js
|--README.md
```

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/relative_path_description/received/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:./article_memo.txt
```

``` js:helloWorld.js:./sample/helloWorld.js
```

``` md:README.md:../README.md
```

`````

````` txt:article_memo.txt:zmce/test/description_case/relative_path_description/received/articles/article_memo.txt
記事メモ
`````

````` js:helloWorld.js:zmce/test/description_case/relative_path_description/received/articles/sample/helloWorld.js
console.log('Hello World!!');
`````

````` yaml:sample_book/config.yaml:zmce/test/description_case/relative_path_description/received/books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

`````

````` md:sample_book/example1.md(コマンド実行前):zmce/test/description_case/relative_path_description/received/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:./fizzbuzz/fizzbuzz.js
```

`````

````` md:sample_book/example2.md(コマンド実行前):zmce/test/description_case/relative_path_description/received/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

``` yaml:config.yaml:./config.yaml
```

``` md:README.md:../../README.md
```

`````

`````  js:fizzbuzz.js:zmce/test/description_case/relative_path_description/received/books/sample_book/fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
`````

````` md:README.md:zmce/test/description_case/relative_path_description/received/README.md
# Docs for zenn.dev
https://zenn.dev/zenn
`````

コマンドを実行すると、 マークダウンファイルが更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/sample_article.md] コードブロックを修正しました。
[books/sample_book/example1.md] コードブロックを修正しました。
[books/sample_book/example2.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 3, 変更無 0, エラー有 0, 対象無 0, スキップ 0)
```

````` md:sample_articles.md(コマンド実行後):zmce/test/description_case/relative_path_description/expected/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:./article_memo.txt
記事メモ
```

``` js:helloWorld.js:./sample/helloWorld.js
console.log('Hello World!!');
```

``` md:README.md:../README.md
# Docs for zenn.dev
https://zenn.dev/zenn
```

`````

````` md:sample_book/example1.md(コマンド実行後):zmce/test/description_case/relative_path_description/expected/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:./fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
```

`````

````` md:sample_book/example2.md(コマンド実行後):zmce/test/description_case/relative_path_description/expected/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

``` yaml:config.yaml:./config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

```

``` md:README.md:../../README.md
# Docs for zenn.dev
https://zenn.dev/zenn
```

`````
:::

---

### ２． 参照用ディレクトリ(submodules) からの相対パスで記載する

`sample/helloWorld.js` や `memo.txt` のように、
**先頭が、`/` `./` `../`ではじまらないパス** は、
**参照用ディレクトリ(submodules) からの相対パス** として扱います。

:::message
記事/本 で Gitのサブモジュール[^submodule]等で、**他のリポジトリなどを参照** することを想定しています。
**参照用ディレクトリ(submodules) 下に参照先を配置** して利用してください。
Gitのサブモジュールの利用を推奨していますが、サブモジュールを使わずに配置しても問題ありません。
:::

:::message
参照用ディレクトリ名は、デフォルト(指定なし)では 「submodules」 です。**zmce.config.ymlに設定すれば、(全体/記事/本/チャプター 毎に) 変更** できます。
:::

[^submodule]: Gitのサブモジュールは、 `git submodule` コマンドで管理できます。例えば、zmce自体のGitリポジトリを参照先リポジトリとして追加するには、 `git submodule add https://github.com/j5c8k6m8/zmce.git submodules/zmce` を実行します。


:::details 使用例を見る
以下のような構成で 記事/本 を作成します。

``` console:フォルダ構成
|--articles
|  |--sample_article.md
|--books
|  |--sample_book
|  |  |--config.yaml
|  |  |--example1.md
|  |  |--example2.md
|--submodules
|  |--article_memo.txt
|  |--sample
|  |  |--helloWorld.js
|  |--fizzbuzz
|  |  |--fizzbuzz.js
|--README.md
```

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/simple_path_description/received/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:article_memo.txt
```

``` js:helloWorld.js:sample/helloWorld.js
```

## 参照用ディレクトリ(submodules) からの相対パスではZennのルートディレクトリは参照できない

`````

````` yaml:sample_book/config.yaml:zmce/test/description_case/simple_path_description/received/books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

`````

````` md:sample_book/example1.md(コマンド実行前):zmce/test/description_case/simple_path_description/received/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:fizzbuzz/fizzbuzz.js
```

`````

````` md:sample_book/example2.md(コマンド実行前):zmce/test/description_case/simple_path_description/received/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

## 参照用ディレクトリ(submodules) からの相対パスではbooksディレクトリ下は参照できない

## 参照用ディレクトリ(submodules) からの相対パスではZennのルートディレクトリは参照できない

`````

````` txt:article_memo.txt:zmce/test/description_case/simple_path_description/received/submodules/article_memo.txt
記事メモ
`````

````` js:helloWorld.js:zmce/test/description_case/simple_path_description/received/submodules/sample/helloWorld.js
console.log('Hello World!!');
`````

`````  js:fizzbuzz.js:zmce/test/description_case/simple_path_description/received/submodules/fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
`````

````` md:README.md:zmce/test/description_case/simple_path_description/received/README.md
# Docs for zenn.dev
https://zenn.dev/zenn
`````

コマンドを実行すると、 マークダウンファイルが更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/sample_article.md] コードブロックを修正しました。
[books/sample_book/example1.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 2, 変更無 0, エラー有 0, 対象無 1, スキップ 0)
```

````` md:sample_articles.md(コマンド実行後):zmce/test/description_case/simple_path_description/expected/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:article_memo.txt
記事メモ
```

``` js:helloWorld.js:sample/helloWorld.js
console.log('Hello World!!');
```

## 参照用ディレクトリ(submodules) からの相対パスではZennのルートディレクトリは参照できない

`````

````` md:sample_book/example1.md(コマンド実行後):zmce/test/description_case/simple_path_description/expected/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
```

`````
:::

---

### ３． 絶対パスで記載する

`/proc/version` や `/etc/os-release` のように、
**先頭が、`/` ではじまるパス** は、**絶対パス** として扱います。

:::message
記事/本 で **OSの情報など実行環境に関する情報を埋め込む** などの利用を想定しています。
:::

:::details 使用例を見る
以下のような、sample_article.md ファイルを作成します。

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/abs_path_description/received/articles/sample_article.md
# Sample Articles(バージョン情報の埋め込み)

``` txt:/proc/version:/proc/version
```

``` txt:/etc/os-release:/etc/os-release
```
`````

筆者のZennの執筆環境で実行すると、以下のように更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/sample_article.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 1, 変更無 0, エラー有 0, 対象無 0, スキップ 0)
```

````` md:sample_article.md(コマンド実行後):zmce/test/description_case/abs_path_description/expected/articles/sample_article.md
# Sample Articles(バージョン情報の埋め込み)

``` txt:/proc/version:/proc/version
Linux version 3.10.0-1062.18.1.el7.x86_64 (mockbuild@kbuilder.bsys.centos.org) (gcc version 4.8.5 20150623 (Red Hat 4.8.5-39) (GCC) ) #1 SMP Tue Mar 17 23:49:17 UTC 2020

```

``` txt:/etc/os-release:/etc/os-release
PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
NAME="Debian GNU/Linux"
VERSION_ID="9"
VERSION="9 (stretch)"
VERSION_CODENAME=stretch
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"

```
`````
:::

---

# zmce.config.yaml によるカスタマイズ

Zennのルートディレクトリに **zmce.config.yaml** を配置することで、
**`npx zmce`コマンドの挙動を変更** できます。
(zmce.config.yamlが存在しない場合、デフォルトの挙動となります。カスタマイズが必要な場合、zmce.config.yamlを手動で作成してください。)
以下の、2点について、 **全体(デフォルト) と、記事/本/チャプター毎** の挙動を設定できます。

 1. **参照用ディレクトリ名** (キー: `relativeRoot`, デフォルト: ` "submodules" `)
 2. **置換対象コードブロック区切り文字列** (キー: `fenceStr`, デフォルト: ` "```" `)
 2. **対象外(skip)設定** (キー: `skip`, デフォルト: ` false `)

コンフィグファイルのフォーマットは以下の通りです。
**各キーは省略可能** です。(省略時はデフォルトを利用します。)

``` yaml
relativeRoot: "submodules"
fenceStr: "```"
articles: # 記事毎の設定(上書き)
    sample_article1: # 記事のスラッグ
        relativeRoot: ""
        fenceStr: "~~~"
    sample_article2: # 記事のスラッグ
        relativeRoot: ""
        fenceStr: "~~~"
        skip: true
books: # 本毎の設定(上書き)
    sample_book1: # 本のスラッグ
        relativeRoot: ""
        fenceStr: "~~~"
    sample_book2: # 本のスラッグ
        relativeRoot: ""
        fenceStr: "~~~"
        skip: true
        chapters: # チャプター毎の設定(本毎の設定を更に上書き)
            sample_chapter1: # チャプターのスラッグ
                relativeRoot: ""
                fenceStr: "~~~"
                skip: false
```

---

### １． 参照用ディレクトリ名のカスタマイズ

キー: `relativeRoot`, デフォルト: ` "submodules" `

参照用のディレクトリ名を設定することができます。

:::message
参照用ディレクトリに、 **`""(空文字)` を設定**すると、
参照用ディレクトリが **Zennのルートディレクトリ** になります。
:::

:::details 使用例を見る
以下のような構成で 記事/本 を作成します。

``` console:フォルダ構成
|--zmce.config.yaml
|--articles
|  |--sample_article.md
|  |--fizzbuzz_article.md
|--books
|  |--sample_book
|  |  |--config.yaml
|  |  |--example1.md
|  |  |--example2.md
|--ref
|  |--article_memo.txt
|  |--sample
|  |  |--helloWorld.js
|  |--fizzbuzz
|  |  |--fizzbuzz.js
|--README.md
```

````` yaml:zmce.config.yaml:zmce/test/description_case/config_relative_root/received/zmce.config.yaml
relativeRoot: "" # Zennのルートディレクトリをデフォルトに
articles:
    # 個別に設定を上書き
    sample_article:
      relativeRoot: "ref"
    fizzbuzz_article:
      relativeRoot: "ref/fizzbuzz"
# デフォルトの設定を使う場合は省略可能
#books:
#    sample_book:
#        relativeRoot: ""

`````

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/config_relative_root/received/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:article_memo.txt
```

``` js:helloWorld.js:sample/helloWorld.js
```

``` md:README.md:../README.md
```

`````

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/config_relative_root/received/articles/fizzbuzz_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# fizzbuzz

``` js:fizzbuzz.js:fizzbuzz.js
```

`````

````` yaml:config.yaml:zmce/test/description_case/config_relative_root/received/books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

`````

````` md:example1.md(コマンド実行前):zmce/test/description_case/config_relative_root/received/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:ref/fizzbuzz/fizzbuzz.js
```

`````

````` md:example2.md(コマンド実行前):zmce/test/description_case/config_relative_root/received/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

``` yaml:config.yaml:books/sample_book/config.yaml
```

``` md:README.md:README.md
```

`````

````` txt:article_memo.txt:zmce/test/description_case/config_relative_root/received/ref/article_memo.txt
記事メモ
`````

````` js:helloWorld.js:zmce/test/description_case/config_relative_root/received/ref/sample/helloWorld.js
console.log('Hello World!!');
`````

````` js:fizzbuzz.js:zmce/test/description_case/config_relative_root/received/ref/fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
`````

````` md:README.md:zmce/test/description_case/config_relative_root/received/README.md
# Docs for zenn.dev
https://zenn.dev/zenn
`````

コマンドを実行すると、 マークダウンファイルが更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/fizzbuzz_article.md] コードブロックを修正しました。
[articles/sample_article.md] コードブロックを修正しました。
[books/sample_book/example1.md] コードブロックを修正しました。
[books/sample_book/example2.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 4, 変更無 0, エラー有 0, 対象無 0, スキップ 0)
```

````` md:sample_article.md(コマンド実行後):zmce/test/description_case/config_relative_root/expected/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` txt:article_memo.txt:article_memo.txt
記事メモ
```

``` js:helloWorld.js:sample/helloWorld.js
console.log('Hello World!!');
```

``` md:README.md:../README.md
# Docs for zenn.dev
https://zenn.dev/zenn
```

`````

````` md:fizzbuzz_article.md(コマンド実行後):zmce/test/description_case/config_relative_root/expected/articles/fizzbuzz_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# fizzbuzz

``` js:fizzbuzz.js:fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
```

`````

````` md:example1.md(コマンド実行後):zmce/test/description_case/config_relative_root/expected/books/sample_book/example1.md
---
title: ""
---

# fizzbuzz

``` js:fizzbuzz.js:ref/fizzbuzz/fizzbuzz.js
for(var i=1;i<101;i++) console.log((i%3?'':'fizz')+(i%5?'':'buzz')||i);
```

`````

````` md:example2.md(コマンド実行後):zmce/test/description_case/config_relative_root/expected/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

``` yaml:config.yaml:books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

```

``` md:README.md:README.md
# Docs for zenn.dev
https://zenn.dev/zenn
```

`````
:::

---

### ２． 置換対象コードブロック区切り文字列のカスタマイズ

キー: `fenceStr`, デフォルト: ` "```" `

Zennのコードブロックは、` ``` `のみではなく、**`~~~` のように、チルダを用いる** ことができます。また、 **` ```` ` のように連続3つ以上** であれば、コードブロックの区切り文字として判定されます。これは、**コードブロック中にコードブロックを記載する時** にも利用できます。[^fence_ref]

[^fence_ref]: **[[小ネタ] 公式に記載されていない、ZennのMarkdown記法 - Zenn](https://zenn.dev/j5c8k6m8/articles/zenn-md-easter-eggs#%E3%82%B3%E3%83%BC%E3%83%89%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF)** 参照

デフォルトでは、**行の先頭からバッククォート3つ以上の文字列** を含むファイルは、参照先として使用できません。　 (警告メッセージが表示されます。)
**コードブロックを含むマークダウンを置換したい場合** は、置換対象コードブロック区切り文字列を `~~~` や、` ```` ` に変更してください。

:::details 使用例を見る
以下のような構成で 記事/本 を作成します。
`README.md` を、`sample_article.md` が参照し、
更に `sample_artcle.md` を `sample_book/example1.md` が参照します。

``` console:フォルダ構成
|--zmce.config.yaml
|--articles
|  |--sample_article.md
|--books
|  |--sample_book
|  |  |--config.yaml
|  |  |--example1.md
|  |  |--example2.md
|--README.md

````` yaml:zmce.config.yaml:zmce/test/description_case/config_fence_str_first/received/zmce.config.yaml
relativeRoot: ""
fenceStr: "~~~"
# デフォルトの設定を使う場合は省略可能
#articles:
#    sample_article:
#      fenceStr: "~~~"
books:
    # 個別に設定を上書き
    sample_book:
        fenceStr: "````"

`````

````` md:sample_article.md(コマンド実行前):zmce/test/description_case/config_fence_str_first/received/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

~~~ md:README.md:README.md
~~~

`````

````` yaml:config.yaml:zmce/test/description_case/config_fence_str_first/received/books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

`````

````` md:example1.md(コマンド実行前):zmce/test/description_case/config_fence_str_first/received/books/sample_book/example1.md
---
title: ""
---

# ref sample article

```` md:sample_article.md:articles/sample_article.md
````

`````

````` md:example2.md(コマンド実行前):zmce/test/description_case/config_fence_str_first/received/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

```` yaml:config.yaml:books/sample_book/config.yaml
````

```` md:README.md:README.md
````

`````

````` md:README.md:zmce/test/description_case/config_fence_str_first/received/README.md
# Docs for zenn.dev
https://zenn.dev/zenn

# CLI install

```
$ npm init --yes # プロジェクトをデフォルト設定で初期化
$ npm install zenn-cli # zenn-cliを導入
```
`````

コマンドを実行すると、 マークダウンファイルが更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/sample_article.md] コードブロックを修正しました。
[books/sample_book/example1.md] コードブロックを修正しました。
[books/sample_book/example2.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 3, 変更無 0, エラー有 0, 対象無 0, スキップ 0)
```

````` md:sample_article.md(コマンド実行後):zmce/test/description_case/config_fence_str_first/expected/articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

~~~ md:README.md:README.md
# Docs for zenn.dev
https://zenn.dev/zenn

# CLI install

```
$ npm init --yes # プロジェクトをデフォルト設定で初期化
$ npm install zenn-cli # zenn-cliを導入
```
~~~

`````

````` md:example1.md(コマンド実行後):zmce/test/description_case/config_fence_str_second/expected/books/sample_book/example1.md
---
title: ""
---

# ref sample article

```` md:sample_article.md:articles/sample_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

~~~ md:README.md:README.md
# Docs for zenn.dev
https://zenn.dev/zenn

# CLI install

```
$ npm init --yes # プロジェクトをデフォルト設定で初期化
$ npm install zenn-cli # zenn-cliを導入
```
~~~

````

`````

````` md:example2.md(コマンド実行後):zmce/test/description_case/config_fence_str_first/expected/books/sample_book/example2.md
---
title: ""
---

# embed config.yaml

```` yaml:config.yaml:books/sample_book/config.yaml
title: ""
summary: ""
topics: []
published: false
price: 0 # 有料の場合200〜5000
# 本に含めるチャプターを順番に並べましょう
chapters:
  - example1
  - example2

````

```` md:README.md:README.md
# Docs for zenn.dev
https://zenn.dev/zenn

# CLI install

```
$ npm init --yes # プロジェクトをデフォルト設定で初期化
$ npm install zenn-cli # zenn-cliを導入
```
````

`````
:::

---


### 3. 対象外(skip)設定

キー: `skip`, デフォルト: ` false `

zmceの拡張記法では、Zennのマークダウン記法で表示されない部分を利用するため、通常、`skip` を明示的に指定する必要はありません。

`skip` を指定せずに、 **zmce拡張記法を用いていない場合、処理結果を「対象無」** としてカウントします。**対象無の場合は、zmceコマンドにおいて、ファイルのパースを行います。**
ファイルサイズが大きい場合や、ファイル数が非常に多い場合は **zmceコマンドの実行時間に影響** します。

`skip` を指定すると、ファイルのパース自体をskipすることで、zmceコマンドの実行時間を短くできます。**skipを指定した場合は、処理結果を「スキップ」としてカウント**します。非常にファイル数が多く、デフォルトの挙動を `skip` にしたい場合は、トップレベルのskipプロパティを `true` に変更することで、記事/本/チャプター毎に明示的に `false` を指定したファイルのみパースする事も実現できます。

また、 `skip` はテンプレートファイル(コピー用のファイル)などで、zmce拡張記法を用いたファイルを更新対象外とする（警告を表示させない）ためにも用いることができるでしょう。

:::details 使用例を見る
以下のような構成で 記事/本 を作成します。

``` console:フォルダ構成
|--zmce.config.yaml
|--articles
|  |--normal_article.md
|  |--skip_article.md
|--books
|  |--skip_book
|  |  |--skip_book_chapter.md
|  |--normal_book
|  |  |--normal_chapter.md
|  |  |--skip_chapter.md
|--submodules
|  |--sample
|  |  |--helloWorld.js
```

````` yaml:zmce.config.yaml:zmce/test/description_case/config_skip/received/zmce.config.yaml
articles:
    skip_article:
        skip: true
books:
    skip_book:
        skip: true
    normal_book:
        chapters:
            skip_chapter:
                skip: true

`````

````` md:normal_article.md(コマンド実行前):zmce/test/description_case/config_skip/received/articles/normal_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:skip_article.md(コマンド実行前):zmce/test/description_case/config_skip/received/articles/skip_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:skip_book_chapter.md(コマンド実行前):zmce/test/description_case/config_skip/received/books/skip_book/skip_book_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:normal_chapter.md(コマンド実行前):zmce/test/description_case/config_skip/received/books/normal_book/normal_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:skip_chapter.md(コマンド実行前):zmce/test/description_case/config_skip/received/books/normal_book/skip_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` js:helloWorld.js:zmce/test/description_case/config_skip/received/submodules/sample/helloWorld.js
console.log('Hello World!!');
`````

コマンドを実行すると、 マークダウンファイルが更新されます。

``` shell:コマンド実行
$ npx zmce
[zmce] 処理を開始します。
[articles/normal_article.md] コードブロックを修正しました。
[books/normal_book/normal_chapter.md] コードブロックを修正しました。
[zmce] 処理を終了します。(変更有 2, 変更無 0, エラー有 0, 対象無 0, スキップ 3)
```

````` md:normal_article.md(コマンド実行後):zmce/test/description_case/config_skip/expected/articles/normal_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
console.log('Hello World!!');
```

`````

````` md:skip_article.md(コマンド実行後):zmce/test/description_case/config_skip/expected/articles/skip_article.md
---
title: ""
emoji: "🙆"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: false
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:skip_book_chapter.md(コマンド実行後):zmce/test/description_case/config_skip/expected/books/skip_book/skip_book_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````

````` md:normal_chapter.md(コマンド実行後):zmce/test/description_case/config_skip/expected/books/normal_book/normal_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
console.log('Hello World!!');
```

`````

````` md:skip_chapter.md(コマンド実行後):zmce/test/description_case/config_skip/expected/books/normal_book/skip_chapter.md
---
title: ""
---

# SAMPLE

``` js:helloWorld.js:sample/helloWorld.js
```

`````
:::


---

# zmceのソースコード

:::details zmceのソースコードを見る
````` ts:zmce.ts:zmce/src/zmce.ts
// ref: https://github.com/zenn-dev/zenn-editor/blob/master/packages/zenn-cli/utils/api/
import fs from "fs-extra";
import { basename, dirname, join } from "path";
import yaml from "js-yaml";
import colors from "colors/safe";

const articlesDirectoryName = "articles";
const booksDirectoryName = "books";
const configFileNameWithoutExtension = "zmce.config";
const defaultRelativeRoot = "submodules";
const defaultFenceStr = "```";
const defaultSkip = false;

type Config = {
  defaultFileConfig: FileConfig;
  articles: { [key: string]: FileConfig };
  books: { [key: string]: FileConfig };
  chapters: { [key: string]: FileConfig };
};

type FileConfig = {
  skip: boolean;
  relativeRoot: string;
  fenceStr: FenceStr;
};

type FenceStr = string;

function isFenceStr(arg: unknown): arg is FenceStr {
  return typeof arg === "string" && /^(````*|~~~~*)$/.test(arg);
}

type ResultCount = {
  change: number;
  noChange: number;
  noTarget: number;
  warn: number;
  skip: number;
};

export function main() {
  consoleInfoSimple(`[zmce] 処理を開始します。`);
  const cwd = process.cwd();
  const config = getConfig(cwd);
  const articleFiles = getArticleFiles(cwd);
  const chapterFiles = getChapterFiles(cwd);
  if (process.exitCode == 1) {
    consoleInfoSimple(
      `[zmce] エラーが発生したため、置換処理を行わずに終了します。`
    );
  } else {
    const rusultCount = {
      change: 0,
      noChange: 0,
      noTarget: 0,
      warn: 0,
      skip: 0,
    };
    articleFilesCodeEmbed(cwd, articleFiles, config, rusultCount);
    chapterFilesCodeEmbed(cwd, chapterFiles, config, rusultCount);
    consoleInfoSimple(
      `[zmce] 処理を終了します。(変更有 ${rusultCount.change}, 変更無 ${rusultCount.noChange}, エラー有 ${rusultCount.warn}, 対象無 ${rusultCount.noTarget}, スキップ ${rusultCount.skip})`
    );
  }
}

function getConfig(basePath: string): Config {
  let fileRaw = null;
  let configFileName = `${configFileNameWithoutExtension}.yaml`;
  try {
    fileRaw = fs.readFileSync(join(basePath, configFileName), "utf8");
  } catch (e) {
    try {
      let configFileName = `${configFileNameWithoutExtension}.yml`;
      fileRaw = fs.readFileSync(join(basePath, configFileName), "utf8");
    } catch (e) {}
  }
  return buildConfig(fileRaw, configFileName);
}

function buildConfig(arg: string | null, configFileName: string): Config {
  let fileConfig: any;
  let relativeRoot = defaultRelativeRoot;
  let fenceStr = defaultFenceStr;
  let skip = defaultSkip;
  const articles: { [key: string]: FileConfig } = {};
  const books: { [key: string]: FileConfig } = {};
  const chapters: { [key: string]: FileConfig } = {};
  if (arg) {
    try {
      fileConfig = yaml.safeLoad(arg);
    } catch (e) {
      consoleError(
        `[${configFileName}] 設定ファイルのフォーマットがyamlファイルとして不正です。`
      );
    }
  }
  if (isHash(fileConfig)) {
    if ("relativeRoot" in fileConfig) {
      if (typeof fileConfig.relativeRoot === "string") {
        relativeRoot = fileConfig.relativeRoot;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルのrelativeRootプロパティには文字列を指定してください。`
        );
      }
    }
    if ("fenceStr" in fileConfig) {
      if (isFenceStr(fileConfig.fenceStr)) {
        fenceStr = fileConfig.fenceStr;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルのfenceStrプロパティには「\`」もしくは「~」の連続した3文字以上の文字列を指定してください。`
        );
      }
    }
    if ("skip" in fileConfig) {
      if (typeof fileConfig.skip === "boolean") {
        skip = fileConfig.skip;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルのskipプロパティにはtrue/falseを指定してください。`
        );
      }
    }
    if ("articles" in fileConfig) {
      if (isHash(fileConfig.articles)) {
        for (let article in fileConfig.articles) {
          articles[article] = buildFileConfig(
            fileConfig.articles[article],
            relativeRoot,
            fenceStr,
            skip,
            `articles.${article}`,
            configFileName
          );
        }
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルのarticlesプロパティは連想配列(ハッシュ)で記載してください。`
        );
      }
    }
    if ("books" in fileConfig) {
      if (isHash(fileConfig.books)) {
        for (let book in fileConfig.books) {
          books[book] = buildFileConfig(
            fileConfig.books[book],
            relativeRoot,
            fenceStr,
            skip,
            `books.${book}`,
            configFileName
          );
          if (isHash(fileConfig.books[book])) {
            if ("chapters" in fileConfig.books[book]) {
              if (isHash(fileConfig.books[book].chapters)) {
                for (let chapter in fileConfig.books[book].chapters) {
                  chapters[`${book}/${chapter}`] = buildFileConfig(
                    fileConfig.books[book].chapters[chapter],
                    books[book].relativeRoot,
                    books[book].fenceStr,
                    books[book].skip,
                    `books.${book}.chapters.${chapter}`,
                    configFileName
                  );
                }
              } else {
                consoleError(
                  `[${configFileName}] 設定ファイルのbooks.${book}.chaptersプロパティは連想配列(ハッシュ)で記載してください。`
                );
              }
            }
          }
        }
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルのbooksプロパティは連想配列(ハッシュ)で記載してください。`
        );
      }
    }
  } else if (fileConfig != null) {
    consoleError(`[${configFileName}] 連想配列(ハッシュ)で記載してください。`);
  }
  return {
    defaultFileConfig: {
      relativeRoot: relativeRoot,
      fenceStr: fenceStr,
      skip: skip,
    },
    articles: articles,
    books: books,
    chapters: chapters,
  };
}

function buildFileConfig(
  arg: any,
  relativeRoot: string,
  fenceStr: FenceStr,
  skip: boolean,
  propertyName: string,
  configFileName: string
): FileConfig {
  if (isHash(arg)) {
    if ("relativeRoot" in arg) {
      if (typeof arg.relativeRoot === "string") {
        relativeRoot = arg.relativeRoot;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルの${propertyName}.relativeRootプロパティには文字列を指定してください。`
        );
      }
    }
    if ("fenceStr" in arg) {
      if (isFenceStr(arg.fenceStr)) {
        fenceStr = arg.fenceStr;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルの${propertyName}.fenceStrプロパティには「\`」もしくは「~」の連続した3文字以上の文字列を指定してください。`
        );
      }
    }
    if ("skip" in arg) {
      if (typeof arg.skip === "boolean") {
        skip = arg.skip;
      } else {
        consoleError(
          `[${configFileName}] 設定ファイルの${propertyName}.skipプロパティにはtrue/falseを指定してください。`
        );
      }
    }
  } else if (arg != null) {
    consoleError(
      `[${configFileName}] 設定ファイルの${propertyName}プロパティは連想配列(ハッシュ)で記載してください。`
    );
  }
  return {
    relativeRoot: relativeRoot,
    fenceStr: fenceStr,
    skip: skip,
  };
}

function isHash(arg: unknown) {
  return typeof arg == "object" && !Array.isArray(arg);
}

function getArticleFiles(basePath: string) {
  let articleFiles: string[] = [];
  try {
    const articleAllFiles = fs.readdirSync(
      join(basePath, articlesDirectoryName)
    );
    articleAllFiles.sort();
    articleFiles = articleAllFiles
      .filter((f) => f.match(/\.md$/))
      .map((f) => join(articlesDirectoryName, f));
  } catch (e) {
    consoleError(
      `プロジェクトルートに${articlesDirectoryName}ディレクトリがありません。`
    );
  }
  return articleFiles;
}

function getChapterFiles(basePath: string) {
  const chapterFiles: string[] = [];
  try {
    let allBookDirs = fs.readdirSync(join(basePath, booksDirectoryName));
    let bookDirs = allBookDirs.filter((f) => {
      try {
        return fs.statSync(join(basePath, booksDirectoryName, f)).isDirectory();
      } catch (e) {
        return false;
      }
    });
    bookDirs.forEach((bookDir) => {
      try {
        const bookChapterAllFiles = fs.readdirSync(
          join(basePath, booksDirectoryName, bookDir)
        );
        bookChapterAllFiles.sort();
        bookChapterAllFiles
          .filter((f) => f.match(/\.md$/))
          .forEach((f) =>
            chapterFiles.push(join(booksDirectoryName, bookDir, f))
          );
      } catch (e) {
        // ファイルの読み込み失敗は無視する
      }
    });
  } catch (e) {
    consoleError(
      `プロジェクトルートに${booksDirectoryName}ディレクトリがありません。`
    );
  }
  return chapterFiles;
}

function articleFilesCodeEmbed(
  basePath: string,
  articleFiles: string[],
  config: Config,
  resultCount: ResultCount
): void {
  articleFiles.forEach((f) => {
    let fileKey = basename(f, ".md");
    codeEmbed(
      basePath,
      f,
      config.articles[fileKey] || config.defaultFileConfig,
      resultCount
    );
  });
}

function chapterFilesCodeEmbed(
  basePath: string,
  chapterFiles: string[],
  config: Config,
  resultCount: ResultCount
): void {
  chapterFiles.forEach((f) => {
    let bookKey = basename(dirname(f));
    let chapterKey = `${bookKey}/${basename(f, ".md")}`;
    codeEmbed(
      basePath,
      f,
      config.chapters[chapterKey] ||
        config.books[bookKey] ||
        config.defaultFileConfig,
      resultCount
    );
  });
}
function codeEmbed(
  basePath: string,
  mdPath: string,
  fileConfig: FileConfig,
  resultCount: ResultCount
): void {
  if (fileConfig.skip) {
    resultCount.skip += 1;
    return;
  }
  let text;
  let targetFlg = false;
  let warnFlg = false;
  try {
    text = fs.readFileSync(join(basePath, mdPath), "utf8");
  } catch (e) {
    return;
  }
  let afterText = text.replace(
    getReplaceCodePattern(fileConfig.fenceStr),
    (
      match,
      beginMark,
      codeType,
      codeName,
      codePath,
      other,
      code,
      afterMark
    ) => {
      targetFlg = true;
      let afterCode;
      codePath = codePath.trim();
      const [codeAbsPath, codeRelativePath] = getCodeAbsRelativePath(
        basePath,
        fileConfig.relativeRoot,
        mdPath,
        codePath
      );
      try {
        afterCode = fs.readFileSync(codeAbsPath, "utf8");
      } catch (e) {
        consoleWarn(`[${mdPath}] 「${codeRelativePath}」ファイルがありません`);
        warnFlg = true;
        return match;
      }
      if (getCheckPattern(fileConfig.fenceStr).test(afterCode)) {
        consoleWarn(
          `[${mdPath}] 「${codeRelativePath}」ファイル内に使用できないパターン(^${fileConfig.fenceStr})が含まれています。`
        );
        warnFlg = true;
        return match;
      }
      return `${beginMark}${codeType}:${codeName}:${codePath}${other}\n${afterCode}\n${afterMark}`;
    }
  );
  if (!targetFlg) {
    resultCount.noTarget += 1;
  } else if (afterText != text) {
    fs.writeFileSync(join(basePath, mdPath), afterText, "utf8");
    consoleInfo(`[${mdPath}] コードブロックを修正しました。`);
    if (warnFlg) {
      resultCount.warn += 1;
    } else {
      resultCount.change += 1;
    }
  } else {
    if (warnFlg) {
      resultCount.warn += 1;
    } else {
      resultCount.noChange += 1;
    }
  }
}

function getReplaceCodePattern(fenceStr: string) {
  return new RegExp(
    `(^${fenceStr})([^${fenceStr[0]}:\n]*):([^:\n]*):([^:\n]+)(.*$)([^]*?)(^${fenceStr}$)`,
    "gm"
  );
}

function getCheckPattern(fenceStr: string) {
  return new RegExp(`^${fenceStr}`, "m");
}

function getCodeAbsRelativePath(
  basePath: string,
  relativeRoot: string,
  mdPath: string,
  codePath: string
): [string, string] {
  if (codePath.startsWith("/")) {
    return [codePath, codePath];
  } else if (/^(\.\/|\.\.\/)/.test(codePath)) {
    let mdDir = dirname(mdPath);
    return [join(basePath, mdDir, codePath), join(mdDir, codePath)];
  } else {
    return [
      join(basePath, relativeRoot, codePath),
      join(relativeRoot, codePath),
    ];
  }
}

function consoleError(msg: string): void {
  process.exitCode = 1;
  console.error(colors.red(msg));
}

function consoleWarn(msg: string): void {
  console.warn(colors.yellow(msg));
}

function consoleInfo(msg: string): void {
  console.info(colors.cyan(msg));
}

function consoleInfoSimple(msg: string): void {
  console.info(msg);
}

`````
:::

# その他

 - 本記事の、使用例や、zmceのソースコードは、zmce拡張記法を用いて埋め込んでいます。  
本記事の原文(テキスト)は、 **[GitHubで参照](https://raw.githubusercontent.com/j5c8k6m8/j5c8k6m8-zenn-contents/master/articles/zmce-introduction.md)** できます。
 - **[コードブロックにファイル名を指定するPR](https://github.com/zenn-dev/zenn-editor/pull/47)** を出しました(merge済)。丁寧に確認して頂き、修正後当日中に(追加で必要な修正入れて頂き)リリースされました。迅速な対応に感謝です。
  - 当初、HTMLのコメント記法を使用した、マクロ置換記法を作ろうと考えました。  
 しかし、Zennのリリース当時はHTMLコメントが使えなかったことと、コードブロックにファイル名を指定するIssuesが既に上がっていたことから、今回のコードブロックのコード名称を拡張する形式にしました。~~先にHTMLのコメント記法に対応が入るとは思わなかった😅~~  
 作成してみて、今の形式もシンプルでわかり易いと考え公開に踏み切りました。
  - 今後、以下のような拡張ができたらいいかと考えています。下記以外でもフィードバックを頂けると励みになります！
    - コマンドの実行結果の埋め込み等への対応
    - HTMLコメント記法を使用した、より制御の細かいマクロ置換記法の作成
    - 初期化コマンド(npx zmce init)対応
    - 特定の行やメソッドのみを参照させるような指定方法の追加
    - コマンド追加(zmce clearなど。ブランチ別でコードを埋め込まない状態で管理するなど)
 - Zennにより、ローカル執筆環境のディレクトリ構成にarticlesやbooksなど、いい意味でConventionが設けられたため、このようなツールが作成しやすい環境ができたと思っています。zmceの思想や使用しているツールは今後、記事を書けたらいいな、と思ってます。
 - 私は心の中で「ジムチェ」と唱えながらコマンドをたたいています。


# link

 - **npm** : https://www.npmjs.com/package/zmce
 - **github** : https://github.com/j5c8k6m8/zmce
