package api

import (
	"github.com/gin-gonic/gin"
	"go-xo-game/xo"
	"net/http"
)

func NewGameHandler(context *gin.Context){
	playerOne := xo.NewPlayer("KA","X")
	playerTwo := xo.NewPlayer("PK","O")
	xoGame := xo.NewGame(playerOne,playerTwo)
	context.JSON(http.StatusCreated, xoGame)
}
