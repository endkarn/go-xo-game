package api

import "go-xo-game/xo"

type FakeGame struct {

}

func (fakeGame FakeGame) Play(player Player, locationX, locationY int) string {
	return "X WIN"
}

func (fakeGame FakeGame) IsStarted() bool {
	return true
}

func (fakeGame FakeGame) GetGameDetail() xo2.Game {
	return xo2.Game{
		Board:       [3][3]string{
			[3]string{"X","O","X"},
		},
		PlayersOne:  xo2.Player{
			Symbol: "X",
			Name:   "KA",
			Score:  10,
		},
		PlayersTwo:  Player{
			Symbol: "O",
			Name:   "PK",
			Score:  5,
		},
		CurrentTurn: "KA",
	}
}
