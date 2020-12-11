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
