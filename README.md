# gin-sample
ginのサンプル  

# 使い方
下のコマンドを打ったら、このURLを読み込んでみる。  
http://localhost:9998/v1/designs    

```
cd ./infrastructure/docker

docker-compose up -d

docker-compose exec app bash

make init

go run main.go
```

# マイグレーション
```
make migrate
```

# シード
```
make seed
```

# テーブルクリア
```
make clear_db
```

# 構文チェック

```
make check
```

# 構文自動修正
```
make fmt
```

# golint
```
make lint
```

# 単体テスト実行
```
make test
```


# ファイル設計  
Domain(Entities)  
ユビキタス言語化したデータを格納  
ユーザ情報、記事情報  

Usecase  
  - port

Http  
  - controller  
  - response  
  - middleware  

repository  
docker  
