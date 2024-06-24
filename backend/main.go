package main

import (
	"context"
	"fmt"
	"log"
	"myapp/api/controller"
	"myapp/api/middleware"
	"myapp/api/router"
	"myapp/application"
	"myapp/config"
	"myapp/domain/repository"
	"myapp/domain/service"
	"myapp/infrastructure/idtoken"
	"myapp/infrastructure/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
	"github.com/swaggest/openapi-go/openapi3"
)

func main() {
	ctx := context.Background()

	injector := do.New()
	injectDependencies(injector)

	app := gin.Default()
	router.SetupRoutes(injector, app)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.HostName, config.Port),
		Handler: app.Handler(),
	}

	// スキーマ情報からOpenAPIファイルの自動生成
	// appDoc := do.MustInvoke[*openapi3.Reflector](injector)
	// generatedDoc, err := appDoc.Spec.MarshalYAML()
	// if err != nil {
	// 	log.Printf("failed to generate openapi doc; %+v\n", err)
	// } else {
	// 	f, err := os.Create("../docs/api.yaml")
	// 	if err != nil {
	// 		log.Fatalf("failed to load file; %+v\n", err)
	// 	}
	// 	defer f.Close()
	// 	f.Write(generatedDoc)
	// }

	// 裏側でサーバを起動
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %+v\n", err)
		}
	}()

	// ctrl + cでサーバをシャットダウンする
	injector.ShutdownOnSIGTERM()
	server.Shutdown(ctx)
}

func injectDependencies(i *do.Injector) {
	// Inject external resources
	do.Provide[*openapi3.Reflector](i, router.NewAppDoc)
	do.Provide[*sqlx.DB](i, store.NewStore)

	// Inject service resources
	do.Provide[service.IdtokenService](i, idtoken.NewIDTokenJwtImpl)

	// Inject repository resources
	do.Provide[repository.HelloWorldRepository](i, store.NewHelloWorldRepository)
	do.Provide[repository.PostRepository](i, store.NewPostRepository)
	do.Provide[repository.UserRepository](i, store.NewUserRepositoryImpl)

	// Inject application usecase resources
	do.Provide[application.ListPostUsecase](i, application.NewListPostUsecase)
	do.Provide[application.GetPostDetailUsecase](i, application.NewGetPostDetailUsecase)
	do.Provide[application.SigninUsecase](i, application.NewSigninUsecase)

	// Inject controller resources
	do.Provide[*controller.PostController](i, controller.NewPostController)
	do.Provide[*controller.AuthController](i, controller.NewAuthController)

	// Inject middleware resources
	do.Provide[*middleware.AuthorizationMiddleware](i, middleware.NewAuthorizationMiddleware)
}
