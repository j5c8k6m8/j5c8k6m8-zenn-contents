---
title: "🧪throw後のコード"
---

|||
|:--|:--|
|🔖|[`中断を利用するパターン`](./p_after)|
|👼|[`中断コード削除`](./a_after_stop_delete) [`ネスト修正による中断コードの移動`](./a_after_stop_move)|
|🧟|[`goto文のラベルによるジャンプ`](./z_goto) [`ホイスティング`](./z_hoisting) [`組込モジュールの上書き`](./z_builtin_override) [`組込モジュールの隠蔽`](./z_builtin_hide)|

``` python:🚩 after_throw.py:./projects/python/src/after_throw.py
try:
    raise Exception()
    print("Am I dead?")
except Exception:
    pass

```

|🔧言語|🔩ツール|🚩|note|
|:--|:--|:--|:--|
|[`Python`](#🔧python)|-|`🆗`|-|
||`flake8`|`🆗`|-|
|[`Ruby`](#🔧ruby)|-|`🆗`|-|
||`rubocop`|`⚠`|`UnreachableCode`|
|[`JavaScript`](#🔧javascript)|-|`🆗`|-|
||`eslint`|`⚠`|`no-unreachable`|
|[`Java`](#🔧java)|-|`🚫`|`unreachable_statement`|
|[`Go`](#🔧go)|-|`⚠`|`unreachable_code`|

`throw` は、多くのプログラミング言語で **例外処理のために大域脱出を行う機能** として用意されている。プログラミング言語によっては、`raise` , `panic` , `fail` といった予約語を用いるものもある。

例外処理と大域脱出を区別するプログラミング言語もある。例えば、Rubyにおいては、例外処理は `raise` だが、大域脱出として、`throw` が提供されている。例外処理を伴わない大域脱出も本レシピに含める。しかし、主にループからの大域脱出で使用する `ラベル付break` は [`🧪break後のコード`](./r_after_break) として別レシピとする。

`assert` も例外処理の大域脱出を行う機能であるが、デバッグ用途であり通常の例外処理とは大きく目的が異なるため、 `🧪常に条件がfalseとなるassert後のコード [12/9 公開予定]` として別レシピとする

[`🔖中断を利用するパターン`](./p_after) である。`throw` 後に書かれているコードは、基本的にはデッドコードとなる。

 - `if` 等の条件内を想定していたが、ネストを誤った。 -> [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
 - マージ等で `throw` が重複した。 -> [`👼中断コード削除`](./a_after_stop_delete)
 - `throw` 追加による修正後、不要なコードを消さなかった。 -> [`🛐供養`](./memorial)
 - 削除対象に `goto文のラベル` が記載されている。 -> [`🧟goto文のラベルによるジャンプ`](./z_goto)
 - 削除対象に `ホイスティング対象` が記載されている。  -> [`🧟ホイスティング`](./z_hoisting)
 - `throw` の定義が上書きされている ->  [`🧟組込モジュールの上書き`](./z_builtin_override)
 - `throw` が内側のスコープで定義されている -> [`🧟組込モジュールの隠蔽`](./z_builtin_hide)


# 🔧Python

`🔧Python` -> `🆗実行可`, `🔩flake8` -> `🆗検知無`

予約語は `raise` を用いる。

``` python:🚩 after_throw.py:./projects/python/src/after_throw.py
try:
    raise Exception()
    print("Am I dead?")
except Exception:
    pass

```

``` console
$ # コード実行
$ python src/after_throw.py
$ # flake8
$ flake8 src/after_throw.py
$ 
```

# 🔧Ruby

`🔧Ruby` -> `🆗警告無`,  `🔩rubocop` -> `⚠検知有`

Rubyでは **予約語ではなくKernelモジュールのメソッドとして提供** されている。

そのため、[`🧟組込モジュールの上書き`](./z_builtin_override) や [`🧟組込モジュールの隠蔽`](./z_builtin_hide) が可能。

メソッド名は `raise` で、 `fail` がエイリアスとして使用できる。
（`fail` は `🔩rubocop` に警告される）

例外処理を伴わない、大域脱出用のメソッドとして、 `throw` が別に用意されている。

``` ruby:🚩 after_throw.rb:./projects/ruby/src/after_throw.rb
begin
  raise
  puts 'Am I dead?'
rescue RuntimeError
  # do nothing
end

```

``` console
$ # コード実行
$ ruby src/after_throw.rb
$ # コンパイル(Syntaxチェック&警告確認)
$ ruby -wc src/after_throw.rb 
Syntax OK
$ # rubocop
$ rubocop src/after_throw.rb
/app/ruby/.rubocop.yml: Warning: no department given for Documentation.
Inspecting 1 file
W

Offenses:

src/after_throw.rb:3:3: W: Lint/UnreachableCode: Unreachable code detected.
  puts 'Am I dead?'
  ^^^^^^^^^^^^^^^^^

1 file inspected, 1 offense detected
$ 
```


# 🔧JavaScript

`🔧JavaScript` -> `🆗実行可`, `🔩eslint` -> `⚠検知有`

``` js:🚩 after_throw.js:./projects/javascript/src/after_throw.js
try {
  throw 'Error';
  console.log("Am I dead?");
} catch (e) {
  // empty
}

```

``` console
$ # コード実行
$ node src/after_throw.js 
$ # eslint
$ eslint src/after_throw.js 

/app/javascript/src/after_throw.js
  3:3  error  Unreachable code  no-unreachable

✖ 1 problem (1 error, 0 warnings)

$ 
```

# 🔧Java

`🔧Java` -> `🚫実行不可`

``` java:🚩 AfterThrow.java:./projects/java/src/main/java/AfterThrow.java
public class AfterThrow {
    public static void main(String[] args) {
        try {
            throw new RuntimeException();
            System.out.println("Am I dead?");
        } catch (RuntimeException e) {
        }
    }
}
```

``` console
$ # コード実行
$ java src/main/java/AfterThrow.java 
src/main/java/AfterThrow.java:5: error: unreachable statement
            System.out.println("Am I dead?");
            ^
1 error
error: compilation failed
$ java src/main/java/AfterThrow.java 
$ 
```

# 🔧Go

`🔧Go` -> `⚠検知有`

予約後 `panic` を用いる。

なお、Goでのエラーハンドリングは、 `errorインタフェース` を `戻り値` の最後に付与するのが一般的であり、`panic` はいわゆる `ランタイムエラー` である点には注意が必要。

``` go:🚩 after_throw.go:./projects/golang/src/after_throw.go
package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	panic("Error")
	fmt.Println("Am I dead?")
}

```

``` console
$ # コード実行
$ go run src/after_throw.go 
$ # 静的解析(標準ツール)実行
$ go vet src/after_throw.go 
# command-line-arguments
src/after_throw.go:10:2: unreachable code
$ 
```
