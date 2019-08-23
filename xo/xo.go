package xo

const nowinmessage = "NO WIN"

type Player struct {
	symbol string
	name   string
	score  int
}

type Game struct {
	board       [3][3]string
	playersOne  Player
	playersTwo  Player
	currentTurn string
}

func NewPlayer(name, symbol string) Player {
	return Player{
		symbol: symbol,
		name:   name,
		score:  0,
	}
}

func NewGame(playerOne, playerTwo Player) Game {
	return Game{
		board:       [3][3]string{},
		playersOne:  playerOne,
		playersTwo:  playerTwo,
		currentTurn: playerOne.name,
	}
}

func (game *Game) Play(player Player, locationX, locationY int) string {
	game.marking(player, locationX, locationY)
	haveWinner, winner := game.checkWin()
	if haveWinner {
		game.getCurrentPlayer().score++
		game.board = [3][3]string{}
		return winner
	}
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, locationX, locationY int) {
	game.board[locationX][locationY] = player.symbol
}

func (game Game) checkWin() (bool, string) {
	currentPlayer := game.getCurrentPlayer()
	// Horizontal Conditions
	if game.board[0][0] == currentPlayer.symbol &&
		game.board[0][1] == currentPlayer.symbol &&
		game.board[0][2] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	if game.board[1][0] == currentPlayer.symbol &&
		game.board[1][1] == currentPlayer.symbol &&
		game.board[1][2] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	if game.board[2][0] == currentPlayer.symbol &&
		game.board[2][1] == currentPlayer.symbol &&
		game.board[2][2] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	// Vertical Conditions
	if game.board[0][0] == currentPlayer.symbol &&
		game.board[1][0] == currentPlayer.symbol &&
		game.board[2][0] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	if game.board[0][1] == currentPlayer.symbol &&
		game.board[1][1] == currentPlayer.symbol &&
		game.board[2][1] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	if game.board[0][2] == currentPlayer.symbol &&
		game.board[1][2] == currentPlayer.symbol &&
		game.board[2][2] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	// Cross Conditions
	if game.board[0][0] == currentPlayer.symbol &&
		game.board[1][1] == currentPlayer.symbol &&
		game.board[2][2] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	if game.board[0][2] == currentPlayer.symbol &&
		game.board[1][1] == currentPlayer.symbol &&
		game.board[2][0] == currentPlayer.symbol {
		return true, currentPlayer.symbol + " WIN"
	}
	return false, nowinmessage
}

func (game *Game) switchTurn() {
	switch game.currentTurn {
	case game.playersOne.name:
		game.currentTurn = game.playersTwo.name
	case game.playersTwo.name:
		game.currentTurn = game.playersOne.name
	}
}

func (game *Game) getCurrentPlayer() *Player {
	if game.currentTurn == game.playersOne.name {
		return &game.playersOne
	}
	return &game.playersTwo
}
