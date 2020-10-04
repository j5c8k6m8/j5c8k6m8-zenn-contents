---
title: "[小ネタ] 公式に記載されていない、ZennのMarkdown記法"
emoji: "🐣"
type: "idea" # tech: 技術記事 / idea: アイデア
topics: [ "Zenn", "md", "markdownit" ]
published: true
---

# はじめに

**公式 [ZennのMarkdown記法](https://zenn.dev/zenn/articles/markdown-guide)** に記載されていない、小ネタ

[markdown-it](https://github.com/markdown-it/markdown-it) を使っている[^1]ようなので、そっちの仕様を含む。

[^1]: https://zenn.dev/catnose99/articles/zenn-dev-stack


# Styling text

公式には、インラインスタイルとして、**太字**まで記載があるが、
***太字 かつ イタリック***の指定ができる。

なお、`*` の他、`_`(アンダースコア)でも同様の効果となる

上記は [GFM](https://docs.github.com/en/github/writing-on-github/basic-writing-and-formatting-syntax#styling-text) と同じ。

```
*Italic(イタリック)*
**Bold(太字)**
***All bold and italic(太字かつ斜体)***
_Italic(イタリック)_
__Bold(太字)__
___All bold and italic(太字かつ斜体)___
```

*Italic(イタリック)*
**Bold(太字)**
***All bold and italic(太字かつ斜体)***
_Italic(イタリック)_
__Bold(太字)__
___All bold and italic(太字かつ斜体)___


# Task list items (extension)

[GFM](https://github.github.com/gfm/#task-list-items-extension-) と同じく使用できる。

```
- [x] foo
  - [ ] bar
  - [x] baz
- [ ] bim
```

- [x] foo
  - [ ] bar
  - [x] baz
- [ ] bim

# Autolinks (extension)

[GFM](https://github.github.com/gfm/#autolinks-extension-) と同じく、有効な `www.` で始まるドメインと、`http:` 及び `https:` についてはリンクで表示されるが、[emailのリンク](https://github.github.com/gfm/#extended-email-autolink)はつかない

www.commonmark.org
http://commonmark.org
foo@bar.baz


# その他

## 矢印の場合にフォント変換?

矢印の表記はフォントが変わる？

`->`: ->
`<->`: <->
`<=>`: <=>
`<-`: <-
`-->`: -->
`<--`: <--
`---->`: ---->

と思ったら、 **Google Web FontsのInterの仕様** みたい。（有名？？）
~~実はこの記事を書こうと思ったきっかけ~~

↓参考
https://fonts.google.com/specimen/Inter?query=Inter&preview.text=%3C----&preview.text_type=custom


## HTML

### インラインコメントは有効

アップデートによりインラインコメントを記載できるようになった。

@[tweet](https://twitter.com/zenn_dev/status/1310821183567192064)

### 上記以外は無効

`<font color="red">aaa</font>`: <font color="red">aaa</font>

などはHTMLとして解釈されない

※ 文字色は強調のパターンを増やしたくないため、サポートされない予定

https://github.com/zenn-dev/zenn-roadmap/issues/87

## Markdownの原文を見る

現在はサポートされていないはず。(公開設定されたGitHub連携していれば確認可能)

