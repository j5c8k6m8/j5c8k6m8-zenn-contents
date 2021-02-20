---
title: "👼ネスト修正による中断コードの移動"
---

[`🔖中断を利用するパターン`](./p_after) と [`🔖終らない処理を利用するパターン`](./p_forever) で利用できるレイズである。

構文木において、 `return` などの中断を伴う処理が、前要素ではなくなるように、インデントや中括弧を追加することで、後続の処理を復活させる。 

``` diff python:after_return_with_if.py:exec diff -U 100 books/dead-code-cookbook/projects/python/src/angel/before/after_stop_move.py books/dead-code-cookbook/projects/python/src/angel/after/after_stop_move.py | tail -n +4
 import sys
 
 
 def f():
     if len(sys.argv) > 1:
         print("Cannot enter arguments.")
-    return
+        return
     print("Am I dead?")
 
 
 f()

```

以下では、上記の様に前段に分岐処理がある [`🧪return後のコード`](./r_after_return) に対する修正のサンプルを示す。


# 言語毎

- [`Python`](#🔧python)
- [`Ruby`](#🔧ruby)
- [`JavaScript`](#🔧javascript)
- [`Java`](#🔧java)
- [`Go`](#🔧go)


## 🔧Python

``` diff python:after_return_with_if.py:exec diff -U 100 books/dead-code-cookbook/projects/python/src/angel/before/after_stop_move.py books/dead-code-cookbook/projects/python/src/angel/after/after_stop_move.py | tail -n +4
 import sys
 
 
 def f():
     if len(sys.argv) > 1:
         print("Cannot enter arguments.")
-    return
+        return
     print("Am I dead?")
 
 
 f()

```


## 🔧Ruby

``` diff ruby:after_return_with_if.rb:exec diff -U 100 books/dead-code-cookbook/projects/ruby/src/angel/before/after_stop_move.rb books/dead-code-cookbook/projects/ruby/src/angel/after/after_stop_move.rb | tail -n +4
 if ARGV.size.positive?
   puts('Cannot enter arguments.')
+  return
 end
-return
 puts 'Am I dead?'

```


## 🔧JavaScript

``` diff js:after_return_with_if.js:exec diff -U 100 books/dead-code-cookbook/projects/javascript/src/angel/before/after_stop_move.js books/dead-code-cookbook/projects/javascript/src/angel/after/after_stop_move.js | tail -n +4
 (() => {
     if (process.argv.length > 2) {
         console.log("Cannot enter arguments.");
+        return
     }
-    return
     console.log("Am I dead?");
 })()
\ No newline at end of file

```


## 🔧Java

``` diff java:AfterReturnWithIf.java:exec diff -U 100 books/dead-code-cookbook/projects/java/src/main/java/angel/before/after_stop_move/AfterReturnWithIf.java books/dead-code-cookbook/projects/java/src/main/java/angel/after/after_stop_move/AfterReturnWithIf.java | tail -n +4
 public class AfterReturnWithIf {
     public static void main(String[] args) {
         if (args.length > 0) {
             System.out.println("Cannot enter arguments.");
+            return;
         }
-        return;
         System.out.println("Am I dead?");
     }
 }
\ No newline at end of file

```


## 🔧Go

``` diff go:after_return_with_if.go:exec diff -U 100 books/dead-code-cookbook/projects/golang/src/angel/before/after_stop_move.go books/dead-code-cookbook/projects/golang/src/angel/after/after_stop_move.go | tail -n +4
 package main
 
 import (
 	"fmt"
 	"os"
 )
 
 func main() {
 	if len(os.Args) > 1 {
 		fmt.Println("Cannot enter arguments.")
+		return
 	}
-	return
 	fmt.Println("Am I dead?")
 }

```
