package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-xo-game/xo"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CreateNewXOGame_Input_PlayerOne_KA_PlayerTwo_PK_Should_Be_New_XOGame_With_Two_Players_Empty_Board_And_No_Score(t *testing.T) {
	expected := `{"board":[["","",""],["","",""],["","",""]],"player_one":{"symbol":"X","name":"KA","score":0},"player_two":{"symbol":"O","name":"PK","score":0},"current_turn":"KA"}`
	requestBody := `{"player_one":{"symbol":"X","name":"KA"},"player_two":{"symbol":"O","name":"PK"}}`
	request := httptest.NewRequest("POST", "/new_game", strings.NewReader(requestBody))
	writer := httptest.NewRecorder()
	gameAPI := GameAPI{}
	testRoute := gin.Default()
	testRoute.POST("/new_game", gameAPI.NewGameHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}

func Test_ViewGameHandler_Should_Be_Get_Current_Stage_Of_Game_Board(t *testing.T) {
	expected := `{"board":[["X","O","X"],["","",""],["","",""]],"player_one":{"symbol":"X","name":"KA","score":10},"player_two":{"symbol":"O","name":"PK","score":5},"current_turn":"KA"}`
	request := httptest.NewRequest("GET", "/game", nil)
	writer := httptest.NewRecorder()
	mockGame := xo.Game{
		Board: [3][3]string{
			[3]string{"X", "O", "X"},
			[3]string{"", "", ""},
			[3]string{"", "", ""},
		},
		PlayersOne: xo.Player{
			Name:   "KA",
			Symbol: "X",
			Score:  10,
		},
		PlayersTwo: xo.Player{
			Name:   "PK",
			Symbol: "O",
			Score:  5,
		},
		CurrentTurn: "KA",
	}
	mockGameAPI := GameAPI{xoGame: mockGame}
	testRoute := gin.Default()
	testRoute.GET("/game", mockGameAPI.ViewGameHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}
