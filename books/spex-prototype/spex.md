---
title: "文字列集合クラス"
---

文字列集合クラスは、論理演算できる正規表現の核となるものです。

文字列集合クラスは、内部的にはDFAとなります。

> #### DFAの定義
> 数学的には、DFAは以下の5つの要素で定義されます。
> 
>  - 状態の集合（丸）
>  - 文字の集合（Unicode文字）
>  - 遷移関数（矢印）
>    - 引数: 今の状態, 文字
>    - 戻り値: 次の状態
>  - 開始状態（最初の矢印がさす丸）
>  - 受理状態の集合（二重丸）
>
>  ~[正規表現エンジンを作ろう （2）](https://codezine.jp/article/detail/3154?p=3) より~

今回は簡易的なため、以下の通り実装します。

 - 状態はDFA毎に整数値(`sid`)で保持する。
   - `sid` を採番/取得する `SidGen` クラスを定義する
 - 開始状態の `sid` は、 `0` とする
 - 遷移関数を表わす `Transition` クラスを定義する。`Transition` は以下のフィールドを保持する。
   - 今の状態(`_fr_sid`)
   - 文字集合(`_chex`)
   - 次の状態(`_to_sid`)
 - 文字列集合クラスは、遷移関数(`Transition`)の集合と、受理状態(`sid`)の集合を保持する

文字列集合クラスをDFAとすることで、否定演算は受理状態を反転させる事で実現できます。

:::message
否定演算を受理状態の反転として実現させるため、DFAについては全ての未受理状態を含めて構築する点に注意してください。
:::

他の、論理和、論理積、連結、繰返しの演算の実装については、NFAをDFAに変換する際に用いる部分集合構成法が参考になります。

> #### 部分集合構成法
>NFAをDFAに変換するには、部分集合構成法（subset construction）と呼ばれる方法を利用します。
>
> (中略)
>
> このような考え方を元にして、NFAの「取りうる状態の集合」をDFAの「状態」としてNFAからDFAを作成する方法を、部分集合構成法（subset construction）と呼びます。
>
> ~[正規表現エンジンを作ろう （5）](https://codezine.jp/article/detail/3168) より~

つまり、遷移先として複数状態を取りえる事があるが、そのパターンを1つの状態として作成していくことで、NFAで実行可能だった演算を、DFAでも実行可能とします。

:::message
NFAでの空文字（ε）での遷移は、DFAでは合成されてしまうため、繰返しを0文字以上ではなく、1文字以上の演算とします。
なお、0文字以上の繰り返しは、論理和と1文字以上の繰り返しの組み合わせで置き換え可能です。
:::

また、文字列集合クラスは構築後、DFAとしてのデータを保持します。
DFAを可視化するために、 `mermaid` メソッドにて、状態遷移を描画するためのコード(文字列)を出力できるようにします。

以下にソースコードを記載します。演算結果例は、ビルダークラス作成後の使用例のページ以降で示します。

``` python:spex.py:spex-m8p-py/src/spexm8p/spex.py
from spexm8p.chex import Chex


# 文字列集合クラス (String Pattern Expression)
# 各状態をsid(state id)として保持する。
# sid =  0 は、開始状態として常に存在する前提とする。
class Spex:
    def __init__(self, tn_list, accepts_sids):
        self._accepts_sids = accepts_sids
        self._tns_dict = {}
        for tn in tn_list:
            fr_sid = tn._fr_sid
            if fr_sid not in self._tns_dict:
                self._tns_dict[fr_sid] = []
            self._tns_dict[fr_sid].append(tn)
        if len(self._accepts_sids) == 0:
            self._kind = 0  # 空集合
        elif len(self._accepts_sids) == len(self._tns_dict) - 1:
            self._kind = 1  # 全集合
        else:
            self._kind = 2

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
        tn_list = []
        accepts_sids = set()
        for sid in self._tns_dict:
            if sid != 0 and sid not in self._accepts_sids:
                accepts_sids.add(sid)
            tn_list.extend(self._tns_dict[sid])
        return Spex(tn_list, accepts_sids)

    # 論理和演算(|)
    def __or__(self, other):
        if self._kind == 2:
            if other._kind == 2:
                new_tn_list = []
                new_accepts_sids = set()
                Spex._calc_and_or(
                    0, new_tn_list, new_accepts_sids,
                    self, other, 0, 0,
                    SidGen(), {}, 0,  # or
                )
                return Spex(new_tn_list, new_accepts_sids)
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
                new_tn_list = []
                new_accepts_sids = set()
                Spex._calc_and_or(
                    0, new_tn_list, new_accepts_sids,
                    self, other, 0, 0,
                    SidGen(), {}, 1,  # and
                )
                return Spex(new_tn_list, new_accepts_sids)
            elif other._kind == 0:
                return other
            elif other._kind == 1:
                return self
        elif self._kind == 0:
            return self
        elif self._kind == 1:
            return other

    def concat(self, other):
        new_tn_list = []
        new_accepts_sids = set()
        Spex._calc_concat(
            0, new_tn_list, new_accepts_sids,
            self, other, 0, [],
            SidGen(), {},
        )
        return Spex(new_tn_list, new_accepts_sids)

    def repeat(self):
        new_tn_list = []
        new_accepts_sids = set()
        Spex._calc_repeat(
            0, new_tn_list, new_accepts_sids,
            self, [0],
            SidGen(), {}
        )
        return Spex(new_tn_list, new_accepts_sids)

    # 包含確認
    def include(self, other):
        if self._kind == 2:
            if other._kind == 2:
                return (~self & other).blank()
            elif other._kind == 0:
                return True
            elif other._kind == 1:
                return False
        elif self._kind == 0:
            return False
        elif self._kind == 1:
            return True

    def __str__(self):
        ret = ""
        for tns in self._tns_dict.values():
            for tn in tns:
                ret += f'{tn._fr_sid} -- {tn._chex} --> {tn._to_sid}\n'
        ret += f'[{", ".join([str(sid) for sid in self._accepts_sids])}]'
        return ret

    def __eq__(self, other):
        if self._kind == other._kind:
            if self._kind != 2:
                return True
            return self.include(other) and other.include(self)
        else:
            return False

    def mermaid(self):
        ret = "graph LR\n"
        for sid in self._tns_dict:
            if sid == 0:
                ret += "    0(( ))\n"
            elif sid == -1:
                ret += "    -1(( ))\n"
            elif sid == -2:
                ret += "    -2(( ))\n"
            else:
                ret += f"    {sid}( )\n"
        for sid in self._tns_dict:
            if sid == 0:
                ret += "    style 0 fill:#000,stroke-width:0px\n"
            elif sid == -1:
                ret += "    style -1 fill:#adb5bd,stroke-width:0px\n"
            elif sid == -2:
                ret += "    style -2 fill:#adb5bd,stroke:#dc3545,stroke-width:4px\n"
            elif sid in self._accepts_sids:
                ret += f"    style {sid} stroke:#dc3545,stroke-width:4px\n"
        for tns in self._tns_dict.values():
            for tn in tns:
                ret += f'    {tn._fr_sid} -- "{tn._chex}" --> {tn._to_sid}\n'
        return ret

    @classmethod
    def build_by_chex(cls, chex):
        if chex.blank():
            return Spex([
                Transition(0, 1, Chex.WHOLE),
                Transition(1, 1, Chex.WHOLE),
            ], set())
        else:
            return Spex([
                Transition(0, 1, chex),
                Transition(0, 2, ~chex),
                Transition(1, 2, Chex.WHOLE),
                Transition(2, 2, Chex.WHOLE),
            ], {1})

    @classmethod
    def build_whole(cls):
        return Spex([
            Transition(0, 1, Chex.WHOLE),
            Transition(1, 1, Chex.WHOLE)
        ], {1})

    @classmethod
    def _calc_and_or(cls, new_fr_sid, new_tn_list, new_accepts_sids, spex1, spex2, spex1_sid, spex2_sid, sid_gen, sid_dict_by_skey, ope_kind):
        # 1. 作成対象となる遷移先のパターンを抽出するため、chexを細分化する
        chex_pattern = [Chex.WHOLE]
        tmp_chex_pattern = []
        for target_chex in chex_pattern:
            for tn in spex1._tns_dict[spex1_sid]:
                and_chex = target_chex & tn._chex
                if not and_chex.blank():
                    tmp_chex_pattern.append(and_chex)
        chex_pattern = tmp_chex_pattern
        tmp_chex_pattern = []
        for target_chex in chex_pattern:
            for tn in spex2._tns_dict[spex2_sid]:
                and_chex = target_chex & tn._chex
                if not and_chex.blank():
                    tmp_chex_pattern.append(and_chex)
        chex_pattern = tmp_chex_pattern
        # 2. パターン毎に新しい遷移を追加する。 遷移先(new_to_sid)が新しい場合は再帰呼出を行う
        for target_chex in chex_pattern:
            # 2-1. 対象の遷移(next)をみつける
            for tn in spex1._tns_dict[spex1_sid]:
                if tn._chex.include(target_chex):
                    next_spex1_sid = tn._to_sid
                    break
            else:
                assert False  # unreachable code
            for tn in spex2._tns_dict[spex2_sid]:
                if tn._chex.include(target_chex):
                    next_spex2_sid = tn._to_sid
                    break
            else:
                assert False  # unreachable code
            # 2-2. 新しい遷移先のキーを作成する(再帰呼出済かの確認)
            skey = f'${next_spex1_sid}/${next_spex2_sid}'
            if skey in sid_dict_by_skey:
                # 2-2-1. 遷移先が登録済みの場合は、遷移先のsidを取得。(再帰呼出は行わない)
                new_to_sid = sid_dict_by_skey[skey]
            else:
                # 2-2-2. 遷移先が未登録の場合
                # 2-2-2-1. 新しい遷移先のsidを採番
                new_to_sid = sid_gen.get()
                # 2-2-2-2. 新しい遷移先を登録
                sid_dict_by_skey[skey] = new_to_sid
                # 2-2-2-3. 新しい遷移先の受理状態を評価(new_accepts_sids追加判定)
                if ope_kind == 0:  # or
                    if next_spex1_sid in spex1._accepts_sids or next_spex2_sid in spex2._accepts_sids:
                        new_accepts_sids.add(new_to_sid)
                elif ope_kind == 1:  # and
                    if next_spex1_sid in spex1._accepts_sids and next_spex2_sid in spex2._accepts_sids:
                        new_accepts_sids.add(new_to_sid)
                else:
                    assert False  # unreachable code
                # 2-2-2-4. 再帰呼出
                cls._calc_and_or(new_to_sid, new_tn_list, new_accepts_sids, spex1, spex2, next_spex1_sid, next_spex2_sid, sid_gen, sid_dict_by_skey, ope_kind)
            # 2-3. new_tn_list追加
            new_tn_list.append(Transition(new_fr_sid, new_to_sid, target_chex))

    @classmethod
    def _calc_concat(cls, new_fr_sid, new_tn_list, new_accepts_sids, spex1, spex2, spex1_sid, spex2_sids, sid_gen, sid_dict_by_skey):
        # 1. 作成対象となる遷移先のパターンを抽出するため、chexを細分化する
        chex_pattern = [Chex.WHOLE]
        tmp_chex_pattern = []
        for target_chex in chex_pattern:
            for tn in spex1._tns_dict[spex1_sid]:
                and_chex = target_chex & tn._chex
                if not and_chex.blank():
                    tmp_chex_pattern.append(and_chex)
        chex_pattern = tmp_chex_pattern
        if spex1_sid in spex1._accepts_sids:
            # spex1が受理状態の場合は、spex2の最初の遷移を考慮する
            tmp_chex_pattern = []
            for target_chex in chex_pattern:
                for tn in spex2._tns_dict[0]:
                    and_chex = target_chex & tn._chex
                    if not and_chex.blank():
                        tmp_chex_pattern.append(and_chex)
            chex_pattern = tmp_chex_pattern
        for spex2_sid in spex2_sids:
            tmp_chex_pattern = []
            for target_chex in chex_pattern:
                for tn in spex2._tns_dict[spex2_sid]:
                    and_chex = target_chex & tn._chex
                    if not and_chex.blank():
                        tmp_chex_pattern.append(and_chex)
            chex_pattern = tmp_chex_pattern
        # 2. パターン毎に新しい遷移を追加する。 遷移先(new_to_sid)が新しい場合は再帰呼出を行う
        for target_chex in chex_pattern:
            # 2-1. 対象の遷移(next)をみつける
            for tn in spex1._tns_dict[spex1_sid]:
                if tn._chex.include(target_chex):
                    next_spex1_sid = tn._to_sid
                    break
            else:
                assert False  # unreachable code
            next_spex2_sids = set()
            if spex1_sid in spex1._accepts_sids:
                # spex1が受理状態の場合は、spex2の最初の遷移を考慮する
                for tn in spex2._tns_dict[0]:
                    if tn._chex.include(target_chex):
                        next_spex2_sids.add(tn._to_sid)
                        break
                else:
                    assert False  # unreachable code
            for spex2_sid in spex2_sids:
                for tn in spex2._tns_dict[spex2_sid]:
                    if tn._chex.include(target_chex):
                        next_spex2_sids.add(tn._to_sid)
                        break
                else:
                    assert False  # unreachable code
            # 2-2. 新しい遷移先のキーを作成する(再帰呼出済かの確認)
            if len(next_spex2_sids) == 0:
                skey = f'${next_spex1_sid}/'
            else:
                skey = f'${next_spex1_sid}/${"-".join([str(x) for x in sorted(next_spex2_sids)])}'
            if skey in sid_dict_by_skey:
                # 2-2-1. 遷移先が登録済みの場合は、遷移先のsidを取得。(再帰呼出は行わない)
                new_to_sid = sid_dict_by_skey[skey]
            else:
                # 2-2-2. 遷移先が未登録の場合
                # 2-2-2-1. 新しい遷移先のsidを採番
                new_to_sid = sid_gen.get()
                # 2-2-2-2. 新しい遷移先を登録
                sid_dict_by_skey[skey] = new_to_sid
                # 2-2-2-3. 新しい遷移先の受理状態を評価(new_accepts_sids追加判定)
                for next_spex2_sid in next_spex2_sids:
                    if next_spex2_sid in spex2._accepts_sids:
                        new_accepts_sids.add(new_to_sid)
                        break
                # 2-2-2-4. 再帰呼出
                cls._calc_concat(new_to_sid, new_tn_list, new_accepts_sids, spex1, spex2, next_spex1_sid, next_spex2_sids, sid_gen, sid_dict_by_skey)
            # 2-3. new_tn_list追加
            new_tn_list.append(Transition(new_fr_sid, new_to_sid, target_chex))

    @classmethod
    def _calc_repeat(cls, new_fr_sid, new_tn_list, new_accepts_sids, spex, spex_sids, sid_gen, sid_dict_by_skey):
        # 1. 作成対象となる遷移先のパターンを抽出するため、chexを細分化する
        chex_pattern = [Chex.WHOLE]
        for spex_sid in spex_sids:
            tmp_chex_pattern = []
            for target_chex in chex_pattern:
                for tn in spex._tns_dict[spex_sid]:
                    and_chex = target_chex & tn._chex
                    if not and_chex.blank():
                        tmp_chex_pattern.append(and_chex)
            chex_pattern = tmp_chex_pattern
        # 受理状態の場合は、最初の遷移を考慮する
        for spex_sid in spex_sids:
            if spex_sid in spex._accepts_sids:
                tmp_chex_pattern = []
                for target_chex in chex_pattern:
                    for tn in spex._tns_dict[0]:
                        and_chex = target_chex & tn._chex
                        if not and_chex.blank():
                            tmp_chex_pattern.append(and_chex)
                chex_pattern = tmp_chex_pattern
                break
        # 2. パターン毎に新しい遷移を追加する。 遷移先(new_to_sid)が新しい場合は再帰呼出を行う
        for target_chex in chex_pattern:
            # 2-1. 対象の遷移(next)をみつける
            next_spex_sids = set()
            for spex_sid in spex_sids:
                if spex_sid in spex._accepts_sids:
                    # 受理状態の場合は、最初の遷移を考慮する
                    for tn in spex._tns_dict[0]:
                        if tn._chex.include(target_chex):
                            next_spex_sids.add(tn._to_sid)
                            break
                    break
            for spex_sid in spex_sids:
                for tn in spex._tns_dict[spex_sid]:
                    if tn._chex.include(target_chex):
                        next_spex_sids.add(tn._to_sid)
                        break
                else:
                    assert False  # unreachable code
            # 2-2. 新しい遷移先のキーを作成する(再帰呼出済かの確認)
            skey = "-".join([str(x) for x in sorted(next_spex_sids)])
            if skey in sid_dict_by_skey:
                # 2-2-1. 遷移先が登録済みの場合は、遷移先のsidを取得。(再帰呼出は行わない)
                new_to_sid = sid_dict_by_skey[skey]
            else:
                # 2-2-2. 遷移先が未登録の場合
                # 2-2-2-1. 新しい遷移先のsidを採番
                new_to_sid = sid_gen.get()
                # 2-2-2-2. 新しい遷移先を登録
                sid_dict_by_skey[skey] = new_to_sid
                # 2-2-2-3. 新しい遷移先の受理状態を評価(new_accepts_sids追加判定)
                for next_spex_sid in next_spex_sids:
                    if next_spex_sid in spex._accepts_sids:
                        new_accepts_sids.add(new_to_sid)
                        break
                # 2-2-2-4. 再帰呼出
                cls._calc_repeat(new_to_sid, new_tn_list, new_accepts_sids, spex, next_spex_sids, sid_gen, sid_dict_by_skey)
            # 2-3. new_tn_list追加
            new_tn_list.append(Transition(new_fr_sid, new_to_sid, target_chex))

    BLANK = None
    WHOLE = None


class Transition:
    def __init__(self, fr_sid, to_sid, chex):
        self._fr_sid = fr_sid
        self._to_sid = to_sid
        self._chex = chex


class SidGen:
    def __init__(self):
        self._sid = 0

    def get(self):
        self._sid += 1
        return self._sid


Spex.BLANK = Spex.build_by_chex(Chex.BLANK)
Spex.WHOLE = Spex.build_whole()

```
