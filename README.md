# Zenn Contents

[✍️ How to use](https://zenn.dev/zenn/articles/zenn-cli-guide)

# Zenn CLIでの開発手順(Docker)

## 1. CLIをインストール / Zenn用セットアップ

### 前提条件

 - [ZennとGitHubリポジトリとの連携](/zenn/articles/connect-to-github)が行われていること
 - Dockerがインストールされていること
 - GitHubのリポジトリを作成し、開発環境にcloneされていること
 - Dockerファイルが作成されていること
 - Zennのルートディレクトリのパスに対して、環境変数(ZENN_HOME)が設定されていること  
 (本README.md内のコマンドでのみ使用)


``` shell-session
$ docker build -t zenn-contents .
$ docker run --name zenn --rm -v $ZENN_HOME:/app/zenn-contents -w /app/zenn-contents -it zenn-contents bash
/app/zenn-contents # npm install
``` 

## 起動

``` shell
docker run --name zenn --rm -d -v $ZENN_HOME:/app/zenn-contents -w /app/zenn-contents -p 80:8000 zenn-contents
```

## ログイン(コマンド実行)

``` shell
docker exec -it zenn bash
``` 
