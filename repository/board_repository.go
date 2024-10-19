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
