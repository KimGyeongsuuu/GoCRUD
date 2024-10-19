package repository

import (
	"context"
	"go-practice/model"

	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{
		db: db,
	}
}

func (r *BoardRepository) CreateBoard(ctx context.Context, board *model.Board) error {
	result := r.db.WithContext(ctx).Create(board)
	return result.Error
}

func (r *BoardRepository) FindAllBoard(ctx context.Context) ([]model.Board, error) {
	var boards []model.Board
	result := r.db.WithContext(ctx).Find(&boards)
	return boards, result.Error
}
