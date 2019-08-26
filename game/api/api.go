package api

import (
	"github.com/gin-gonic/gin"
	"go-xo-game/xo"
	"net/http"
)

type GameAPI struct {
	XOGame xo.OXGame
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
	game := xo.NewGame(playerOne, playerTwo)
	gameAPI.XOGame = &game

	context.JSON(http.StatusCreated, gameAPI.XOGame)
}

func (gameAPI *GameAPI) PlayHandler(context *gin.Context) {
	var reqPlayGame PlayGame
	err := context.BindJSON(&reqPlayGame)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !gameAPI.XOGame.IsStarted() {
		context.JSON(http.StatusNotFound, gin.H{"error": "game did not created yet , not fond the game"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": gameAPI.XOGame.Play(reqPlayGame.Player, reqPlayGame.LocationX, reqPlayGame.LocationY),
	})
}

func (gameAPI *GameAPI) ViewGameHandler(context *gin.Context) {
	if !gameAPI.XOGame.IsStarted() {
		context.JSON(http.StatusNotFound, gin.H{"error": "game did not created yet , not fond the game"})
		return
	}
	context.JSON(http.StatusOK, gameAPI.XOGame.GetGameDetail())
}
