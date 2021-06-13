---
title: "ビルダーモジュール"
---

最後に、今まで作成してきたクラス/モジュールを組み合わせて、以下のように `from spexm8p.builder import spex` 後に、 `spex()` 関数で、文字列集合クラスのインスタンスが作成できるように、ビルダークラスを作成します。

``` console:sample
# python3
Python 3.9.2 (default, Mar  4 2021, 15:30:45)
[GCC 8.3.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> from spexm8p.builder import spex
>>> # 文字列'a'の補集合は、'a'以外の1文字、もしくは、2文字以上の文字列であることの確認
>>> spex('!a') == spex('[^a]|..+')
True
>>>
```

以下にソースコードを記載します。これにて、全てのソースコードの作成が完了しました。次のページから実際の動作結果を見ていきます。

``` python:builder.py:spex-m8p-py/src/spexm8p/builder.py
from functools import reduce
from spexm8p.spex import Chex
from spexm8p.spex import Spex
from spexm8p.parser import tokenize, parse


def spex(spex_str):
    return _build_spex(parse(tokenize(spex_str)))


def _build_spex(parsed):
    if parsed['type'] == 'inc_chex':
        return Spex.build_by_chex(Chex(parsed['tokens'], True))
    elif parsed['type'] == 'exc_chex':
        return Spex.build_by_chex(Chex(parsed['tokens'], False))
    elif parsed['type'] == 'or':
        return _build_spex(parsed['left']) | _build_spex(parsed['right'])
    elif parsed['type'] == 'and':
        return _build_spex(parsed['left']) & _build_spex(parsed['right'])
    elif parsed['type'] == 'invert':
        return ~_build_spex(parsed['node'])
    elif parsed['type'] == 'repeat':
        return _build_spex(parsed['node']).repeat()
    elif parsed['type'] == 'concat':
        return reduce(lambda a, b: a.concat(b), [_build_spex(x) for x in parsed['nodes']])
    else:
        assert False  # unreachable code

```