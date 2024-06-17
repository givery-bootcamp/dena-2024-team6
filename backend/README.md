### ディレクトリ構成

```
.
├── cmd/ #実行するアプリ (例：API,マイグレーションコマンドなど) 単位で管理するディレクトリ
│   ├── api/ #APIのロジックや実行を司るディレクトリ
│   │   ├── controller/ #APIの各エンドポイントの処理をここに記述
│   │   ├── middleware/ #認証・認可の共通処理をここに記述
│   │   ├── schema/ APIスキーマモデルをここに定義する
│   │   ├── di/ APIに必要な依存関係を組み立てる処理をここに記述
│   │   └── api.go #APIの起動・実行をここに記述 (これを go run api.goするイメージ)
│   └── migration/ #マイグレーションの実行を司るディレクトリ
├── config/ #環境変数などで読み取る共通の設定をここに置く
│   └── config.go
├── core/ #プロジェクト全体
│   ├── model/
│   │   └── habit.go
│   ├── repository/
│   │   └── post_repository.go
│   └── usecase/
└── pkg/
    └── store/ #MySQLの呼び出しやRepositoryの実装をここに書く (自動生成できるといいな)
        ├── store.go
        └── post_repository_impl.go
```
