package store_test

import (
	"myapp/config"
	"myapp/domain/repository"
	"myapp/infrastructure/store"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
)

var testInjector *do.Injector

func init() {
	testInjector = do.New()
	// NOTE: 本来は環境変数を使うべきだが、便宜上変数代入している
	config.DBHostName = "127.0.0.1"

	do.Provide[*sqlx.DB](testInjector, store.NewStore)

	// Inject repository resources
	do.Provide[repository.HelloWorldRepository](testInjector, store.NewHelloWorldRepository)
	do.Provide[repository.PostRepository](testInjector, store.NewPostRepository)
	do.Provide[repository.UserRepository](testInjector, store.NewUserRepositoryImpl)
}
