---
title: "付録A. 対象言語と環境"
---

**バージョンの違いによる挙動の違いまでは考慮せず、執筆時点での最新のバージョン** を用いる。例えば、本書でのpythonは、python3を対象として、python2は対象外とする。


# 環境について

静的解析ツール等にも言及するため、projects下に

paiza

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
