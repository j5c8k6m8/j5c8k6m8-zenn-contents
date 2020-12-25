---
title: "🧪throwする関数呼出後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move) [`関数修正`](./a_func_fix)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` python:🚩 after_func_throw.py:./projects/python/src/after_func_throw.py
def f():
    raise Exception()


try:
    f()
    print("Am I dead?")
except Exception:
    pass

```

|🔧言語|🔩ツール|🚩|note|
|:--|:--|:--|:--|
|[`Python`](#🔧python)|-|`🆗`|-|
||`flake8`|`🆗`|-|
|[`Ruby`](#🔧ruby)|-|`🆗`|-|
||`rubocop`|`🆗`|-|
|[`JavaScript`](#🔧javascript)|-|`🆗`|-|
||`eslint`|`🆗`|-|
|[`Java`](#🔧java)|-|`🆗`|-|
|[`Go`](#🔧go)|-|`🆗`|-|

中断させる処理は、予約語や、組み込み関数のみではなく、自身で作成することもできる。

中断させる関数を作る最も一般的な方法は、 `throw` を利用することだろう。 `throw` を利用することで、 [`🧪throw後のコード`](./r_after_throw) のように、大域脱出を行うことが出来る。

デッドコードを作るために、 常にthrowする関数である必要はない。 **与えられ得る引数(条件)に対して常にthrowする関数** であれば後続のコードはデッドコードとなる。

::: message
本書では、常にthrowする関数のサンプルを掲載する。しかし、常にthrowする関数を検知可能なプログラミング言語を掲載する時に、本レシピ内に 🏁 を付けて条件付のthrowのサンプルを掲載する予定。
:::

エラー処理用の関数などであれば、常に `throw` する関数を使うこともある。エラー処理用の関数で無ければ、常に `throw` しないように関数呼出に関連した部分を見直すことになるだろう。

 - 呼出対象の関数、関数の引数、呼出先の関数の処理が誤っている。 -> [`👼関数修正`](./a_func_fix)
 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で 関数呼出 が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - 関数呼出 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)


# 🔧Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`

``` python:🚩 after_func_throw.py:./projects/python/src/after_func_throw.py
def f():
    raise Exception()


try:
    f()
    print("Am I dead?")
except Exception:
    pass

```

``` console
$ # コード実行
$ python src/after_func_throw.py 
$ # flake8
$ flake8 src/after_func_throw.py 
$ 
```

# 🔧Ruby

`🔧Ruby` -> `🆗警告無`, `rubocop` -> `🆗検知無`

``` ruby:🚩 after_func_throw.rb:./projects/ruby/src/after_func_throw.rb
def f
  raise
end

begin
  f
  puts 'Am I dead?'
rescue RuntimeError
  # do nothing
end

```

``` console
$ # コード実行
$ ruby src/after_func_throw.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_func_throw.rb 
Syntax OK
$ # rubocop
$ rubocop src/after_func_throw.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
.

1 file inspected, no offenses detected
$ 
```

# 🔧JavaScript

`🔧JavaScript` -> `🆗実行可`, `eslint` -> `🆗検知無`

``` js:🚩 after_func_throw.js:./projects/javascript/src/after_func_throw.js
const f = () => {
  throw 'Error';
};

try {
  f();
  console.log("Am I dead?");
} catch (e) {
  // empty
}
  
```

``` console
$ # コード実行
$ node src/after_func_throw.js 
$ # eslint
$ eslint src/after_func_throw.js 
$ 
```

# 🔧Java

`🔧Java` -> `🆗実行可`

``` java:🚩 AfterFuncThrow.java:./projects/java/src/main/java/AfterFuncThrow.java
public class AfterFuncThrow {
    public static void main(String[] args) {
        try {
            f();
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }

    private static void f() {
        throw new RuntimeException();
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterFuncThrow.java 
$ 
```

# 🔧Go

`🔧go` -> `🆗検知無`

``` go:🚩 after_func_throw.go:./projects/golang/src/after_func_throw.go
package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	f()
	fmt.Println("Am I dead?")
}

func f() {
	panic("Error")
}

```

``` console
$ # コード実行
$ go run src/after_func_throw.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_func_throw.go 
$ 
```
