package xo

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func Test_CheckWin_Input_XXX_In_First_Line_Should_Be_X_WIN(t *testing.T){
	expected := "X WIN"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)
	game.board = [3][3]string{
		[3]string{"X","X","X"},
		[3]string{"","",""},
		[3]string{"","",""},
	}
	_,actual := game.checkWin()

	assert.Equal(t,expected,actual)
}

func Test_CheckWin_Input_Board_From_NewGame_Should_Be_NO_WIN(t *testing.T){
	expected := "NO WIN"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	_,actual := game.checkWin()

	assert.Equal(t,expected,actual)
}

func Test_SwitchTurn_Input_PlayerOne_Play_First_Turn_Should_Be_PlayerTwo_Turn(t *testing.T){
	expected := "PK"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	game.Play(playerOne,0,0)
	actual := game.currentTurn

	assert.Equal(t,expected,actual)
}

func Test_SwitchTurn_Input_PlayerTwo_Play_Second_Turn_Should_Be_PlayerOne_Turn(t *testing.T){
	expected := "KA"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	game.Play(playerOne,0,0)
	game.Play(playerTwo,0,1)
	actual := game.currentTurn

	assert.Equal(t,expected,actual)
}

func Test_XOGame_PlayerOne_Win_In_Horizontal_SecondLine(t *testing.T) {
	expected := "X WIN"
	playerOne := NewPlayer("KA","X")
	playerTwo := NewPlayer("PK","O")
	game := NewGame(playerOne,playerTwo)

	game.Play(playerOne,1,0)
	game.Play(playerTwo,0,0)
	game.Play(playerOne,1,1)
	game.Play(playerTwo,0,1)
	actual := game.Play(playerOne,1,2)

	assert.Equal(t,expected,actual)
}