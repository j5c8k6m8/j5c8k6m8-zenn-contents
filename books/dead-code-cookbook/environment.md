---
title: "付録A. 対象言語と環境"
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
|バージョン|3.9.0|
|ベースイメージ|python:3.9.0-buster|
|flake8|3.8.4|

``` sh
# Docker build
docker build -t dcc-python $DCC_HOME/projects/python/
# コード実行
docker run --rm -v $DCC_HOME/projects/python:/app/python dcc-python python /app/python/src/after_return.py && echo success
# flake8実行
docker run --rm -v $DCC_HOME/projects/python:/app/python dcc-python flake8 /app/python/src/after_return.py && echo success
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

### ビルドと実行

## Ruby

|||
|:--|:--|
|DOCKERイメージ|ruby:2.7.2-buster|

``` console
docker run --rm -v $ZENN_HOME/books/dead-code-cookbook/src/ruby:/app/ruby ruby:2.7.2-buster ruby /app/ruby/hello_world.rb && echo success
```
