---
title: "👼ネスト修正による中断コードの移動"
---

構文木において、 `return` を前要素ではなくなるように、インデントや中括弧を追加することで、後続の処理を復活させる。 

``` diff:after_return_with_if.py
 import sys
 
 
 def main():
     if len(sys.argv) > 1:
         print("Cannot enter arguments.")
-    return
+        return
     print("Am I dead?")
 
 
 main()
```
