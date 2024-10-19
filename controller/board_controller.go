package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"go-practice/model"
	"go-practice/model/data/input"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var input input.BoardInput

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

func (c *BoardController) DeleteBoard(ctx *gin.Context) {
	boardIDParam := ctx.Param("boardID")
	boardID, err := strconv.ParseUint(boardIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	err = c.boardUseCase.DeleteBoard(context.Background(), boardID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
		}
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (c *BoardController) UpdateBoard(ctx *gin.Context) {
	var input input.BoardInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	boardIDParam := ctx.Param("boardID")
	boardID, err := strconv.ParseUint(boardIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	if err := c.boardUseCase.UpdateBoard(ctx, boardID, &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update board"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": ""})

}
