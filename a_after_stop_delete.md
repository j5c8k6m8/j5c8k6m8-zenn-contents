---
title: "👼中断コード削除"
---

[`🔖中断を利用するパターン`](./p_after) と [`🔖終らない処理を利用するパターン`](./p_forever) で利用できるレイズである。

`return` などの中断を伴う処理削除することで、後続の処理を復活させる。デバッグのために `return` を入れて消し忘れた場合や、マージ等で `return` が重複したケースが考えられる。

``` diff:after_return.rb:./projects/ruby/diff/angel/after_stop_delete.diff
-return
 puts 'Am I dead?'

```

以下では、[`🧪return後のコード`](./r_after_return) に対する修正のサンプルを示す。

# 言語毎

- [`Python`](#🔧Python)
- [`Ruby`](#🔧Ruby)
- [`JavaScript`](#🔧JavaScript)
- [`Java`](#🔧Java)
- [`Go`](#🔧Go)


## 🔧Python

``` diff:after_return.py:./projects/python/diff/angel/after_stop_delete.diff
 def f():
-    return
     print("Am I dead?")
 
 
 f()

```


## 🔧Ruby

``` diff:after_return.rb:./projects/ruby/diff/angel/after_stop_delete.diff
-return
 puts 'Am I dead?'

```


## 🔧JavaScript

``` diff:after_return.js:./projects/javascript/diff/angel/after_stop_delete.diff
 (() => {
-    return
     console.log("Am I dead?")
 })()
```


## 🔧Java

``` diff:AfterReturn.java:./projects/java/diff/angel/after_stop_delete.diff
 public class AfterReturn {
     public static void main(String[] args) {
-        return;
         System.out.println("Am I dead?");
     }
 }
```


## 🔧Go

``` diff:after_return.go:./projects/golang/diff/angel/after_stop_delete.diff
 package main
 
 import "fmt"
 
 func main() {
-	return
 	fmt.Println("Am I dead?")
 }

```
