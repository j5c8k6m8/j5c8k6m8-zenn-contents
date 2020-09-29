---
title: "ZennのMarkdownのコードを参照させるnpm「zmce」を作ってみた"
emoji: "🐣"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: [ "Zenn", "md", "markdownit" ]
published: true
---

``` ruby:zmce.ts:zmce/src/zmce.ts
a = "123"
put a
b = a + 345
[b, 345].each do |x|
  put x
end
```
