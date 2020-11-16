---
title: "付録A. 対象言語と環境"
---

# 前提


本書のサンプルコードは、 **本書のリポジトリのprojectsフォルダ下にプログラミング言語毎に配置** する。

各プログラミング言語の実行環境は原則として、 **プログラミング言語フォルダ下のDOCKERファイルで構築したコンテナ** とする。静的解析ツール等の環境は、ツールのインストールをDOCKERファイルで記述する。

**バージョンの違いによる挙動の違いまでは考慮せず、執筆時点での最新のバージョン** を用いる。例えば、本書でのpythonは、python3を対象として、python2は対象外とする。


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
