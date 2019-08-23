package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-xo-game/api"
)

func main(){
	host := flag.String("host", "3000", "start api host port")
	flag.Parse()
	route := gin.Default()
	route.POST("/new_game",api.NewGameHandler)
	route.POST("/play", api.PlayHandler)
	route.Run(":"+*host)
}
