---
title: "🧟ホイスティング"
---

`ホイスティング` は、ほとんどのデッドコードパターンで利用できる。

任意の文を実行することは出来ないが、定義の巻き上げに伴い、外側の変数を隠蔽し挙動を替えることが可能だ。以下のコードの挙動を考えて欲しい。

``` js:hoisting_zombie.js:./projects/javascript/src/zombie/hoisting_zombie.js
var d = "Am I dead?";
(() => {
    if (d) {
        console.log(d);
    }
    return;
    var d = "I am zombie!";
})()
```

上記の実行結果は以下のようになる。

``` console
$ # コード実行
$ node src/zombie/hoisting_zombie.js
$ 
```

**コンソールには何も出力されない** 。

また、`🔩eslint` では以下の通り `Unreachable code` が出力される。

``` console
$ # eslint
$ eslint src/zombie/hoisting_zombie.js 

/app/javascript/src/zombie/hoisting_zombie.js
  1:5  error  'd' is assigned a value but never used  no-unused-vars
  7:5  error  Unreachable code                        no-unreachable

✖ 2 problems (2 errors, 0 warnings)

$ 
```

`🔩eslint` の `Unreachable code` を信じて、以下のようにコードを修正すると、挙動が変わるだろう。

``` js:hoisting_zombie_2.js:./projects/javascript/src/zombie/hoisting_zombie_2.js
var d = "Am I dead?";
(() => {
    if (d) {
        console.log(d);
    }
    return;
})()
```

``` console
$ # コード実行
$ node src/zombie/hoisting_zombie_2.js 
Am I dead?
$ 
```

**コンソールには「Am I dead?」と出力された** 。

**つまり、ゾンビを ~~殺して~~ 削除して しまったことで、別の ~~死体~~ デッドコード からゾンビを復活させてしまったのである。**

原理は知っていれば簡単である。 **`ホイスティング` は、定義のみを巻き上げ、初期化は巻き上げない** [^1] という理解があれば十分だろう。

[^1]: https://developer.mozilla.org/ja/docs/Glossary/Hoisting

なお、本書では以下の様な `function` によるホイスティングはデッドコードとして扱わない。

``` js:hoisting_function.js:./projects/javascript/src/zombie/hoisting_function.js
var d = "Am I dead?";
(() => {
    if (d) {
        console.log(d);
    }
    return;
    function d() {
        console.log("I am zombie!");
    }
})()
```

宣言であることが明確であり、`🔩eslint` でも `Unreachable code` としては扱われないためだ。

``` console
$ # コード実行
$ node src/zombie/hoisting_function.js 
[Function: d]
$ # eslint
$ eslint src/zombie/hoisting_function.js 

/app/javascript/src/zombie/hoisting_function.js
  1:5  error  'd' is assigned a value but never used  no-unused-vars

✖ 1 problem (1 error, 0 warnings)

$ 
```

`ホイスティング` を行う言語は多くない。本書の対象範囲では、 `🔧JavaScript` のみである。また、`🔧JavaScript` においても `var` を使用せず、 `let`, `const` を使用すれば回避策になるだろう。


# 言語毎

- `Python (対象外)`
- `Ruby (対象外)`
- [`JavaScript`](#🔧javascript)
- `Java (対象外)`
- `Go (対象外)`

## 🔧JavaScript

``` js:hoisting_zombie.js:./projects/javascript/src/zombie/hoisting_zombie.js
var d = "Am I dead?";
(() => {
    if (d) {
        console.log(d);
    }
    return;
    var d = "I am zombie!";
})()
```
