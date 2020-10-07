---
title: "Zennの執筆体験を向上させるnpm「zmce」を公開した"
emoji: "🔧"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: [ "Zenn", "md", "markdownit", "zmce" ]
published: true
---

# 「zmce」とは

**Zenn Markdown Code Embed** の頭文字を取った **npmパッケージ** です。

**コマンドひとつで、記事中のコードブロックを更新する** ためのツールです。

https://www.npmjs.com/package/zmce
https://github.com/j5c8k6m8/zmce

:::message
現在のバージョンは、v0.0.2です。今後、破壊的な仕様変更をする可能性があります。
:::

# 前提環境 / 想定ユーザ

 1. **[Zenn](https://zenn.dev/)** で **記事** または **本** を執筆している。
 2. **[GitHubリポジトリでZennのコンテンツを管理](https://zenn.dev/zenn/articles/connect-to-github)** している。
 3. 記事または本で、 **別のGitリポジトリのコードを参照** している。
 4. 記事または本を執筆する際、 **参照コードを都度コピペしたくない** と感じている。

:::message alert
現在(10/7)、ロードマップ内の **[MD）コードブロックにファイル名を指定できるように](https://github.com/zenn-dev/zenn-roadmap/issues/22)** がまだされていない状態のため、本ツールの方法ではコードブロックのシンタックスハイライトが無効になります。~~プルリク出せるかな？~~
また、本ツールは、コードブロックのコード名称部分に独自の記法を採用しています。上記のIssuesの対応次第では、本ツールを用いた記事では、参照パスが記事中に表示される可能性があります。
:::

# 導入手順

## 0. 事前準備

 - あらかじめ、 **[Zenn CLIを導入](https://zenn.dev/zenn/articles/install-zenn-cli)** する必用があります。

## 1. zmceをインストールする

Zennのルートディレクトリで、以下のコマンドを実行します。

``` sh
$ npm install zmce # zmceを導入
```

これでディレクトリにzmceがインストールされます。

## 2. submodulesディレクトリを作成する

Zennのルートディレクトリ下に、submodulesディレクトリを作成します。

``` sh
$ mkdir submodules
```

## 3. 導入完了🎉

これでzmceの導入は完了です。以下のコマンドを実行することで、記事または本が更新されます。

:::message
この時点では参照先のGitリポジトリが存在しないため、実際には更新されることはありません。
コマンドが正常終了することを確認してください。
:::

``` sh
$ npx zmce
[START] zmce
[ END ] zmce
```


# 使い方

## 1. サブモジュールを追加する

zmceでは、記事または本で参照する対象のGitリポジトリを **Gitのサブモジュールとして管理** します。
submodulesディレクトリ下に、対象のGitリポジトリをサブモジュールとして追加してください。

``` sh
$ # 以下の例では、zmce自体のGitリポジトリを参照先リポジトリとして設定します。
$ git submodule add https://github.com/j5c8k6m8/zmce.git submodules/zmce
```

## 2. 参照先ファイルを指定する

コードを参照させたい記事、または、本に以下のような記法で、空のコードブロックを記載してください。

````` md:test.md
# テスト記事

``` コード:名称(表示用):参照先パス
```
`````

## 3. コマンドを実行する

以下の、コマンドを実行することで、空のコードブロック内に参照先のパスの内容が記載されます。

``` sh
$ npx zmce
[START] zmce
[articles/test.md] コードブロックを修正しました。
[ END ] zmce
```

````` md:test.md
# テスト記事

``` コード:名称(表示用):参照先パス
参照先パスのファイルの内容
```
`````

参照先パスの内容が変わった場合は、再度コマンドを実行することで、参照内容を最新化できます。


# 注意事項

 - Zennのコードブロック記法は、 **[[小ネタ] 公式に記載されていない、ZennのMarkdown記法#コードブロック](https://zenn.dev/j5c8k6m8/articles/zenn-md-easter-eggs#%E3%82%B3%E3%83%BC%E3%83%89%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF)** の通り、複数の記載方法がありますが、 **「^```(行の先頭からバッククォート3つ)」のみに対応** しています。
 - **「^```(行の先頭からバッククォート3つ)」を含むファイルは、参照先として使用できません。**  
 (警告メッセージが表示されます。)

# 使用例

本記事は以下に、zmceのソースを、zmceコマンドを用いて埋め込んでいます。
本記事の原文(テキスト)は、 **[GitHubで参照](https://github.com/j5c8k6m8/j5c8k6m8-zenn-contents/blob/master/articles/zmce-introduction.md)** できます。

:::details zmceのソースコード
``` ts:zmce.ts:zmce/src/zmce.ts
// ref: https://github.com/zenn-dev/zenn-editor/blob/master/packages/zenn-cli/utils/api/
import fs from "fs-extra";
import { join } from "path";
import colors from "colors/safe";

const articlesDirectoryName = "articles";
const booksDirectoryName = "books";
const modulesDirectoryName = "submodules";
const replaceCodeSymbol = "```";
const replaceCodePattern = new RegExp(`(^${replaceCodeSymbol})([^:\n]*):([^:\n]*):([^:\n]+)(.*$)([^]*?)(^${replaceCodeSymbol}$)`, 'gm');
const checkPattern = new RegExp(`^${replaceCodeSymbol}`, 'm');
const mdRegex = /\.md$/;

export function main() {
  console.info(colors.cyan(`[START] zmce`));

  try {
    fs.readdirSync(join(process.cwd(), modulesDirectoryName));
  } catch (e) {
    console.error(
      colors.red(
        `プロジェクトルートに${modulesDirectoryName}ディレクトリを作成してください`
      )
    );
    process.exitCode = 1;
  }

  let articleFiles;
  try {
    const articleAllFiles = fs.readdirSync(
      join(process.cwd(), articlesDirectoryName)
    );
    articleAllFiles.sort();
    articleFiles = articleAllFiles
      .filter((f) => f.match(mdRegex))
      .map((f) => join(articlesDirectoryName, f));
  } catch (e) {
    console.error(
      colors.red(
        `プロジェクトルートに${articlesDirectoryName}ディレクトリを作成してください`
      )
    );
    process.exitCode = 1;
  }

  const chapterFiles: string[] = [];
  try {
    let allBookDirs = fs.readdirSync(join(process.cwd(), booksDirectoryName));
    let bookDirs = allBookDirs.filter((f) => {
      try {
        return fs
          .statSync(join(process.cwd(), booksDirectoryName, f))
          .isDirectory();
      } catch (e) {
        return false;
      }
    });
    bookDirs.forEach((bookDir) => {
      try {
        const bookChapterAllFiles = fs.readdirSync(
          join(process.cwd(), booksDirectoryName, bookDir)
        );
        bookChapterAllFiles.sort();
        bookChapterAllFiles
          .filter((f) => f.match(mdRegex))
          .forEach((f) =>
            chapterFiles.push(join(booksDirectoryName, bookDir, f))
          );
      } catch (e) {
        // ファイルの読み込み失敗は無視する
      }
    });
  } catch (e) {
    console.error(
      colors.red(
        `プロジェクトルートに${booksDirectoryName}ディレクトリを作成してください`
      )
    );
    process.exitCode = 1;
  }

  if (process.exitCode == 1) {
    console.info(colors.cyan(`[ END ] zmce`));
  } else {
    const codeEmbed = (relativePath: string) => {
      let text;
      try {
        text = fs.readFileSync(join(process.cwd(), relativePath), "utf8");
      } catch (e) {
        return;
      }
      let afterText = text.replace(
        replaceCodePattern,
        (
          match,
          beginMark,
          codeType,
          codeName,
          codePath,
          other,
          code,
          afterMark
        ) => {
          let afterCode;
          try {
            afterCode = fs.readFileSync(
              join(process.cwd(), modulesDirectoryName, codePath.trim()),
              "utf8"
            );
          } catch (e) {
            console.warn(
              colors.yellow(
                `[${relativePath}] モジュールディレクトリに「${codePath.trim()}」ファイルがありません`
              )
            );
            return match;
          }
          if (checkPattern.test(afterCode)) {
            console.warn(
              colors.yellow(
                `[${relativePath}] 「${codePath.trim()}」ファイル内に使用できないパターン(^${replaceCodeSymbol})が含まれています。`
              )
            );
            return match;
          }
          return `${beginMark}${codeType}:${codeName}:${codePath}${other}\n${afterCode}\n${afterMark}`;
        }
      );
      if (afterText != text) {
        fs.writeFileSync(join(process.cwd(), relativePath), afterText, "utf8");
        console.info(`[${relativePath}] コードブロックを修正しました。`);
      }
    };
    articleFiles?.forEach((f) => codeEmbed(f));
    chapterFiles.forEach((f) => codeEmbed(f));
    console.info(colors.cyan(`[ END ] zmce`));
  }
}

```
:::

# その他

 - 今回はじめてnpmパッケージを公開しました。
 - 当初、HTMLのコメント記法を使用した、マクロ置換記法を作ろうと考えました。  
 しかし、Zennのリリース当時はHTMLコメントが使えなかったことと、コードブロックにファイル名を指定するIssuesが既に上がっていたことから、今回のコードブロックのコード名称を拡張する形式にしました。~~先にHTMLのコメント記法に対応が入るとは思わなかった😅~~  
 作成してみて、今の形式もシンプルでわかり易いと考え公開に踏み切りました。
  - 今後、以下のような拡張ができたらいいかと考えています。下記以外でもフィードバックを頂けると励みになります！
    - 除外対象ファイルなどを細かく指定できるconfig(zmceconfig.json?)の作成
    - configにより「```」以外の形式(「````」や「~~~」)への対応
    - 初期化コマンド(npx zmce init)対応
    - HTMLコメント記法を使用した、マクロ置換記法の作成
    - 特定の行のみを参照させるような指定方法の追加
    - コマンドの実行結果の埋め込み等への対応
 - Zennにより、ローカル執筆環境のディレクトリ構成にarticlesやbooksなど、いい意味でConventionが設けられたため、このようなツールが作成しやすい環境ができたと思っています。zmceの思想や使用しているツールは今後、記事を書けたらいいな、と思ってます。
 - 私は心の中で「ジムチェ」と唱えながらコマンドをたたいています。
