# HOW TO USE?

## 起動方法

- docker-compose を使用して API を起動します。
- docker-compose がインストールされていることが前提となります。
  - インストールされていない場合は、[こちら](https://docs.docker.com/compose/install/)を参考に Docker Desktop をインストールしてください。

### 1. docker コンテナを起動する

```bash
$ make docker_up
```

### 2. Enjoy!

- http://localhost:8080 で API が動作します。


## 停止方法

```bash
$ make docker_down
```

## リスタート方法

```bash
$ make docker_restart
```


## データベースのデータを全削除

- データベースのデータを全削除する場合は以下のコマンドを実行してください。

```bash
$ make docker_remove_vol
```