# 自動テストメモ

## ファイル比較

```sh
function compare() {
  diff $1 $2
  ERR=$?
  if [[ $ERR != 0 ]]; then
    echo $1 $2 NG
  else
    echo $1 $2 OK
  fi
}
```

## evans 実行

```sh
$ echo {} > query.json
$ cat query.json | evans -r --host localhost -p 50051 cli call TableRepository.List
```

## SQL 実行

```sh
#  -e オプションで実行、結果をCSVファイル化
$ docker compose exec mysql /bin/mysql -u root -ppass mysqldb -e'show tables;' | tr "\\t" ","
```

```sh
# 標準入力で実行
$ echo "select now();" | docker-compose exec -T mysql /bin/mysql -u root -ppass mysqldb
```

## JSON を jq で CSV 化

```json
{
  "body": [
    {
      "fields": [
        {
          "key": "column1",
          "result": "OK",
          "value": "100"
        },
        {
          "key": "column2",
          "result": "NG",
          "value": "200"
        },
        {
          "key": "column3",
          "result": "OK",
          "value": "300"
        }
      ]
    }
  ]
}
```

```sh
$ cat query.json | jq -r '(.body[0].fields[0]|to_entries|map(.key)),(.body[].fields[]|[.[]])|@csv'
"key","result","value"
"column1","OK","100"
"column2","NG","200"
"column3","OK","300"
```
