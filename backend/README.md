### ディレクトリ構成

```
.
├── api/ #APIのロジックや実行を司るディレクトリ
│   ├── controller/ #APIの各エンドポイントの処理をここに記述
│   │   └── posts_controller.go
│   ├── middleware/ #認証・認可の共通処理をここに記述
│   │   └── cors.go
│   ├── schema/ APIスキーマモデルをここに定義する
│   │   └── post_schema.go
│   └── router/ ルーティング設定やミドルウェアの埋め込みをここで記述
│       └── router.go
├── application/ #API全体の具体的なロジック(ユースケース)をここに記述
│   ├── list_posts_usecase.go
│   └── get_post_usecase.go
├── config/ #環境変数などで読み取る共通の設定をここに置く
│   └── config.go
├── domain/ #APIのコアな部分
│   ├── apperror/ #API共通のエラー
│   ├── model/ #API共通のドメインモデル
│   │   └── post.go
│   └── repository/ #ドメインモデルをCRUD操作するためのインタフェース
│       └── post_repository.go
├── infrastructure/
│   └── store/ #MySQLの呼び出しやRepositoryの実装をここに書く (自動生成できるといいな)
│       ├── dao/ #テーブル定義をここに書く
│       │   ├── post.go
│       │   └── user.go
│       ├── store.go
│       └── post_repository_impl.go
└── main.go #実行ファイル & DI
```
