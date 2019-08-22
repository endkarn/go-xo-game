package xo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_XOGame_PlayerOne_Win_In_Horizontal_FirstLine(t *testing.T) {
	expected := "X WIN"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	game.Play(playerOne,0,0)
	game.Play(playerTwo,1,0)
	game.Play(playerOne,0,1)
	game.Play(playerTwo,1,1)
	actual := game.Play(playerOne,0,2)

	assert.Equal(t,expected,actual)
}

func Test_Marking_Input_PlayerOne_LocationX_0_LocationY_0_Should_Be_X_In_Board_0_0(t *testing.T){
	expected := [3][3]string{
		[3]string{"X","",""},
		[3]string{"","",""},
		[3]string{"","",""},
	}
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	game.marking(playerOne,0,0)
	actual := game.board

	assert.Equal(t,expected,actual)
}