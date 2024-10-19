package model

import (
	"context"
	"go-practice/model/data/input"
)

type Board struct {
	BoardID int `gorm:"primaryKey;autoIncrement"`
	Title   string
	Content string
}

type BoardUseCase interface {
	CreateBoard(ctx context.Context, input *input.CreateBoardInput) error
	GetBoard(ctx context.Context) ([]Board, error)
}

type BoardRepository interface {
	CreateBoard(ctx context.Context, board *Board) error
	FindAllBoard(ctx context.Context) ([]Board, error)
}
