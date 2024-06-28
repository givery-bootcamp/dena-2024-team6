package controller

import (
	"context"
	"log"
	"myapp/api/middleware"
	"myapp/api/schema"
	"myapp/application"
	"myapp/config"
	"myapp/domain/apperror"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type PostController struct {
	createPostUsecase    application.CreatePostUsecase
	listPostUsecase      application.ListPostUsecase
	getPostUsecase       application.GetPostDetailUsecase
	updatePostUsecase    application.UpdatePostUsecase
	deletePostUsecase    application.DeletePostUsecase
	listCommentsUsecase  application.ListCommentsUsecase
	createCommentUsecase application.CreateCommentUsecase
	updateCommentUsecase application.UpdateCommentUsecase
	deleteCommentUsecase application.DeleteCommentUsecase
}

func NewPostController(i *do.Injector) (*PostController, error) {
	createPostUsecase := do.MustInvoke[application.CreatePostUsecase](i)
	listPostUsecase := do.MustInvoke[application.ListPostUsecase](i)
	getPostUsecase := do.MustInvoke[application.GetPostDetailUsecase](i)
	updatePostUsecase := do.MustInvoke[application.UpdatePostUsecase](i)
	deletePostUsecase := do.MustInvoke[application.DeletePostUsecase](i)
	listCommentsUsecase := do.MustInvoke[application.ListCommentsUsecase](i)
	createCommentUsecase := do.MustInvoke[application.CreateCommentUsecase](i)
	updateCommentUsecase := do.MustInvoke[application.UpdateCommentUsecase](i)
	deleteCommentUsecase := do.MustInvoke[application.DeleteCommentUsecase](i)
	return &PostController{
		createPostUsecase:    createPostUsecase,
		listPostUsecase:      listPostUsecase,
		getPostUsecase:       getPostUsecase,
		updatePostUsecase:    updatePostUsecase,
		deletePostUsecase:    deletePostUsecase,
		listCommentsUsecase:  listCommentsUsecase,
		createCommentUsecase: createCommentUsecase,
		updateCommentUsecase: updateCommentUsecase,
		deleteCommentUsecase: deleteCommentUsecase,
	}, nil
}

func (pc PostController) ListPost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	result, err := pc.listPostUsecase.Execute(ctx)
	if apperror.Is(err, apperror.CodeNotFound) {
		c.JSON(404, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	resp := make([]schema.PostResponse, len(result.Posts))
	for i, p := range result.Posts {
		resp[i] = schema.PostResponse{
			ID:     p.ID,
			Title:  p.Title,
			UserID: p.UserID,
		}
	}

	c.JSON(200, resp)
}

func (pc PostController) GetPost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	result, err := pc.getPostUsecase.Execute(ctx, application.GetPostDetailUsecaseInput{
		ID: postID,
	})
	if apperror.Is(err, apperror.CodeNotFound) {
		c.JSON(404, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.JSON(200, schema.PostDetailResponse{
		ID:        result.Post.ID,
		Title:     result.Post.Title,
		Body:      result.Post.Body,
		CreatedAt: result.Post.CreatedAt,
		UpdatedAt: result.Post.UpdatedAt,
		UserResponse: schema.UserResponse{
			ID:       result.Post.UserID,
			UserName: result.Post.UserName,
		},
	})
}

func (pc PostController) CreatePost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	var req schema.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "リクエストの形式が誤っています")))
		return
	}

	// TODO: そのうちバリデーションライブラリ導入したい
	// RUNEを使っているのは、日本語の文字数をちゃんと正しく取るため
	if utf8.RuneCountInString(req.Title) > 20 || len(req.Title) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "タイトルの長さは1~20文字である必要があります")))
		return
	}
	if utf8.RuneCountInString(req.Body) > 100 || len(req.Body) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "本文の長さは1~100文字である必要があります")))
		return
	}

	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "投稿を作成するにはログインする必要があります")))
		return
	}

	result, err := pc.createPostUsecase.Execute(ctx, application.CreatePostUsecaseInput{
		UserID: user.ID,
		Title:  req.Title,
		Body:   req.Body,
	})
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
	}

	c.JSON(201, schema.MutationSchema{
		TargetID: result.PostID,
		Message:  "新しい投稿を作成しました",
	})
}

