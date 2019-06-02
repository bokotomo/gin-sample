# gin-sample
ginのサンプル  

# 使い方
```
cd ./infrastructure/docker

docker-compose up -d

docker-compose exec app bash

make init

go run migrate.go

go run seed.go

go run main.go
```

# 構文チェック

```
make check
```

# 構文自動修正
```
make lint
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
