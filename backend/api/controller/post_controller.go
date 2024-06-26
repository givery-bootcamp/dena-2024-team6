package controller

import (
	"context"
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
	listPostUsecase      application.ListPostUsecase
	getPostUsecase       application.GetPostDetailUsecase
	listCommentsUsecase  application.ListCommentsUsecase
	createCommentUsecase application.CreateCommentUsecase
}

func NewPostController(i *do.Injector) (*PostController, error) {
	listPostUsecase := do.MustInvoke[application.ListPostUsecase](i)
	getPostUsecase := do.MustInvoke[application.GetPostDetailUsecase](i)
	listCommentsUsecase := do.MustInvoke[application.ListCommentsUsecase](i)
	createCommentUsecase := do.MustInvoke[application.CreateCommentUsecase](i)
	return &PostController{
		listPostUsecase:      listPostUsecase,
		getPostUsecase:       getPostUsecase,
		listCommentsUsecase:  listCommentsUsecase,
		createCommentUsecase: createCommentUsecase,
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
			ID:    p.ID,
			Title: p.Title,
			UserResponse: schema.UserResponse{
				ID:       p.UserID,
				UserName: p.UserName,
			},
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

	c.JSON(200, schema.PostResponse{
		ID:    result.Post.ID,
		Title: result.Post.Title,
		Body:  result.Post.Body,
		UserResponse: schema.UserResponse{
			ID:       result.Post.UserID,
			UserName: result.Post.UserName,
		},
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

	result, err := p.createCommentUsecase.Execute(ctx, application.CreateCommentUsecaseInput{
		UserID: 1,
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
