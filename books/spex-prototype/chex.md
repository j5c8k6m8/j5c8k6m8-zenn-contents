---
title: "文字集合クラス"
---

文字集合クラスを定義します。簡易的な正規表現の文字クラスです。

本書での、文字集合クラスの文字列表現は以下の通りとします。

- 正規表現と同様、`[]` 内に対象の文字を列挙します。
- `.` は全ての文字を表わす全体集合を表わします。
- 対象の文字を取り除いた補集合を定義したい場合は `[]` 内の先頭に `^` を付与します。
- 単一文字の場合は、 `[]` を省略できます。

:::message
**範囲指定、文字クラスの略記法、論理演算記法は対象外** とします。
:::

文字集合クラスは、論理演算が可能です。以下に演算結果例とソースコードを記載します。

# 演算結果例

### 単項演算

|#|chex|invert|備考|
|---|---|---|---|
|1|`^.`|`.`|空集合|
|2|`.`|`^.`|全体集合|
|3|`a`|`^a`|-|
|4|`^a`|`a`|-|
|5|`[ab]`|`^[ab]`|-|
|6|`^[ab]`|`[ab]`|-|

#### 二項演算

|#|chex1|chex2|chex1 or chex2|chex1 and chex2|include[^1]|備考|
|---|---|---|---|---|---|---|
|1|`^.`|`.`|`.`|`^.`|`False` / `True`|-|
|2|`^.`|`a`|`a`|`^.`|`False` / `True`|-|
|3|`.`|`a`|`.`|`a`|`True` / `False`|-|
|4|`a`|`a`|`a`|`a`|`True` / `True`|-|
|5|`a`|`b`|`[ab]`|`^.`|`False` / `False`|-|
|6|`a`|`^a`|`.`|`^.`|`False` / `False`|-|
|7|`a`|`^b`|`^b`|`a`|`False` / `True`|-|
|8|`^a`|`^b`|`.`|`^[ab]`|`False` / `Flase`|-|
|9|`a`|`[ab]`|`[ab]`|`a`|`False` / `True`|-|
|10|`[ab]`|`[bc]`|`[abc]`|`b`|`False` / `False`|-|
|11|`[ab]`|`^[bc]`|`^c`|`a`|`False` / `False`|-|
|12|`^[ab]`|`^[bc]`|`^b`|`^[abc]`|`False` / `False`|-|

[^1]: chex1がchex2を含むか/ chex2がchex1を含むか

``` python:chex.py:spex-m8p-py/src/spexm8p/chex.py
from spexm8p.token import Token


# 文字集合クラス (Character Expression)
class Chex:
    def __init__(self, chars, include_flg):
        self._char_set = set(chars)
        # self._include_flg => True: include, False: exclude(補集合)
        self._include_flg = True if include_flg else False
        self._len = len(self._char_set)
        if self._len == 0:
            if self._include_flg:
                self._kind = 0  # 空集合
            else:
                self._kind = 1  # 全集合
        else:
            self._kind = 2
        self._str = None  # 遅延評価

    # 空集合確認
    def blank(self):
        if self._kind == 0:
            return True
        else:
            return False

    # 全集合確認
    def whole(self):
        if self._kind == 1:
            return True
        else:
            return False

    # 否定論理演算(~)
    def __invert__(self):
        if self._kind == 2:
            return Chex(self._char_set, not self._include_flg)
        else:
            if self._kind == 0:
                return Chex.WHOLE
            elif self._kind == 1:
                return Chex.BLANK

    # 論理和演算(|)
    def __or__(self, other):
        if self._kind == 2:
            if other._kind == 2:
                if self._include_flg:
                    if other._include_flg:
                        return Chex(self._char_set | other._char_set, True)
                    else:
                        return Chex(other._char_set - self._char_set, False)
                else:
                    if other._include_flg:
                        return Chex(self._char_set - other._char_set, False)
                    else:
                        return Chex(self._char_set & other._char_set, False)
            elif other._kind == 0:
                return self
            elif other._kind == 1:
                return other
        elif self._kind == 0:
            return other
        elif self._kind == 1:
            return self

    # 論理積演算(&)
    def __and__(self, other):
        if self._kind == 2:
            if other._kind == 2:
                if self._include_flg:
                    if other._include_flg:
                        return Chex(self._char_set & other._char_set, True)
                    else:
                        return Chex(self._char_set - other._char_set, True)
                else:
                    if other._include_flg:
                        return Chex(other._char_set - self._char_set, True)
                    else:
                        return Chex(self._char_set | other._char_set, False)
            elif other._kind == 0:
                return other
            elif other._kind == 1:
                return self
        elif self._kind == 0:
            return self
        elif self._kind == 1:
            return other

    # 包含確認
    def include(self, other):
        if self._kind == 2:
            if other._kind == 2:
                if self._include_flg and not other._include_flg:
                    return False
                return (~self & other).blank()
            elif other._kind == 0:
                return True
            elif other._kind == 1:
                return False
        elif self._kind == 0:
            return False
        elif self._kind == 1:
            return True

    def __eq__(self, other):
        if self._kind == other._kind:
            if self._kind != 2:
                return True
            return str(self) == str(other)
        else:
            return False

    def __str__(self):
        if self._str is None:
            if self._kind == 2:
                chars = ''.join(sorted(self._char_set))
                if self._include_flg and self._len == 1:
                    self._str = chars
                else:
                    prefix = '' if self._include_flg else Token.DENY
                    self._str = f'{Token.CH_S}{prefix}{chars}{Token.CH_E}'
            elif self._kind == 0:
                self._str = f'{Token.CH_S}{Token.CH_E}'
            elif self._kind == 1:
                self._str = Token.WHOL
        return self._str

    BLANK = None
    WHOLE = None


Chex.BLANK = Chex([], True)
Chex.WHOLE = Chex([], False)

```
