package api

import (
	"github.com/gin-gonic/gin"
	"go-xo-game/xo"
	"net/http"
)

func NewGameHandler(context *gin.Context){
	var reqXOGame xo.Game
	err := context.BindJSON(&reqXOGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	playerOne := xo.NewPlayer(reqXOGame.PlayersOne.Name,reqXOGame.PlayersOne.Symbol)
	playerTwo := xo.NewPlayer(reqXOGame.PlayersTwo.Name,reqXOGame.PlayersTwo.Symbol)
	xoGame := xo.NewGame(playerOne,playerTwo)
	context.JSON(http.StatusCreated, xoGame)
}
