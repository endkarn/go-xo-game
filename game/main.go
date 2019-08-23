package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-xo-game/api"
)

func main() {
	host := flag.String("host", "3000", "start api host port")
	flag.Parse()
	route := gin.Default()
	gameAPI := api.GameAPI{}
	route.POST("/new_game", gameAPI.NewGameHandler)
	route.POST("/play", gameAPI.PlayHandler)
	route.GET("/game", gameAPI.ViewGameHandler)
	route.Run(":" + *host)
}
