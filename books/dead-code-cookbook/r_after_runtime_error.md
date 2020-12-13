---
title: "🧪実行時エラーとなる処理後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`実行時エラー修正`](./a_runtime_error_fix)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|


``` python:🚩 after_runtime_error.py:./projects/python/src/after_runtime_error.py
try:
    None.a
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

特定の関数を呼び出さなくても、実行時エラーを発生させることができる。実行時エラーについては、大きく２種類に分けられるだろう。１つは実行時のリソースに関連して発生するものであり、もう１つはソースコードの不具合によって発生するものである。本レシピでは後者のソースコードの不具合によって常に実行時エラーを発生させることでデッドコードを作成する。利用する実行時エラーとしては `NullPointerException` や `NoMethodError` が代表的な例だろう。

プログラミング言語によって、発生させることが出来る実行時エラーは異なるだろう。本レシピではその種類に言及することはせずに、代表的な実行時エラーの例を一つあげる。

::: message
今後、プログラミング言語毎の実行時エラーの比較をデッドコードツールとしてまとめる予定
:::

 - 不具合により実行時エラーが発生する -> [`👼実行時エラー修正`]((./a_runtime_error_fix))


# 🔧Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`


``` python:🚩 after_runtime_error.py:./projects/python/src/after_runtime_error.py
try:
    None.a
    print("Am I dead?")
except Exception:
    pass

```

``` console
$ # コード実行
$ python src/after_runtime_error.py 
$ # flake8
$ flake8 src/after_runtime_error.py 
$ 
```

# 🔧Ruby

``` ruby:🚩 after_runtime_error.rb:./projects/ruby/src/after_runtime_error.rb
begin
  nil.a
  puts 'Am I dead?'
rescue Exception
  # do nothing
end

```

``` console
$ # コード実行
$ ruby src/after_runtime_error.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_runtime_error.rb 
Syntax OK
$ # rubocop
$ rubocop src/after_runtime_error.rb 
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
W

Offenses:

src/after_runtime_error.rb:4:1: W: Lint/RescueException: Avoid rescuing the Exception class. Perhaps you meant to rescue StandardError?
rescue Exception
^^^^^^^^^^^^^^^^

1 file inspected, 1 offense detected
$ 
```

# 🔧JavaScript

``` js:🚩 after_runtime_error.js:./projects/javascript/src/after_runtime_error.js
try {
  null.a;
  console.log("Am I dead?");
} catch (e) {
  // empty
}

```

``` console
$ # コード実行
$ node src/after_runtime_error.js 
$ # eslint
$ eslint src/after_runtime_error.js 
$ 
```

# 🔧Java

``` java:🚩 AfterRuntimeError.java:./projects/java/src/main/java/AfterRuntimeError.java
public class AfterRuntimeError {
    public static void main(String[] args) {
        try {
            Integer a = null, b = null;
            int x = a + b;
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterRuntimeError.java 
$ 
```

# 🔧Go

``` go:🚩 after_runtime_error.go:./projects/golang/src/after_runtime_error.go
package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	a := 0
	b := 1
	fmt.Println(b / a)
	fmt.Println("Am I dead?")
}

```

``` console
$ # コード実行
$ go run src/after_runtime_error.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_runtime_error.go 
$ 
```
