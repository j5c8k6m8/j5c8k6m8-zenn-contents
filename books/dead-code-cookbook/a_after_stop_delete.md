---
title: "👼中断コード削除"
---


`return` を削除することで、後続の処理を復活させる。

``` diff:after_return.rb
-return
 puts 'Am I dead?'
```
