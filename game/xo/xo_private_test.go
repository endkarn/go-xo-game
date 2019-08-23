package xo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marking_Input_PlayerOne_LocationX_0_LocationY_0_Should_Be_X_In_Board_0_0(t *testing.T) {
	expected := [3][3]string{
		[3]string{"X", "", ""},
		[3]string{"", "", ""},
		[3]string{"", "", ""},
	}
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.marking(playerOne, 0, 0)
	actual := game.Board

	assert.Equal(t, expected, actual)
}

func Test_Marking_Input_PlayerOne_LocationX_0_LocationY_0_And_PlayerTwo_LocationX_0_LocationY_1_Should_Be_X_In_Board_0_0_And_O_In_Board_0_1(t *testing.T) {
	expected := [3][3]string{
		[3]string{"X", "O", ""},
		[3]string{"", "", ""},
		[3]string{"", "", ""},
	}
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.marking(playerOne, 0, 0)
	game.marking(playerTwo, 0, 1)
	actual := game.Board

	assert.Equal(t, expected, actual)
}

func Test_CheckWin_Put_XXX_In_First_Line_Should_Be_X_WIN_As_Winner_And_True_As_GameDone(t *testing.T) {
	expectedGameDone, expectedWinner := true, "X WIN"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)
	game.Board = [3][3]string{
		[3]string{"X", "X", "X"},
		[3]string{"", "", ""},
		[3]string{"", "", ""},
	}
	actualGameDone, actualWinner := game.checkWin()

	assert.Equal(t, expectedWinner, actualWinner)
	assert.Equal(t, expectedGameDone, actualGameDone)
}

func Test_CheckWin_Input_Board_From_NewGame_Should_Be_NO_WIN_As_Winner_And_False_As_GameDone(t *testing.T) {
	expectedGameDone, expectedWinner := false, "NO WIN"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	actualGameDone, actualWinner := game.checkWin()

	assert.Equal(t, expectedWinner, actualWinner)
	assert.Equal(t, expectedGameDone, actualGameDone)
}

func Test_SwitchTurn_Input_PlayerOne_Play_First_Turn_Should_Be_PlayerTwo_Turn(t *testing.T) {
	expected := "PK"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.Play(playerOne, 0, 0)
	actual := game.CurrentTurn

	assert.Equal(t, expected, actual)
}

func Test_SwitchTurn_Input_PlayerTwo_Play_Second_Turn_Should_Be_PlayerOne_Turn(t *testing.T) {
	expected := "KA"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.Play(playerOne, 0, 0)
	game.Play(playerTwo, 0, 1)
	actual := game.CurrentTurn

	assert.Equal(t, expected, actual)
}

func Test_GetCurrentPlayer_FirstTurn_Should_Be_PlayerOne_Turn(t *testing.T) {
	expected := "KA"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")

	game := NewGame(playerOne, playerTwo)
	actual := game.getCurrentPlayer().Name

	assert.Equal(t, expected, actual)
}

func Test_GetCurrentPlayer_SecondTurn_Should_Be_PlayerTwo_Turn(t *testing.T) {
	expected := "PK"
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")

	game := NewGame(playerOne, playerTwo)
	game.Play(playerOne, 0, 0)
	actual := game.getCurrentPlayer().Name

	assert.Equal(t, expected, actual)
}

func Test_XOGame_PlayerOne_Have_Score_1_After_Win_FirstGame(t *testing.T) {
	expected := 1
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.Play(playerOne, 0, 0)
	game.Play(playerTwo, 1, 0)
	game.Play(playerOne, 0, 1)
	game.Play(playerTwo, 1, 1)
	game.Play(playerOne, 0, 2)
	actual := game.PlayersOne.Score

	assert.Equal(t, expected, actual)
}

func Test_XOGame_PlayerTwo_Have_Score_1_After_Win_FirstGame(t *testing.T) {
	expected := 1
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.Play(playerOne, 0, 1)
	game.Play(playerTwo, 0, 0)
	game.Play(playerOne, 1, 0)
	game.Play(playerTwo, 1, 1)
	game.Play(playerOne, 2, 1)
	game.Play(playerTwo, 2, 2)
	actual := game.PlayersTwo.Score

	assert.Equal(t, expected, actual)
}

func Test_XOGame_GameBoard_Should_Be_Empty_After_Player_Win_Game(t *testing.T) {
	expected := [3][3]string{}
	playerOne := NewPlayer("KA", "X")
	playerTwo := NewPlayer("PK", "O")
	game := NewGame(playerOne, playerTwo)

	game.Play(playerOne, 0, 1)
	game.Play(playerTwo, 0, 0)
	game.Play(playerOne, 1, 0)
	game.Play(playerTwo, 1, 1)
	game.Play(playerOne, 2, 1)
	game.Play(playerTwo, 2, 2)
	actual := game.Board

	assert.Equal(t, expected, actual)
}
