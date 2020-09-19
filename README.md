# Zenn Contents

[✍️ How to use](https://zenn.dev/zenn/articles/zenn-cli-guide)

# Zenn CLIでの開発手順(Docker)


## 1. CLIをインストール / Zenn用セットアップ

### 前提条件

 - [ZennとGitHubリポジトリとの連携](/zenn/articles/connect-to-github)が行われていること
 - Dockerがインストールされていること
 - GitHubのリポジトリを作成し、開発環境にcloneされていること
 - Dockerファイルが作成されていること

``` shell
$ docker build -t zenn-contents .
$ docker run --name zenn --rm -v $PWD:/app/zenn-contents -w /app/zenn-contents -it zenn-contents ash
/app # cd zenn-contents
/app/zenn-contents # npm init --yes
/app/zenn-contents # npm install zenn-cli
/app/zenn-contents # npx zenn init
``` 

## 起動

``` shell
docker run --name zenn --rm -d -v $PWD:/app/zenn-contents -w /app/zenn-contents -p 80:8000 zenn-contents
# Linuxで開発ユーザとコンテナのユーザを合わせる場合
# docker run --name zenn --rm -d -v $PWD:/app/zenn-contents -w /app/zenn-contents -p 80:8000 -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -u $(id -u $USER):$(id -g $USER) zenn-contents


## コマンド実行

``` shell
docker exec -it zenn ash
# Linuxで開発ユーザとコンテナのユーザを合わせる場合
# docker exec -it -u $(id -u $USER):$(id -g $USER) zenn ash 
``` 


