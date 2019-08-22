package xo

type Player struct {
	symbol string
	name string
	score int
}

type Game struct {
	board [3][3]string
	playersOne Player
	playersTwo Player
	currentTurn string
}

func NewPlayer(name, symbol string) Player {
	return Player{
		symbol: symbol,
		name:   name,
		score:  0,
	}
}

func NewGame(playerOne , playerTwo Player) Game {
	return Game{
		board:       [3][3]string{},
		playersOne:  playerOne,
		playersTwo:  playerTwo,
		currentTurn: playerOne.name,
	}
}

func (game *Game) Play(player Player, locationX,locationY int) string {
	game.marking(player,locationX,locationY)
	gameDone , winner := game.checkWin()
	if gameDone {
		return winner
	}
	game.switchTurn()
	return "NO WIN"
}

func (game *Game) marking(player Player, locationX,locationY int) {
	game.board[locationX][locationY] = player.symbol
}

func (game Game) checkWin() (bool , string) {
	if (
			game.board[0][0] == game.playersOne.symbol &&
			game.board[0][1] == game.playersOne.symbol &&
			game.board[0][2] == game.playersOne.symbol ) {
			return true , "X WIN"
	}
	if (
			game.board[1][0] == game.playersOne.symbol &&
			game.board[1][1] == game.playersOne.symbol &&
			game.board[1][2] == game.playersOne.symbol ) {
		return true , "X WIN"
	}
	return false , "NO WIN"
}

func (game *Game) switchTurn() {
	if game.currentTurn == game.playersOne.name {
		game.currentTurn = game.playersTwo.name
	} else {
		game.currentTurn = game.playersOne.name
	}
}