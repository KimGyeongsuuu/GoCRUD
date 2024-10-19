package service

import (
	"context"
	"go-practice/model"
	"go-practice/model/data/input"
)

type BoardService struct {
	boardRepo model.BoardRepository
}

func NewBoardService(boardRepo model.BoardRepository) model.BoardUseCase {
	return &BoardService{
		boardRepo: boardRepo,
	}
}

func (u *BoardService) CreateBoard(ctx context.Context, input *input.CreateBoardInput) error {

	board := &model.Board{
		Title:   input.Title,
		Content: input.Content,
	}

	return u.boardRepo.CreateBoard(ctx, board)

}
