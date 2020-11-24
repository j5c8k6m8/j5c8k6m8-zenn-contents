---
title: "📑目次"
---

- [`はじめに`](./introduction)
- [`おわりに`](./afterword)

# デッドコードレシピ一覧

1. [`🔖中断を利用するパターン`](./p_after)
    1. [`🧪return後のコード`](./r_after_return)
    1. [`🧪throw後のコード`](./r_after_throw)
    1. [`🧪break後のコード`](./r_after_break)
    1. [`🧪goto後のコード`](./r_after_goto)
    1. [`🧪exit後のコード`](./r_after_exit)
    1. [`🧪常にthrowする関数呼出後のコード`](./r_after_func_throw)
    1. [`🧪常にbreakする高階関数呼出後のコード`](./r_after_break_yield)
    1. [`🧪常に実行時エラーとなる処理後のコード`](./r_after_runtime_error)
    1. [`🧪常に条件がfalseとなるassert後のコード`](./r_after_assert)
1. [`🔖終らない処理を利用するパターン`](./p_forever)
    1. [`🧪無限ループ後のコード`](./r_forever_loop)
    1. [`🧪終らない関数呼出後のコード`](./r_forever_func)
1. [`🔖単独で常にfalseとなる条件を利用するパターン`](./p_simple_if)
    1. [`🧪単独で常にfalse(リテラル)となる条件のif文`](./r_simple_if_literal)
    1. [`🧪単独で常にfalse(定数)となる条件のif文`](./r_simple_if_const)
    1. [`🧪単独で常にfalse(変数)となる条件のif文`](./r_simple_if_variable)
    1. [`🧪単独で常にfalse(演算結果)となる条件のif文`](./r_simple_if_operation)
1. [`🔖並列な複数の条件を利用するパターン`](./p_parallel_if)
    1. [`🧪部分集合条件のelseif`](./r_parallel_elseif)
    1. [`🧪部分集合条件のcase(switch)`](./r_parallel_switch_case)
    1. [`🧪空集合条件のelse`](./r_parallel_else)
    1. [`🧪空集合条件のdefault(switch)`](./r_parallel_switch_default)
1. [`🔖ネストした複数の条件を利用するパターン`](./p_nest_if)
    1. [`🧪部分集合条件のif(ネスト)`](./r_nest_if)
1. [`🔖if以外の条件を利用するパターン`](./p_cond_other)
    1. [`🧪for継続条件が空集合条件`](./r_cond_other_for)
    1. [`🧪while継続条件が空集合条件`](./r_cond_other_while)
    1. [`🧪三項演算子が空集合条件`](./r_cond_other_ternary)
    1. [`🧪短絡評価が空集合条件`](./r_cond_other_short_circuit)
    1. [`🧪catchすることのないcatch節`](./r_cond_other_catch)
    1. [`🧪次要素のないイテレータ`](./r_cond_other_iterator)
    1. [`🧪常にnullに対するnull条件演算子`](./r_cond_other_null)
1. [`🔖高階関数を利用するパターン`](./p_func)
    1. [`🧪空集合要素に対する高階関数`](./r_func_empty)
    1. [`🧪余分な高階関数の引数`](./r_func_arg)
1. [`🔖定義のみで使用されないパターン`](./p_def) ※代入のみは対象外、パブリックも対象外/無名関数は対象外
    1. [`🧪定義のみの関数`](./r_def_func) ⇒上書き含む
    1. [`🧪定義のみのプライベートメソッド`](./r_def_method)
    1. [`🧪オーバライドされたメソッド`](./r_def_override)
1. [`🔖その他のパターン`](./p_other)
    1. [`🧪１回しか実行されないのジェネレータのyield後のコード`](./r_other_generator_yield)
    1. マクロによる置換を使った何か？
    1. プロミスを使った何か？
    1. 遅延評価を使った何か？？
    1. 呼び出されないデストラクタとか？

# レイズ一覧

1. [`👼中断コード削除`](./a_after_stop_delete)
1. [`👼ネスト修正による中断コードの移動`](./a_after_stop_move)
1. [`👼実行時エラー修正`](./a_after_fix_runtime_error)
1. [`👼関数呼出処理修正`](./a_fix_call)
1. `👼条件削除`
1. `👼条件修正`


# ゾンビ化一覧

1. [`🧟goto文のラベルによるジャンプ`](./z_goto)
1. [`🧟ホイスティング`](./z_hoisting)
1. [`🧟実行時除去`](./z_compile_delete)

# デッドコードツール一覧

1. [`🔪終らない関数`](./t_forever_func)
1. [`🔪部分集合条件と空集合条件`](./t_cond_set)
1. [`🔪比較演算子を用いた常に同じ真偽値となる演算`](./t_cond_comp)
1. [`🔪数値演算を用いた常に同じ真偽値となる演算`](./t_cond_num)
1. [`🔪文字列を用いた常に同じ真偽値となる演算`](./t_cond_str)
1. [`🔪正規表現を用いた常に同じ真偽値となる演算`](./t_cond_regex)
1. [`🔪空集合要素を作る演算`](./t_empty_set)

# 付録

1. [`🧰対象言語と環境`](./environment)







void hoge(double a) {
    if (a != a) {

    }
}
↑NaN

void (Double.isInfinited(d1) && Double.isinfinite(d2) && d1 != d2)
↑プラスマイナス無限大



静的な言語で、型指定を無視したメソッドを追加する。

あとがき。
デッドコードを作れないプログラミング言語
