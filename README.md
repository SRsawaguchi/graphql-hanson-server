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

