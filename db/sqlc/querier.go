// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error)
	CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error)
	CreateQuestion(ctx context.Context, text string) (Question, error)
	DeleteAnswer(ctx context.Context, id int64) error
	DeleteBlog(ctx context.Context, id int64) error
	DeleteQuestion(ctx context.Context, id int64) error
	GetAnswer(ctx context.Context, id int64) (Answer, error)
	GetAnswerForUpdate(ctx context.Context, id int64) (Answer, error)
	GetBlog(ctx context.Context, id int64) (Blog, error)
	GetBlogForUpdate(ctx context.Context, id int64) (Blog, error)
	GetQuestion(ctx context.Context, id int64) (Question, error)
	GetQuestionForUpdate(ctx context.Context, id int64) (Question, error)
	ListAnswer(ctx context.Context, arg ListAnswerParams) ([]Answer, error)
	ListBlog(ctx context.Context, arg ListBlogParams) ([]Blog, error)
	ListQuestion(ctx context.Context, arg ListQuestionParams) ([]Question, error)
	UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) (Answer, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error)
	UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) (Question, error)
}

var _ Querier = (*Queries)(nil)