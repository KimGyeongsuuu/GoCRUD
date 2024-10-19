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

func (service *BoardService) CreateBoard(ctx context.Context, input *input.BoardInput) error {

	board := &model.Board{
		Title:   input.Title,
		Content: input.Content,
	}

	return service.boardRepo.CreateBoard(ctx, board)

}

func (service *BoardService) GetBoard(ctx context.Context) ([]model.Board, error) {
	return service.boardRepo.FindAllBoard(ctx)
}

func (service *BoardService) DeleteBoard(ctx context.Context, boardId uint64) error {
	return service.boardRepo.DeleteBoard(ctx, boardId)
}

func (service *BoardService) UpdateBoard(ctx context.Context, boardId uint64, input *input.BoardInput) error {
	board := &model.Board{
		Title:   input.Title,
		Content: input.Content,
	}

	return service.boardRepo.UpdateBoard(ctx, boardId, board)
}
