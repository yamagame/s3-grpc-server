# GRPC-S3 サンプル
## 実行方法

```bash
# grpsサーバー起動
$ docker-compose up -d
$ docker compose exec grpc-server /bin/bash
/app # ./script/start.sh
```

```bash
# evans起動
$ docker compose exec grpc-server /bin/bash
/app # ./script/repl.sh
```
