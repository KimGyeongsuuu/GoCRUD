package model

import (
	"context"
	"go-practice/model/data/input"
)

type Board struct {
	BoardID int
	Title   string
	Content string
}

type BoardUseCase interface {
	CreateBoard(ctx context.Context, input *input.CreateBoardInput) error
}

type BoardRepository interface {
	CreateBoard(ctx context.Context, board *Board) error
}
