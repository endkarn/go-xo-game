package xo

import "testing"

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

	if expected != actual {
		t.Errorf("Expected %s but get %s",expected,actual)
	}
}