---
title: "🧰対象言語と環境"
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

# 言語毎

- [`Python`](#🔧Python)
- [`Ruby`](#🔧Ruby)
- [`JavaScript`](#🔧JavaScript)
- [`Java`](#🔧Java)
- [`Go`](#🔧Go)

## 🔧Python

|||
|:--|:--|
|バージョン|Python 3.9.0|
|ベースイメージ|python:3.9.0-buster|
|flake8|3.8.4 (mccabe: 0.6.1, pycodestyle: 2.6.0, pyflakes: 2.2.0)|

``` sh
# Docker build
docker build -t dcc-python $DCC_HOME/projects/python/
# コンテナ起動
docker run -it --rm -v $DCC_HOME/projects/python:/app/python -w /app/python dcc-python bash
```

``` sh:コンテナ内で実行
# コード実行
python src/after_return.py
# flake8
flake8 src/after_return.py
```


## 🔧Ruby

|||
|:--|:--|
|バージョン|ruby 2.7.2p137 (2020-10-01 revision 5445e04352) [x86_64-linux]|
|ベースイメージ|ruby:2.7.2-buster|
|rubocop|1.3.1|

```  sh
# Docker build
docker build -t dcc-ruby $DCC_HOME/projects/ruby/
# コンテナ起動
docker run -it --rm -v $DCC_HOME/projects/ruby:/app/ruby -w /app/ruby dcc-ruby bash
```

``` sh
# コード実行
ruby src/after_return.rb
# コンパイル(Syntaxチェック&警告確認)
ruby -wc src/after_return.rb
# rubocop
rubocop src/after_return.rb
```

## 🔧JavaScript

|||
|:--|:--|
|バージョン|ruby 2.7.2p137 (2020-10-01 revision 5445e04352) [x86_64-linux]|
|ベースイメージ|node:15.2.1-buster|
|rubocop|1.3.1|

```  sh
# Docker build
docker build -t dcc-javascript $DCC_HOME/projects/javascript/
# コンテナ起動
docker run -it --rm -v $DCC_HOME/projects/javascript:/app/javascript -w /app/javascript dcc-javascript bash
```

``` sh
# コード実行
node src/after_return.js
# eslint
eslint src/after_return.js
```


## 🔧Java

|||
|:--|:--|
|バージョン|openjdk 15.0.1 2020-10-20 OpenJDK Runtime Environment (build 15.0.1+9-18) OpenJDK 64-Bit Server VM (build 15.0.1+9-18, mixed mode, sharing)|
|ベースイメージ|openjdk:15.0.1|

```  sh
# Docker build
docker build -t dcc-java $DCC_HOME/projects/java/
# コンテナ起動
docker run -it --rm -v $DCC_HOME/projects/java:/app/java -w /app/java dcc-java bash
```

``` sh
# コード実行
java src/main/java/AfterReturn.java
```

## 🔧Go

|||
|:--|:--|
|バージョン|go version go1.15.4 linux/amd64|
|ベースイメージ|golang:1.15.4-buster|


``` console
# Docker build
docker build -t dcc-golang $DCC_HOME/projects/golang/
# コンテナ起動
docker run -it --rm -v $DCC_HOME/projects/golang:/app/golang -w /app/golang dcc-golang bash
```

```
# コード実行
go run src/after_return.go
# 静的解析(標準ツール)実行
go vet src/after_return.go
```

