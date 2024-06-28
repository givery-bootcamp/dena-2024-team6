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
	"os"

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
	if err := router.SetupRoutes(injector, app); err != nil {
		log.Fatalf("failed to setup routes: %+v\n", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.HostName, config.Port),
		Handler: app.Handler(),
	}

	// スキーマ情報からOpenAPIファイルの自動生成
	appDoc := do.MustInvoke[*openapi3.Reflector](injector)
	generatedDoc, err := appDoc.Spec.MarshalYAML()
	if err != nil {
		log.Printf("failed to generate openapi doc; %+v\n", err)
	} else {
		f, err := os.Create("../docs/api.yaml")
		if err != nil {
			log.Fatalf("failed to load file; %+v\n", err)
		}
		defer f.Close()
		if _, err = f.Write(generatedDoc); err != nil {
			log.Fatalf("failed to generate doc; %+v\n", err)
		}
	}

	// 裏側でサーバを起動
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %+v\n", err)
		}
	}()

	// ctrl + cでサーバをシャットダウンする
	if err := injector.ShutdownOnSIGTERM(); err != nil {
		log.Fatalf("failed to signal shutdown; %+v\n", err)
	}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("something is wrong on shutdown server; %+v\n", err)
	}
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
	do.Provide[repository.CommentRepository](i, store.NewCommentRepositoryImpl)

	// Inject application usecase resources
	do.Provide[application.CreatePostUsecase](i, application.NewCreatePostUsecase)
	do.Provide[application.ListPostUsecase](i, application.NewListPostUsecase)
	do.Provide[application.GetPostDetailUsecase](i, application.NewGetPostDetailUsecase)
	do.Provide[application.UpdatePostUsecase](i, application.NewUpdatePostUsecase)
	do.Provide[application.DeletePostUsecase](i, application.NewDeletePostUsecase)
	do.Provide[application.SigninUsecase](i, application.NewSigninUsecase)
	do.Provide[application.SignupUsecase](i, application.NewSignupUsecase)
	do.Provide[application.GetUserUsecase](i, application.NewGetUserUsecase)
	do.Provide[application.ListCommentsUsecase](i, application.NewListCommentsUsecase)
	do.Provide[application.CreateCommentUsecase](i, application.NewCreateCommentUsecase)
	do.Provide[application.UpdateCommentUsecase](i, application.NewUpdateCommentUsecase)
	do.Provide[application.DeleteCommentUsecase](i, application.NewDeleteCommentUsecase)
	// Inject controller resources
	do.Provide[*controller.PostController](i, controller.NewPostController)
	do.Provide[*controller.AuthController](i, controller.NewAuthController)

	// Inject middleware resources
	do.Provide[*middleware.AuthorizationMiddleware](i, middleware.NewAuthorizationMiddleware)
}
