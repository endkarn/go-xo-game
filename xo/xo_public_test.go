package xo_test

import (
	"github.com/stretchr/testify/assert"
	"go-xo-game/xo"
	"testing"
)

func Test_XOGame_PlayerOne_Win_In_Horizontal_FirstLine(t *testing.T) {
	expected := "X WIN"
	playerOne := xo.NewPlayer("KA","X")
	playerTwo := xo.NewPlayer("PK","O")
	game := xo.NewGame(playerOne,playerTwo)

	game.Play(playerOne,0,0)
	game.Play(playerTwo,1,0)
	game.Play(playerOne,0,1)
	game.Play(playerTwo,1,1)
	actual := game.Play(playerOne,0,2)

	assert.Equal(t,expected,actual)
}

func Test_XOGame_PlayerOne_Win_In_Horizontal_SecondLine(t *testing.T) {
	expected := "X WIN"
	playerOne := xo.NewPlayer("KA","X")
	playerTwo := xo.NewPlayer("PK","O")
	game := xo.NewGame(playerOne,playerTwo)

	game.Play(playerOne,1,0)
	game.Play(playerTwo,0,0)
	game.Play(playerOne,1,1)
	game.Play(playerTwo,0,1)
	actual := game.Play(playerOne,1,2)

	assert.Equal(t,expected,actual)
}

func Test_XOGame_PlayerOne_Win_In_Horizontal_ThirdLine(t *testing.T) {
	expected := "X WIN"
	playerOne := xo.NewPlayer("KA","X")
	playerTwo := xo.NewPlayer("PK","O")
	game := xo.NewGame(playerOne,playerTwo)

	game.Play(playerOne,2,0)
	game.Play(playerTwo,0,0)
	game.Play(playerOne,2,1)
	game.Play(playerTwo,0,1)
	actual := game.Play(playerOne,2,2)

	assert.Equal(t,expected,actual)
}