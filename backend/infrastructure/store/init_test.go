package store_test

import (
	"myapp/domain/repository"
	"myapp/infrastructure/store"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
)

var testInjector *do.Injector

func init() {
	testInjector = do.New()

	do.Provide[*sqlx.DB](testInjector, store.NewStore)

	// Inject repository resources
	do.Provide[repository.HelloWorldRepository](testInjector, store.NewHelloWorldRepository)
	do.Provide[repository.PostRepository](testInjector, store.NewPostRepository)
	do.Provide[repository.UserRepository](testInjector, store.NewUserRepositoryImpl)
}
