---
title: "🧰付録A. 対象言語と環境"
---

# 前提


本書のサンプルコードは、 **本書のリポジトリのprojectsフォルダ下にプログラミング言語毎に配置** する。

各プログラミング言語の実行環境は原則として、 **プログラミング言語フォルダ下のDOCKERファイルで構築したコンテナ** とする。静的解析ツール等の環境は、ツールのインストールをDOCKERファイルで記述する。

::: message
作成するDOCKERイメージ内にはサンプルコードは含めず、**データ・ボリュームとしてホスト上のディレクトリをマウント**[^1] して利用することを想定する。
:::

[^1]: https://docs.docker.jp/engine/userguide/dockervolumes.html#id5

**バージョンの違いによる挙動の違いまでは考慮せず、執筆時点での最新のバージョン** を用いる。例えば、本書でのpythonは、python3を対象として、python2は対象外とする。

::: message
本付録のコマンド内では本書のリポジトリのパスを `DCC_HOME` の環境変数として記載する。
**本コマンドをコピペして使用する場合は、 `DCC_HOME` に本書のリポジトリのパスを設定** すること。
:::

# プログラミング言語一覧

- Python
- JavaScript(Node.js)
- Ruby
- Go
- Java
- TypeScript
- Rust

# プログラミング言語毎詳細

## Python

|||
|:--|:--|
|バージョン|Python 3.9.0|
|ベースイメージ|python:3.9.0-buster|
|flake8|3.8.4 (mccabe: 0.6.1, pycodestyle: 2.6.0, pyflakes: 2.2.0)|

``` sh
# Docker build
docker build -t dcc-python $DCC_HOME/projects/python/
# コード実行
docker run --rm -v $DCC_HOME/projects/python:/app/python dcc-python python /app/python/src/after_return.py && echo success
# flake8実行
docker run --rm -v $DCC_HOME/projects/python:/app/python dcc-python flake8 /app/python/src/after_return.py
```

## Go

|||
|:--|:--|
|DOCKERイメージ|1.15.4-buster|

``` console
docker run --rm -v $ZENN_HOME/books/dead-code-cookbook/src/golang:/app/golang golang:1.15.4-buster go run /app/golang/hello_world.go && echo success
```

## Java

|||
|:--|:--|
|DOCKERイメージ|openjdk:15.0.1|

## Ruby

|||
|:--|:--|
|バージョン|ruby 2.7.2p137 (2020-10-01 revision 5445e04352) [x86_64-linux]|
|ベースイメージ|ruby:2.7.2-buster|
|rubocop|1.3.1|

```  sh
# Docker build
docker build -t dcc-ruby $DCC_HOME/projects/ruby/
# コード実行
docker run --rm -v $DCC_HOME/projects/ruby:/app/ruby dcc-ruby ruby /app/ruby/after_return.rb && echo success
# コード実行(コンパイル(Syntaxチェック&警告確認)
docker run --rm -v $DCC_HOME/projects/ruby:/app/ruby dcc-ruby ruby -wc /app/ruby/after_return.rb
# rubocop
docker run --rm -v $DCC_HOME/projects/ruby:/app/ruby dcc-ruby rubocop /app/ruby/after_return.rb
```

## JavaScript

|||
|:--|:--|
|バージョン|ruby 2.7.2p137 (2020-10-01 revision 5445e04352) [x86_64-linux]|
|ベースイメージ|node:15.2.1-buster|
|rubocop|1.3.1|

```  sh
# Docker build
docker build -t dcc-javascript $DCC_HOME/projects/javascript/
# コード実行
docker run --rm -v $DCC_HOME/projects/javascript:/app/javascript dcc-javascript node /app/javascript/src/after_return.js && echo success
# eslint
docker run --rm -v $DCC_HOME/projects/javascript:/app/javascript dcc-javascript eslint /app/javascript/src/after_return.js
```
