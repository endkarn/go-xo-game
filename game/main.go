package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-xo-game/api"
	"go-xo-game/xo"
)

func main() {
	host := flag.String("host", "3000", "start api host port")
	flag.Parse()
	route := gin.Default()

	playerOne := xo.NewPlayer(reqXOGame.PlayersOne.Name, reqXOGame.PlayersOne.Symbol)
	playerTwo := xo.NewPlayer(reqXOGame.PlayersTwo.Name, reqXOGame.PlayersTwo.Symbol)
	game := xo.NewGame(playerOne, playerTwo)
	gameAPI := api.GameAPI{XOGame:&game}
	route.POST("/new_game", gameAPI.NewGameHandler)
	route.POST("/play", gameAPI.PlayHandler)
	route.GET("/game", gameAPI.ViewGameHandler)
	route.Run(":" + *host)
}
