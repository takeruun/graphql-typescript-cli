# 概要

GrapQL サーバー と Next.js クライアント
graphql-codegen　をしようしてスキーマ駆動開発サンプル

## 使い方
1. 初期化
```
yarn graphql-code-generator init
```

2. もろもろ設定
```
? What type of application are you building? Application built with React
? Where is your schema?: (path or url) http://localhost:3000/query
? Where are your operations and fragments?: ./src/graph/*.graphql
? Pick plugins: TypeScript (required by other typescript plugins), TypeScript Operations (operations and fragments), TypeScript React Apollo (typed
components and HOCs)
? Where to write the output: ./src/@generated/graphql.ts
? Do you want to generate an introspection file? Yes
? How to name the config file? codegen.yml
? What script in package.json should run the codegen? graphql-codegen
```
3. 自動生成
スキーマファイル修正後、下記コマンド実行
```
yarn graphql-codegen
```

# 参考サイト

Apollo 公式

https://www.apollographql.com/docs/

graphql-codegen 公式

https://www.graphql-code-generator.com/
