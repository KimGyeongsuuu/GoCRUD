package router

import (
	"go-practice/controller"
	"go-practice/repository"
	"go-practice/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	boardRepo := repository.NewBoardRepository(db)
	boardUseCase := service.NewBoardService(boardRepo)
	boardController := controller.NewBoardController(boardUseCase)

	health := r.Group("/health")
	{
		health.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "go-server is running",
			})
		})
	}
	boards := r.Group("/boards")
	{
		boards.POST("", boardController.CreateBoard)
		boards.GET("", boardController.GetBoard)
	}

	return r
}
