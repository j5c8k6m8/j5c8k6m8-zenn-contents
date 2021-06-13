---
title: "パーサーモジュール"
---

簡易的にパーサーを作成します。正規表現エンジンの理論とは直接関係無いため、説明は割愛します。

なお、本書のパーサーとは異なりますが、 **[正規表現エンジンを作ろう （3）](https://codezine.jp/article/detail/3158)** において、正規表現エンジンのコンパイラの作成についての説明が記載されています。

:::message
本モジュールのparse関数では、構文木の作成までを行い、次のbuilderモジュールで文字集合クラスの構築を行います。
:::

以下にソースコードを記載します。

``` python:parser.py:spex-m8p-py/src/spexm8p/parser.py
from spexm8p.token import Token

# 入力文字を、制御用として扱う文字。文字として扱う場合はエスケープが必要
ESCAPES = [
    Token.AND, Token.OR, Token.INVT,
    Token.CH_S, Token.CH_E, Token.WHOL, Token.DENY,
    Token.REPT,
    Token.SP_S, Token.SP_E,
]


def tokenize(input_str):
    ret = []
    escape_flg = False
    for char in list(input_str):
        if escape_flg:
            if char == Token.ESC:
                ret.append(Token.ESC + Token.ESC)
            elif char in ESCAPES:
                ret.append(Token.ESC + char)
            else:
                ret.append(Token.ESC)
                ret.append(char)
            escape_flg = False
        elif char == Token.ESC:
            escape_flg = True
        else:
            ret.append(char)
    if escape_flg:
        ret.append(Token.ESC + Token.ESC)
    return ret


def parse(tokens):
    return _parse_and_or(tokens)


def _parse_and_or(tokens):
    group_level = 0
    for i, token in enumerate(tokens):
        if token == Token.SP_S:
            group_level += 1
        elif token == Token.SP_E:
            group_level -= 1
        elif group_level == 0:
            if token == Token.AND:
                return _get_and_node(
                    _parse_and_or(tokens[:i]),
                    _parse_and_or(tokens[i+1:])
                )
            elif token == Token.OR:
                return _get_or_node(
                    _parse_and_or(tokens[:i]),
                    _parse_and_or(tokens[i+1:])
                )
    if group_level != 0:
        raise Exception(f'SyntaxError {Token.SP_E} not enough')
    return _parse_invert(tokens)


def _parse_invert(tokens):
    if len(tokens) == 0:
        raise Exception('SyntaxError invalid blank node')
    elif tokens[0] == Token.INVT:
        return _get_invert_node(_parse_concat(tokens[1:]))
    else:
        return _parse_concat(tokens)


def _parse_concat(tokens):
    group_level = 0
    in_ch = False
    current_tokens = []
    node_kind = 0  # 0: 登録不要, 1: group登録, 2: 単一文字登録, 3: 複数文字登録
    nodes = []
    for i, token in enumerate(tokens):
        if node_kind != 0:
            if node_kind == 1:
                node = _parse_and_or(current_tokens)
            elif node_kind == 2:
                node = _parse_inc_chex(current_tokens)
            elif node_kind == 3:
                node = _parse_chex(current_tokens)
            else:
                assert False  # unreachable code
            current_tokens = []
            node_kind = 0
            if token == Token.REPT:
                node = _get_repeat_node(node)
                nodes.append(node)
                continue
            nodes.append(node)
        if group_level == 0:
            if token == Token.INVT:
                raise Exception(f'SyntaxError {Token.INVT} Not at the beginning')
            elif token == Token.SP_E:
                raise Exception(f'SyntaxError {Token.SP_E} invalid position')
            elif token == Token.REPT:
                raise Exception(f'SyntaxError {Token.REPT} invalid position')
            elif token == Token.SP_S:
                if in_ch:
                    raise Exception(f'SyntaxError {Token.SP_S} invalid position')
                group_level += 1
            elif token == Token.CH_S:
                if in_ch:
                    raise Exception(f'SyntaxError {Token.CH_S} invalid position')
                in_ch = True
            elif token == Token.CH_E:
                if not in_ch:
                    raise Exception(f'SyntaxError {Token.CH_E} invalid position')
                in_ch = False
                node_kind = 3
            else:
                current_tokens.append(token)
                if not in_ch:
                    node_kind = 2
        else:
            if token == Token.SP_S:
                group_level += 1
                current_tokens.append(token)
            elif token == Token.SP_E:
                group_level -= 1
                if group_level != 0:
                    current_tokens.append(token)
                else:
                    node_kind = 1
            else:
                current_tokens.append(token)
    if group_level != 0:
        raise Exception(f'SyntaxError {Token.SP_E} not enough')
    if in_ch:
        raise Exception(f'SyntaxError {Token.CH_E} not enough')
    if node_kind != 0:
        if node_kind == 1:
            nodes.append(_parse_and_or(current_tokens))
        elif node_kind == 2:
            nodes.append(_parse_inc_chex(current_tokens))
        elif node_kind == 3:
            nodes.append(_parse_chex(current_tokens))
        else:
            assert False  # unreachable code
    if len(nodes) == 0:
        raise Exception('SyntaxError invalid blank node')
    elif len(nodes) == 1:
        return nodes[0]
    else:
        return _get_concat_node(nodes)


def _parse_chex(tokens):
    if len(tokens) > 0 and tokens[0] == Token.DENY:
        return _parse_exc_chex(tokens[1:])
    else:
        return _parse_inc_chex(tokens)


def _parse_inc_chex(tokens):
    if Token.WHOL in tokens:
        return _get_exc_chex([])
    else:
        return _get_inc_chex(tokens)


def _parse_exc_chex(tokens):
    if Token.WHOL in tokens:
        return _get_inc_chex([])
    else:
        return _get_exc_chex(tokens)


def _get_and_node(left, right):
    return {'type': 'and', 'left': left, 'right': right}


def _get_or_node(left, right):
    return {'type': 'or', 'left': left, 'right': right}


def _get_concat_node(nodes):
    return {'type': 'concat', 'nodes': nodes}


def _get_invert_node(node):
    return {'type': 'invert', 'node': node}


def _get_repeat_node(node):
    return {'type': 'repeat', 'node': node}


def _get_inc_chex(tokens):
    return {'type': 'inc_chex', 'tokens': tokens}


def _get_exc_chex(tokens):
    return {'type': 'exc_chex', 'tokens': tokens}

```