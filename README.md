# GraphQL ハンズオン(Server)

GraphQLを試してみる。  

参考  
https://www.howtographql.com/graphql-go/0-introduction/

## gqlgenのインストール
https://github.com/99designs/gqlgen  
`gqlgen`はGoでGraphQLサーバを開発するためのライブラリや各種ツール。  

```
go get -u github.com/99designs/gqlgen
```

## 初期化

```
go mod init
gqlgen init
```

## Schemaを書き換える

`/graph/schema.graphql`にスキーマを定義する。  
定義した後、以下のコマンドを実行してコードを自動生成する。  

```
gqlgen generate
```

## resolverを書き換える

`/graph/schema.resolvers.go`を書き換える。  
そして、 `go run server.go`を実行してサーバを起動する。  

サンプルでは、`Links()`メソッドを書きかえ。（commitログ参照)

## playground経由で実行
ブラウザで`http://localhost:8080`にアクセスし、playgroundから以下のようなクエリを送信する。  

```
query {
	links{
    title
    address,
    user{
      name
    }
  }
}
```

指定したデータが得られる事を確認する。

## mutations
`graph/schema.resolvers.go`を編集する。  
ここでは、`CreateLink()`を編集する。  
※最初はstaticなデータを生成して返すように実装し、クライアント側（playground）から呼び出せることを確認しよう。  

実装できたら、playgroundから以下のようなリクエストを送り、上記で設定したレスポンスが返ってくるかどうか確認しよう。  

```
mutation {
  createLink(input: { title: "new link", address: "http://address.org" }) {
    title
    user {
      name
    }
    address
  }
}
```

## RDB(PostgreSQL)のセットアップ

### RDBMSの起動
以下のコマンドを利用してPostgreSQLのコンテナを起動する。  

```
make rundb
```

※一度`make rundb`したら、`make stopdb`で停止、`make startdb`で再起動できる。  

### マイグレーション
```
make migrate
```
