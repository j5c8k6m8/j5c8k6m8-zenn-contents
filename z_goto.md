---
title: "🧟goto文のラベルによるジャンプ"
---

[`🔖中断を利用するパターン`](./p_after) と [`🔖終らない処理を利用するパターン`](./p_forever) で利用できるゾンビ化である。
`goto` 先が同一ブロックという制限がない言語であれば、他のパターンでもゾンビ化が可能。

`goto` 文のラベルをデッドコード中に埋め込みジャンプすることで、ゾンビ化を可能にする。

``` go:goto_zombie.go:./projects/golang/src/zombie/goto_zombie.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Cannot enter arguments.")
		goto L
	}
	return
L:
	fmt.Println("Am I dead?")
}

```

上記は、 `🔧Go` でのサンプルである。`vet(標準ツール)` で `⚠検知可能` である。 

`🔧Go` の場合は **`goto` 先が同一ブロックでなくてはいけないという制限がある**。

``` go:goto_zombie_err1.go:./projects/golang/src/zombie/goto_zombie_err1.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Cannot enter arguments.")
		goto L
	}
	return
	if len(os.Args) > 1 {
	L:
		fmt.Println("Am I dead?")
	}
}

```

上記のプログラムは、 **ネスト内にラベルがあり `🚫実行不可`** である。

``` console
$ # コード実行
$ go run src/zombie/goto_zombie_err1.go 
# command-line-arguments
src/zombie/goto_zombie_err1.go:11:8: goto L jumps into block starting at src/zombie/goto_zombie_err1.go:14:22
$
```

また、 **未使用のラベルも定義できない** 。

``` go:goto_zombie_err2.go:./projects/golang/src/zombie/goto_zombie_err2.go
package main

import "fmt"

func main() {
	goto A
A:
	goto C
B:
	fmt.Println("Am I dead?")
	goto C
C:
	return
}

```

上記のプログラムは、 **未使用なラベルがあるため `🚫実行不可`** である。

``` console
$ # コード実行
$ go run src/zombie/goto_zombie_err2.go 
# command-line-arguments
src/zombie/goto_zombie_err2.go:9:1: label B defined and not used
$
```

上記のような制約も考慮すれば、 **可読性を損なわないため、`分別のあるプログラミング` だと主張** する人もいるだろう。

しかし、以下のプログラムを見て欲しい。

``` go:goto_zombie_multi.go:./projects/golang/src/zombie/goto_zombie_multi.go
package main

import "fmt"

func main() {
	goto A
A:
	goto D
B:
	fmt.Println("Am I dead?")
	goto C
C:
	fmt.Println("Am I dead?")
	goto B
D:
	return
}

```

上記プログラムのラベルBとラベルCが実行されることはないが、`🆗実行可` であり、`vet(標準ツール)` でも `🆗検知無` である。

本書では、**ラベルによるジャンプでしか到達できないコードはデッドコード** として扱い、**`goto` 文のラベルによるジャンプをゾンビ化** として扱う。

`goto` 文を扱える言語は多くない。本書の対象範囲では、 `🔧Go` のみである。

# 言語毎

- `Python (対象外)`
- `Ruby (対象外)`
- `JavaScript (対象外)`
- `Java (対象外)`
- [`Go`](#🔧go)

## 🔧Go

``` go:goto_zombie.go:./projects/golang/src/zombie/goto_zombie.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Cannot enter arguments.")
		goto L
	}
	return
L:
	fmt.Println("Am I dead?")
}

```
