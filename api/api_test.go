package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CreateNewXOGame_Input_PlayerOne_KA_PlayerTwo_PK_Should_Be_New_XOGame_With_Two_Players_Empty_Board_And_No_Score(t *testing.T){
	expected := `{"board":[["","",""],["","",""],["","",""]],"player_one":{"symbol":"X","name":"KA","score":0},"player_two":{"symbol":"O","name":"PK","score":0},"current_turn":"KA"}`
	requestBody := `{"player_one":{"symbol":"X","name":"KA"},"player_two":{"symbol":"O","name":"PK"}}`
	request := httptest.NewRequest("POST","/new_game",strings.NewReader(requestBody))
	writer := httptest.NewRecorder()
	testRoute := gin.Default()
	testRoute.POST("/new_game", NewGameHandler)
	testRoute.ServeHTTP(writer,request)
	response := writer.Result()

	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}

