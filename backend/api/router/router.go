package router

import (
	"myapp/api/controller"
	"myapp/api/middleware"
	"myapp/api/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

func SetupRoutes(i *do.Injector, app *gin.Engine) error {
	// 依存するControllerの依存関係の取得
	postController := do.MustInvoke[*controller.PostController](i)
	authController := do.MustInvoke[*controller.AuthController](i)

	authorizationMiddleware := do.MustInvoke[*middleware.AuthorizationMiddleware](i)

	// OpenAPIの依存関係の取得
	appDoc := do.MustInvoke[*openapi3.Reflector](i)

	// OpenAPIの基本定義を設定
	appDoc.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	appDoc.Spec.Info.
		WithTitle("Web開発研修6班 API").
		WithVersion("1.0.0").
		WithDescription("FY24卒Web開発研修6班のAPI仕様書です")

	appDoc.Spec.Servers = append(appDoc.Spec.Servers, openapi3.Server{
		URL: "http://localhost:9000",
	})

	app.Use(middleware.CorsMiddleware())
	// HealthCheckOpe / GET
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	healthCheckOpe, _ := appDoc.NewOperationContext(http.MethodGet, "/")
	healthCheckOpe.SetID("healthCheck")
	healthCheckOpe.SetSummary("APIのセルフチェック")
	healthCheckOpe.SetTags("operation")
	if err := appDoc.AddOperation(healthCheckOpe); err != nil {
		return err
	}

	// listPostsOpe /posts GET
	app.GET("/posts", postController.ListPost)
	listPostsOpe, _ := appDoc.NewOperationContext(http.MethodGet, "/posts")
	listPostsOpe.SetID("listPosts")
	listPostsOpe.SetSummary("投稿の一覧を取得")
	listPostsOpe.SetTags("post")
	listPostsOpe.AddRespStructure(new([]schema.PostResponse), openapi.WithHTTPStatus(http.StatusOK))
	listPostsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	listPostsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	listPostsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(listPostsOpe); err != nil {
		return err
	}

	// getPostOpe /posts/{id} GET
	app.GET("/posts/:postid", postController.GetPost)
	getPostOpe, _ := appDoc.NewOperationContext(http.MethodGet, "posts/{id}")
	getPostOpe.SetID("getPost")
	getPostOpe.AddReqStructure(new(schema.PostRequest))
	getPostOpe.SetSummary("投稿をIDから取得")
	getPostOpe.SetTags("post")
	getPostOpe.AddRespStructure(new(schema.PostResponse), openapi.WithHTTPStatus(http.StatusOK))
	getPostOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	getPostOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	getPostOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusNotFound))
	getPostOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(getPostOpe); err != nil {
		return err
	}

	// listPostComments /posts/{postId}/comments GET
	app.GET("/posts/:postid/comments", postController.ListComments)
	listCommnetsOpe, _ := appDoc.NewOperationContext(http.MethodGet, "posts/{postId}/comments")
	listCommnetsOpe.SetID("listPostComments")
	listCommnetsOpe.AddReqStructure(new(schema.CommentListRequest))
	listCommnetsOpe.SetSummary("対象の投稿のコメント一覧を取得")
	listCommnetsOpe.SetTags("post")
	listCommnetsOpe.AddRespStructure(new([]schema.CommentResponse), openapi.WithHTTPStatus(http.StatusOK))
	listCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	listCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	listCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusForbidden))
	listCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(listCommnetsOpe); err != nil {
		return err
	}

	// createPostComments /posts/{postId}/comments POST
	app.POST("/posts/:postid/comments", postController.CreateComment)
	createCommnetsOpe, _ := appDoc.NewOperationContext(http.MethodPost, "posts/{postId}/comments")
	createCommnetsOpe.SetID("createPostComments")
	createCommnetsOpe.AddReqStructure(new(schema.CreateCommentRequest))
	createCommnetsOpe.SetSummary("対象の投稿のコメントを追加")
	createCommnetsOpe.SetTags("post")
	createCommnetsOpe.AddRespStructure(new(schema.CommentResponse), openapi.WithHTTPStatus(http.StatusCreated))
	createCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	createCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	createCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusForbidden))
	listCommnetsOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(createCommnetsOpe); err != nil {
		return err
	}

	// signInOpe /signin POST
	app.POST("/signin", authController.SignIn)
	signInOpe, _ := appDoc.NewOperationContext(http.MethodPost, "/signin")
	signInOpe.SetID("signIn")
	signInOpe.SetSummary("ユーザのログインを実行")
	signInOpe.SetTags("auth")
	signInOpe.AddReqStructure(new(schema.LoginRequest))
	signInOpe.AddRespStructure(new(schema.UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	signInOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	signInOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	signInOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(signInOpe); err != nil {
		return err
	}

	authRequired := app.Group("/")
	authRequired.Use(authorizationMiddleware.Exec())

	// signOutOpe /signout POST
	authRequired.POST("/signout", authController.SignOut)
	signOutOpe, _ := appDoc.NewOperationContext(http.MethodPost, "/signout")
	signOutOpe.SetID("signOut")
	signOutOpe.SetTags("auth")
	signOutOpe.SetSummary("ユーザのログアウトを実行")
	signOutOpe.AddSecurity("user auth")
	signOutOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	signOutOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusBadRequest))
	signOutOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(signOutOpe); err != nil {
		return err
	}

	// getCurrentUserOpe /user GET
	authRequired.GET("/user", authController.GetCurrentUser)
	getCurrentUserOpe, _ := appDoc.NewOperationContext(http.MethodGet, "/user")
	getCurrentUserOpe.SetID("getCurrentUser")
	getCurrentUserOpe.SetSummary("現在ログインしているユーザを取得")
	getCurrentUserOpe.SetTags("auth")
	getCurrentUserOpe.AddSecurity("user auth")
	getCurrentUserOpe.AddRespStructure(new(schema.UserResponse), openapi.WithHTTPStatus(http.StatusOK))
	getCurrentUserOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusUnauthorized))
	getCurrentUserOpe.AddRespStructure(new(schema.ErrorResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	if err := appDoc.AddOperation(getCurrentUserOpe); err != nil {
		return err
	}

	return nil
}

func NewAppDoc(i *do.Injector) (*openapi3.Reflector, error) {
	return openapi3.NewReflector(), nil
}
