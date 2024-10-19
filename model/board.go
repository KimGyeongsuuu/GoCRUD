package model

import (
	"context"
	"go-practice/model/data/input"
)

type Board struct {
	BoardID uint64 `gorm:"primaryKey;autoIncrement"`
	Title   string
	Content string
}

type BoardUseCase interface {
	CreateBoard(ctx context.Context, input *input.CreateBoardInput) error
	GetBoard(ctx context.Context) ([]Board, error)
	DeleteBoard(ctx context.Context, boardID uint64) error
}

type BoardRepository interface {
	CreateBoard(ctx context.Context, board *Board) error
	FindAllBoard(ctx context.Context) ([]Board, error)
	DeleteBoard(ctx context.Context, boardId uint64) error
}
