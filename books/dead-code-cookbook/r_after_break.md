---
title: "🧪break後のコード"
---


|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting)|

``` python:🚩 after_break.py:./projects/python/src/after_break.py
while True:
    break
    print("Am I dead?")

```

``` java:🏁 AfterBreakLabel.java:./projects/java/src/main/java/AfterBreakLabel.java
public class AfterBreakLabel {
    public static void main(String[] args) {
        l:{
            while(true) {
                while(true) {
                    break l;
                }
                System.out.println("Am I dead?");
            }
        }
    }
}
```

|🔧言語|🔩ツール|🚩|note|🏁|note|
|:--|:--|:--|:--|:--|:--|
|[`Python`](#🔧python)|-|`🆗`|-|-|対象外|
||`flake8`|`🆗`|-|-|対象外|
|[`Ruby`](#🔧ruby)|-|`⚠`|`warning:_statement_not_reached`|-|対象外|
||`rubocop`|`⚠`|`UnreachableCode`|-|対象外|
|[`JavaScript`](#🔧javascript)|-|`🆗`|-|`🆗`|-|
||`eslint`|`⚠`|`no-unreachable`|`⚠`|`no-unreachable`|
|[`Java`](#🔧java)|-|`🚫`|`unreachable_statement`|`🚫`|`unreachable_statement`|
|[`Go`](#🔧go)|-|`⚠`|`unreachable_code`|`⚠`|`unreachable_code`|

`break` は、多くのプログラミング言語で、 **ループ文から抜け出すために使用** する。プログラミング言語によっては、 `switch文` から抜け出すためにも使う。

`continue` などのループ文で用いられるジャンプ文を利用した場合も、本レシピに含める。

また、`label付break` も `🏁追加サンプル` として本レシピに含める。`label付break` では、 `break` の直後ではなく、ネストしたループにおいて、 `break` よりも上位のループの後続でデッドコードとなるパターンを示す。

[`🔖中断を利用するパターン`](./p_after) である。`break` 後に書かれているコードは、基本的にはデッドコードとなる。デバッグのために使われることもある。

 - デバッグのために残した `break` が残ってしまった。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `break` 文が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `break` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)

# 🔧Python

`🚩` : `🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`
`🏁` : 対象外

`switch文` は提供していない。

``` python:🚩 after_break.py:./projects/python/src/after_break.py
while True:
    break
    print("Am I dead?")

```

``` console
$ # コード実行
$ python src/after_break.py 
$ # flake8
$ flake8 src/after_break.py 
$ 
```

# 🔧Ruby

`🚩` : `🔧Ruby` -> `⚠警告有`,  `🔩rubocop` -> `⚠検知有`
`🏁` : 対象外

Rubyでは、`continue` ではなく、 `next` を用いる。また、 `break` , `next` の他に、 `redo` `retry` がジャンプできる予約語として用意されている。

`switch文` である `case` では `break` は使わない。

``` ruby:🚩 after_break.rb:./projects/ruby/src/after_break.rb
while true
  break
  puts 'Am I dead?'
end

```

``` console
$ # コード実行
$ ruby src/after_break.rb 
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_break.rb 
src/after_break.rb:3: warning: statement not reached
Syntax OK
$ # rubocop
$ rubocop src/after_break.rb 
Inspecting 1 file
W

Offenses:

src/after_break.rb:1:1: W: Lint/UnreachableLoop: This loop will have at most one iteration.
while true ...
^^^^^^^^^^
src/after_break.rb:1:1: C: Style/InfiniteLoop: Use Kernel#loop for infinite loops.
while true
^^^^^
src/after_break.rb:3:3: W: Lint/UnreachableCode: Unreachable code detected.
  puts 'Am I dead?'
  ^^^^^^^^^^^^^^^^^

1 file inspected, 3 offenses detected, 1 offense auto-correctable
$ 
```

# 🔧JavaScript

`🚩` : `🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `⚠検知有`
`🏁` : `🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `⚠検知有`

`switch` 文中で `break` が使用できる。

``` js:🚩 after_break.js:./projects/javascript/src/after_break.js
while(true) {
  break;
  console.log("Am I dead?");
}

```

``` console
$ # コード実行
$ node src/after_break.js 
$ # eslint
$ eslint src/after_break.js 

/app/javascript/src/after_break.js
  1:7  error  Unexpected constant condition  no-constant-condition
  3:3  error  Unreachable code               no-unreachable

✖ 2 problems (2 errors, 0 warnings)

$ 
```

``` js:🏁 after_break_label.js:./projects/javascript/src/after_break_label.js
l: while (true) {
  while(true) {
    break l;
  }
  console.log("Am I dead?");
}

```

``` console
$ # コード実行
$ node src/after_break_label.js 
$ # eslint
$ eslint src/after_break_label.js 

/app/javascript/src/after_break_label.js
  1:11  error  Unexpected constant condition  no-constant-condition
  2:9   error  Unexpected constant condition  no-constant-condition
  5:3   error  Unreachable code               no-unreachable

✖ 3 problems (3 errors, 0 warnings)

$ 
```

# 🔧Java

`🚩` : `🔧Java` -> `🚫実行不可`
`🏁` : `🔧Java` -> `🚫実行不可`

`switch` 文中で `break` が使用できる。

``` java:🚩 AfterBreak.java:./projects/java/src/main/java/AfterBreak.java
public class AfterBreak {
    public static void main(String[] args) {
        while(true) {
            break;
            System.out.println("Am I dead?");
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterBreak.java 
src/main/java/AfterBreak.java:5: error: unreachable statement
            System.out.println("Am I dead?");
            ^
1 error
error: compilation failed
$ 
```

``` java:🏁 AfterBreakLabel.java:./projects/java/src/main/java/AfterBreakLabel.java
public class AfterBreakLabel {
    public static void main(String[] args) {
        l:{
            while(true) {
                while(true) {
                    break l;
                }
                System.out.println("Am I dead?");
            }
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterBreakLabel.java 
src/main/java/AfterBreakLabel.java:8: error: unreachable statement
                System.out.println("Am I dead?");
                ^
1 error
error: compilation failed
$ 
```


# 🔧Go

`🚩` : `🔧Go` -> `⚠検知有`
`🏁` : `🔧Go` -> `⚠検知有`

``` go:🚩 after_break.go:./projects/golang/src/after_break.go
package main

import "fmt"

func main() {
	for {
		break
		fmt.Println("Am I dead?")
	}
}

```

``` console
$ # コード実行
$ go run src/after_break.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_break.go 
# command-line-arguments
src/after_break.go:8:3: unreachable code
$ 
```

``` go:🏁 after_break_label.go:./projects/golang/src/after_break_label.go
package main

import "fmt"

func main() {
L:
	for {
		for {
			break L
		}
		fmt.Println("Am I dead?")
	}
}

```

``` console
$ # コード実行
$ go run src/after_break_label.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_break_label.go 
# command-line-arguments
src/after_break_label.go:11:3: unreachable code
$ 
```