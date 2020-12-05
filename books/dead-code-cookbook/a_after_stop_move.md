---
title: "👼ネスト修正による中断コードの移動"
---

[`🔖中断を利用するパターン`](./p_after) と [`🔖終らない処理を利用するパターン`](./p_forever) で利用できるレイズである。

構文木において、 `return` などの中断を伴う処理が、前要素ではなくなるように、インデントや中括弧を追加することで、後続の処理を復活させる。 

``` diff:after_return_with_if.py:./projects/python/diff/angel/after_stop_move.diff
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

``` diff:after_return_with_if.py:./projects/python/diff/angel/after_stop_move.diff
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

``` diff:after_return_with_if.rb:./projects/ruby/diff/angel/after_stop_move.diff
 if ARGV.size.positive?
   puts('Cannot enter arguments.')
+  return
 end
-return
 puts 'Am I dead?'

```


## 🔧JavaScript

``` diff:after_return_with_if.js:./projects/javascript/diff/angel/after_stop_move.diff
 (() => {
     if (process.argv.length > 2) {
         console.log("Cannot enter arguments.");
+        return
     }
-    return
     console.log("Am I dead?");
 })()
```


## 🔧Java

``` diff:AfterReturnWithIf.java:./projects/java/diff/angel/after_stop_move.diff
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
```


## 🔧Go

``` diff:after_return_with_if.go:./projects/golang/diff/angel/after_stop_move.diff
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
