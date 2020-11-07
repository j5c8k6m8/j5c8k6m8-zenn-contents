---
title: "Zennでアドベントカレンダーを行うためのGitHub Actions"
emoji: "🎄"
type: "idea" # tech: 技術記事 / idea: アイデア
topics: ["zenn", "GitHubActions"]
published: true
---

2020年は Zenn のアドベントカレンダー機能のリリースが間に合わなそうだ。

アドベンドカレンダーはWikipedia[^1]によると

- インターネット上などで、12月の1日から25日までに、特定のテーマに沿って毎日ブログなどに記事を投稿していくという企画
- 複数人実施の場合は、カレンダーを管理するサイトを軸にある程度の範囲（プログラミング言語や使用する技術など）を決めて参加者を募り、順に投稿を行うというスタイルが多い。

[^1]: https://ja.wikipedia.org/wiki/%E3%82%A2%E3%83%89%E3%83%99%E3%83%B3%E3%83%88%E3%82%AB%E3%83%AC%E3%83%B3%E3%83%80%E3%83%BC#%E4%BC%81%E7%94%BB%E3%81%A8%E3%81%97%E3%81%A6%E3%81%AE%E3%82%A2%E3%83%89%E3%83%99%E3%83%B3%E3%83%88%E3%82%AB%E3%83%AC%E3%83%B3%E3%83%80%E3%83%BC

つまり、アドベンドカレンダーは **記事投稿先だけあれば実施** できる 。Zennでもアドベンドカレンダー目次記事のようなものを作り、 ~~勝手に~~ アドベンドカレンダーを実施することはできるだろう。

Zennで、一人アドベンドカレンダーに挑戦したいと考えたが、その際に面倒だと感じることは、
**毎日、記事を公開する操作をしなければならない** ことだ。

つまり、**事前に記事を書いておき、公開をスケジューリングする** 機能が一番必要だ。

**[GitHubリポジトリでZennのコンテンツを管理](https://zenn.dev/zenn/articles/connect-to-github)** していれば、GitHubActionsの `on.schedule` [^2] で実現できそうなのでやってみた。

[^2]: https://docs.github.com/ja/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#onschedule

なお、公開用ブランチへのマージする設定の記載は **[Github Actions でブランチの操作を行う - RitoLabo](https://www.ritolab.com/entry/206)** を参考にさせていただきました。


# GitHub Actionsで記事公開をスケジューリングする

日付毎に、公開する記事を書くブランチを用意する。
今回は、アドベンドカレンダー用なので、 `adcal-1` ～ `adcal-25` というブランチを用意する。

25ブランチだけなので、コマンドで作成し、GitHubにもpushしておく。

``` console:ブランチを作成し、pushする
git checkout -b adcal-1 && git push --set-upstream origin adcal-1
git checkout -b adcal-2 && git push --set-upstream origin adcal-2
git checkout -b adcal-3 && git push --set-upstream origin adcal-3
git checkout -b adcal-4 && git push --set-upstream origin adcal-4
git checkout -b adcal-5 && git push --set-upstream origin adcal-5
git checkout -b adcal-6 && git push --set-upstream origin adcal-6
git checkout -b adcal-7 && git push --set-upstream origin adcal-7
git checkout -b adcal-8 && git push --set-upstream origin adcal-8
git checkout -b adcal-9 && git push --set-upstream origin adcal-9
git checkout -b adcal-10 && git push --set-upstream origin adcal-10
git checkout -b adcal-11 && git push --set-upstream origin adcal-11
git checkout -b adcal-12 && git push --set-upstream origin adcal-12
git checkout -b adcal-13 && git push --set-upstream origin adcal-13
git checkout -b adcal-14 && git push --set-upstream origin adcal-14
git checkout -b adcal-15 && git push --set-upstream origin adcal-15
git checkout -b adcal-16 && git push --set-upstream origin adcal-16
git checkout -b adcal-17 && git push --set-upstream origin adcal-17
git checkout -b adcal-18 && git push --set-upstream origin adcal-18
git checkout -b adcal-19 && git push --set-upstream origin adcal-19
git checkout -b adcal-20 && git push --set-upstream origin adcal-20
git checkout -b adcal-21 && git push --set-upstream origin adcal-21
git checkout -b adcal-22 && git push --set-upstream origin adcal-22
git checkout -b adcal-23 && git push --set-upstream origin adcal-23
git checkout -b adcal-24 && git push --set-upstream origin adcal-24
git checkout -b adcal-25 && git push --set-upstream origin adcal-25
```

ブランチをマージして、公開用ブランチにpushするワークフローを作成する。
GitHub Actions(1日目と2日目)は以下の通り。

:::message
 - 公開用のブランチを、`zenn` という名称でデフォルトブランチに設定しています。
 - --ff-onlyを指定していないため、マージコミットのためにuserの設定をしています。
:::

``` yml:.github/workflows/adcal-1.yml:../.github/workflows/adcal-1.yml
name: adcal-1
on:
  schedule:
    # 12/1 の 7:00 (JST) に処理を実行する。(UTC の 22:00 は JST だと翌日の 7:00)
    - cron: '0 22 30 11 *'
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: merge adcal branch
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          git config user.name j5c8k6m8
          git config user.email j5c8k6m8@gmail.com
          git merge -m "Merge branch for Advent calendar" origin/adcal-1
      - name: push default branch
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: git push origin zenn

```

``` yml:.github/workflows/adcal-2.yml:../.github/workflows/adcal-2.yml
name: adcal-2
on:
  schedule:
    # 12/2 の 7:00 (JST) に処理を実行する。(UTC の 22:00 は JST だと翌日の 7:00)
    - cron: '0 22 1 12 *'
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: merge adcal branch
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          git config user.name j5c8k6m8
          git config user.email j5c8k6m8@gmail.com
          git merge -m "Merge branch for Advent calendar" origin/adcal-2
      - name: push default branch
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: git push origin zenn

```

これも、日付を変えて、1～25まで準備すれば準備は完了[^3]。

[^3]: 全量は **[本記事のGitHubリポジトリ](https://github.com/j5c8k6m8/j5c8k6m8-zenn-contents/tree/zenn/.github/workflows)** を参照してください。

:::message 

### GitHub Actionsの設定ではまったところ

#### 実行間隔

スケジュールされた時間から、10分くらいはタイミングがずれることがあった。

公式にも、 **「スケジュールされたワークフローを実行できる最短の間隔は、5 分ごとに 1 回です。」** と記載があり、動作確認の為、5分未満後に実行されるようなワークフローを記述しても実行されないことがある。

#### fetch-depth

`fetch-depth: 0` にしておかないと、デフォルトではマージ対象のブランチの情報を取得できないため失敗する。

:::

あとは、公開したい日付に合わせたブランチで 記事/本 を書いてpushしておけばＯＫ。

GitHub + Zennのみで完結(ブラウザのみでも実施できる)するところがいい。

~~公開リポジトリの場合は、ブランチを見れば公開前の情報が見えてしまうのは仕方ない。~~

# さいごに

Zennでは本が書けるので、 **毎日チャプターの内容が公開されていく本** を書いてみようかと思っている。

1～25までworkflowを作らずに、実行時の時間から判断するなどもう少しworkflowを賢くすることも考えたが、来年はアドベントカレンダー機能がリリースされているだろうし、普段使いのアイディアも現状はないので、まずはYAGNIで。
