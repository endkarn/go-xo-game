package xo

const nowinmessage = "NO WIN"

type Player struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

type Game struct {
	Board       [3][3]string `json:"board"`
	PlayersOne  Player       `json:"player_one"`
	PlayersTwo  Player       `json:"player_two"`
	CurrentTurn string       `json:"current_turn"`
}

func NewPlayer(name, symbol string) Player {
	return Player{
		Symbol: symbol,
		Name:   name,
		Score:  0,
	}
}

func NewGame(playerOne, playerTwo Player) Game {
	return Game{
		Board:       [3][3]string{},
		PlayersOne:  playerOne,
		PlayersTwo:  playerTwo,
		CurrentTurn: playerOne.Name,
	}
}

func (game *Game) Play(player Player, locationX, locationY int) string {
	game.marking(player, locationX, locationY)
	haveWinner, winner := game.checkWin()
	if haveWinner {
		game.getCurrentPlayer().Score++
		game.Board = [3][3]string{}
		return winner
	}
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, locationX, locationY int) {
	game.Board[locationX][locationY] = player.Symbol
}

func (game Game) checkWin() (bool, string) {
	currentPlayer := game.getCurrentPlayer()
	// Horizontal Conditions
	for x := range game.Board {
		if game.Board[x][0] == currentPlayer.Symbol &&
			game.Board[x][1] == currentPlayer.Symbol &&
			game.Board[x][2] == currentPlayer.Symbol {
			return true, currentPlayer.Symbol + " WIN"
		}
	}
	// Vertical Conditions
	for y := range game.Board {
		if game.Board[0][y] == currentPlayer.Symbol &&
			game.Board[1][y] == currentPlayer.Symbol &&
			game.Board[2][y] == currentPlayer.Symbol {
			return true, currentPlayer.Symbol + " WIN"
		}
	}
	// Cross Conditions
	if game.Board[0][0] == currentPlayer.Symbol &&
		game.Board[1][1] == currentPlayer.Symbol &&
		game.Board[2][2] == currentPlayer.Symbol {
		return true, currentPlayer.Symbol + " WIN"
	}
	if game.Board[0][2] == currentPlayer.Symbol &&
		game.Board[1][1] == currentPlayer.Symbol &&
		game.Board[2][0] == currentPlayer.Symbol {
		return true, currentPlayer.Symbol + " WIN"
	}
	return false, nowinmessage
}

func (game *Game) switchTurn() {
	switch game.CurrentTurn {
	case game.PlayersOne.Name:
		game.CurrentTurn = game.PlayersTwo.Name
	case game.PlayersTwo.Name:
		game.CurrentTurn = game.PlayersOne.Name
	}
}

func (game *Game) getCurrentPlayer() *Player {
	if game.CurrentTurn == game.PlayersOne.Name {
		return &game.PlayersOne
	}
	return &game.PlayersTwo
}
