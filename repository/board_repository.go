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

func (r *BoardRepository) DeleteBoard(ctx context.Context, boardId uint64) error {
	var board model.Board
	result := r.db.WithContext(ctx).Delete(&board, boardId)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error

}

func (r *BoardRepository) UpdateBoard(ctx context.Context, boardId uint64, board *model.Board) error {
	var existingBoard model.Board

	if err := r.db.WithContext(ctx).First(&existingBoard, boardId).Error; err != nil {
		return err
	}

	existingBoard.Title = board.Title
	existingBoard.Content = board.Content

	result := r.db.WithContext(ctx).Save(&existingBoard)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error

}
