package api

import (
	"github.com/gin-gonic/gin"
	"go-xo-game/xo"
	"net/http"
)

var xoGame xo.Game

type PlayGame struct {
	Player xo.Player `json:"player"`
	LocationX int `json:"location_x"`
	LocationY int `json:"location_y"`
}

func NewGameHandler(context *gin.Context){
	var reqXOGame xo.Game
	err := context.BindJSON(&reqXOGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	playerOne := xo.NewPlayer(reqXOGame.PlayersOne.Name,reqXOGame.PlayersOne.Symbol)
	playerTwo := xo.NewPlayer(reqXOGame.PlayersTwo.Name,reqXOGame.PlayersTwo.Symbol)
	xoGame = xo.NewGame(playerOne,playerTwo)
	context.JSON(http.StatusCreated, xoGame)
}

func PlayHandler(context *gin.Context){
	var reqPlayGame PlayGame
	err := context.BindJSON(&reqPlayGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":xoGame.Play(reqPlayGame.Player,reqPlayGame.LocationX,reqPlayGame.LocationY),
	})
}