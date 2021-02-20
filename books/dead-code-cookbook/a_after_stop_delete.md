---
title: "👼中断コード削除"
---

[`🔖中断を利用するパターン`](./p_after) と [`🔖終らない処理を利用するパターン`](./p_forever) で利用できるレイズである。

`return` などの中断を伴う処理削除することで、後続の処理を復活させる。デバッグのために `return` を入れて消し忘れた場合や、マージ等で `return` が重複したケースが考えられる。

``` diff ruby:after_return.rb:exec diff -U 100 books/dead-code-cookbook/projects/ruby/src/after_return.rb books/dead-code-cookbook/projects/ruby/src/angel/after/after_stop_delete.rb | tail -n +4
-return
 puts 'Am I dead?'

```

以下では、[`🧪return後のコード`](./r_after_return) に対する修正のサンプルを示す。

# 言語毎

- [`Python`](#🔧python)
- [`Ruby`](#🔧ruby)
- [`JavaScript`](#🔧javascript)
- [`Java`](#🔧java)
- [`Go`](#🔧go)


## 🔧Python

``` diff python:after_return.py:exec diff -U 100 books/dead-code-cookbook/projects/python/src/after_return.py books/dead-code-cookbook/projects/python/src/angel/after/after_stop_delete.py | tail -n +4
 def f():
-    return
     print("Am I dead?")
 
 
 f()

```


## 🔧Ruby

``` diff ruby:after_return.rb:exec diff -U 100 books/dead-code-cookbook/projects/ruby/src/after_return.rb books/dead-code-cookbook/projects/ruby/src/angel/after/after_stop_delete.rb | tail -n +4
-return
 puts 'Am I dead?'

```


## 🔧JavaScript

``` diff javascript:after_return.js:exec diff -U 100 books/dead-code-cookbook/projects/javascript/src/after_return.js books/dead-code-cookbook/projects/javascript/src/angel/after/after_stop_delete.js | tail -n +4
 (() => {
-  return;
   console.log("Am I dead?");
 })();
\ No newline at end of file

```


## 🔧Java

``` diff java:AfterReturn.java:exec diff -U 100 books/dead-code-cookbook/projects/java/src/main/java/AfterReturn.java books/dead-code-cookbook/projects/java/src/main/java/angel/after/after_stop_delete/AfterReturn.java | tail -n +4
 public class AfterReturn {
     public static void main(String[] args) {
-        return;
         System.out.println("Am I dead?");
     }
 }
\ No newline at end of file

```


## 🔧Go

``` diff go:after_return.go:exec diff -U 100 books/dead-code-cookbook/projects/golang/src/after_return.go books/dead-code-cookbook/projects/golang/src/angel/after/after_stop_delete.go | tail -n +4
 package main
 
 import "fmt"
 
 func main() {
-	return
 	fmt.Println("Am I dead?")
 }

```
