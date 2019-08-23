package api

import (
	"github.com/gin-gonic/gin"
	"go-xo-game/xo"
	"net/http"
)

type GameAPI struct {
	xoGame xo.Game
}

type PlayGame struct {
	Player    xo.Player `json:"player"`
	LocationX int       `json:"location_x"`
	LocationY int       `json:"location_y"`
}

func (gameAPI *GameAPI) NewGameHandler(context *gin.Context) {
	var reqXOGame xo.Game
	err := context.BindJSON(&reqXOGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	playerOne := xo.NewPlayer(reqXOGame.PlayersOne.Name, reqXOGame.PlayersOne.Symbol)
	playerTwo := xo.NewPlayer(reqXOGame.PlayersTwo.Name, reqXOGame.PlayersTwo.Symbol)
	gameAPI.xoGame = xo.NewGame(playerOne, playerTwo)
	context.JSON(http.StatusCreated, gameAPI.xoGame)
}

func (gameAPI *GameAPI) PlayHandler(context *gin.Context) {
	var reqPlayGame PlayGame
	err := context.BindJSON(&reqPlayGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !gameIsStarted(gameAPI.xoGame) {
		context.JSON(http.StatusNotFound, gin.H{"error": "game did not created yet , not fond the game"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": gameAPI.xoGame.Play(reqPlayGame.Player, reqPlayGame.LocationX, reqPlayGame.LocationY),
	})
}

func (gameAPI *GameAPI) ViewGameHandler(context *gin.Context) {
	if !gameIsStarted(gameAPI.xoGame) {
		context.JSON(http.StatusNotFound, gin.H{"error": "game did not created yet , not fond the game"})
		return
	}
	context.JSON(http.StatusOK, gameAPI.xoGame)
}

func gameIsStarted(game xo.Game) bool {
	if game.PlayersOne.Symbol == "" || game.PlayersTwo.Symbol == "" {
		return false
	}
	return true
}