// update post
func (pc PostController) UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	var req schema.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "リクエストの形式が誤っています")))
		return
	}

	// TODO: そのうちバリデーションライブラリ導入したい
	// RUNEを使っているのは、日本語の文字数をちゃんと正しく取るため
	if utf8.RuneCountInString(req.Title) > 20 || len(req.Title) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "タイトルの長さは1~20文字である必要があります")))
		return
	}
	if utf8.RuneCountInString(req.Body) > 100 || len(req.Body) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "本文の長さは1~100文字である必要があります")))
		return
	}

	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "unauthorized")))
		return
	}

	err = pc.updatePostUsecase.Execute(ctx, application.UpdatePostUsecaseInput{
		PostID: postID,
		UserID: user.ID,
		Title:  req.Title,
		Body:   req.Body,
	})
	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.JSON(200, schema.MutationSchema{
		TargetID: postID,
		Message:  "投稿の更新に成功しました",
	})
}

// delete psot
func (pc PostController) DeletePost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}
	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "unauthorized")))
		return
	}

	err = pc.deletePostUsecase.Execute(ctx, application.DeletePostUsecaseInput{
		PostID: postID,
		UserID: user.ID,
	})

	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	// 削除は204なので、204を返す
	c.JSON(204, schema.MutationSchema{
		TargetID: postID,
		Message:  "投稿を削除しました",
	})
}

func (pc PostController) ListComments(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	result, err := pc.listCommentsUsecase.Execute(ctx, application.ListCommentsUsecaseInput{
		PostID: postID,
	})
	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
	}

	resp := make([]schema.CommentResponse, len(result.Comments))
	for i, com := range result.Comments {
		resp[i] = schema.CommentResponse{
			ID:     com.ID,
			PostID: com.PostID,
			Body:   com.Body,
			UserResponse: schema.UserResponse{
				ID:       com.UserID,
				UserName: com.UserName,
			},
			CreatedAt: com.CreatedAt,
			UpdatedAt: com.UpdatedAt,
		}
	}
	c.JSON(200, resp)
}

func (p PostController) CreateComment(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	var req schema.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "リクエストの形式が誤っています")))
		return
	}
	// TODO: そのうちバリデーションライブラリ導入したい
	// RUNEを使っているのは、日本語の文字数をちゃんと正しく取るため
	if utf8.RuneCountInString(req.Body) > 100 || len(req.Body) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "コメントの長さは1~100文字である必要があります")))
		return
	}
	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "unauthorized")))
		return
	}
	result, err := p.createCommentUsecase.Execute(ctx, application.CreateCommentUsecaseInput{
		UserID: user.ID,
		PostID: postID,
		Body:   req.Body,
	})
	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.JSON(201, schema.MutationSchema{
		// TODO: IDをとってくる
		TargetID: result.CommentID,
		Message:  "新しいコメントを投稿しました",
	})
}

func (p PostController) UpdateComment(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
		return
	}

	var req schema.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "リクエストの形式が誤っています")))
		return
	}
	// TODO: そのうちバリデーションライブラリ導入したい
	// RUNEを使っているのは、日本語の文字数をちゃんと正しく取るため
	if utf8.RuneCountInString(req.Body) > 100 || len(req.Body) == 0 {
		c.JSON(400, schema.NewErrorResponse(apperror.New(apperror.CodeInvalidArgument, "コメントの長さは1~100文字である必要があります")))
		return
	}
	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "unauthorized")))
		return
	}

	err = p.updateCommentUsecase.Execute(ctx, application.UpdateCommentUsecaseInput{
		UserID:    user.ID,
		PostID:    postID,
		CommentID: commentID,
		Body:      req.Body,
	})
	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.JSON(201, schema.MutationSchema{
		// TODO: IDをとってくる
		TargetID: commentID,
		Message:  "コメントを更新しました",
	})
}

func (p PostController) DeleteComment(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	user, ok := middleware.GetUserAuthContext(c)
	if !ok {
		c.JSON(401, schema.NewErrorResponse(apperror.New(apperror.CodeUnauthorized, "unauthorized")))
		return
	}
	log.Printf("user: %v", user)

	err = p.deleteCommentUsecase.Execute(ctx, application.DeleteCommentUsecaseInput{
		CommentID: commentID,
	})

	if apperror.Is(err, apperror.CodeForbidden) {
		c.JSON(403, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) || err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	// 削除は204なので、204を返す
	c.JSON(204, schema.MutationSchema{
		TargetID: commentID,
		Message:  "コメントを削除しました",
	})
}
