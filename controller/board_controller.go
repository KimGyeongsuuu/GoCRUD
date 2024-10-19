package controller

import (
	"context"
	"net/http"

	"go-practice/model"
	"go-practice/model/data/input"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	boardUseCase model.BoardUseCase
}

func NewBoardController(boardUseCase model.BoardUseCase) *BoardController {
	return &BoardController{
		boardUseCase: boardUseCase,
	}
}

func (c *BoardController) CreateBoard(ctx *gin.Context) {
	var input input.CreateBoardInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.boardUseCase.CreateBoard(context.Background(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Board created successfully"})
}

func (c *BoardController) GetBoard(ctx *gin.Context) {

	boards, err := c.boardUseCase.GetBoard(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve boards"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"boards": boards})

}
