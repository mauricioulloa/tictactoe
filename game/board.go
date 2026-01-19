package game

// Cell represents a cell on the board
type Cell rune

const (
	Empty Cell = ' '
	X     Cell = 'X'
	O     Cell = 'O'
)

// Board represents the tic-tac-toe game state
type Board struct {
	cells         [9]Cell
	currentPlayer Cell
	winner        Cell
	gameOver      bool
}

// NewBoard creates a new game board
func NewBoard() *Board {
	b := &Board{
		currentPlayer: X,
	}
	b.Reset()
	return b
}

// Reset clears the board for a new game
func (b *Board) Reset() {
	for i := range b.cells {
		b.cells[i] = Empty
	}
	b.currentPlayer = X
	b.winner = Empty
	b.gameOver = false
}

// GetCell returns the cell at the given position (0-8)
func (b *Board) GetCell(pos int) Cell {
	if pos < 0 || pos > 8 {
		return Empty
	}
	return b.cells[pos]
}

// CurrentPlayer returns whose turn it is
func (b *Board) CurrentPlayer() Cell {
	return b.currentPlayer
}

// Winner returns the winner (Empty if no winner yet)
func (b *Board) Winner() Cell {
	return b.winner
}

// IsGameOver returns true if the game has ended
func (b *Board) IsGameOver() bool {
	return b.gameOver
}

// MakeMove attempts to place the current player's mark at position (1-9)
// Returns true if the move was valid
func (b *Board) MakeMove(pos int) bool {
	// Convert from 1-9 to 0-8 index
	idx := pos - 1

	// Validate position
	if idx < 0 || idx > 8 {
		return false
	}

	// Check if cell is empty
	if b.cells[idx] != Empty {
		return false
	}

	// Check if game is already over
	if b.gameOver {
		return false
	}

	// Make the move
	b.cells[idx] = b.currentPlayer

	// Check for winner
	if b.checkWinner() {
		b.winner = b.currentPlayer
		b.gameOver = true
	} else if b.isFull() {
		b.gameOver = true
	} else {
		// Switch player
		if b.currentPlayer == X {
			b.currentPlayer = O
		} else {
			b.currentPlayer = X
		}
	}

	return true
}

// IsFull returns true if all cells are filled
func (b *Board) isFull() bool {
	for _, cell := range b.cells {
		if cell == Empty {
			return false
		}
	}
	return true
}

// checkWinner checks if the current player has won
func (b *Board) checkWinner() bool {
	// All winning combinations (indices)
	winPatterns := [][3]int{
		{0, 1, 2}, // Top row
		{3, 4, 5}, // Middle row
		{6, 7, 8}, // Bottom row
		{0, 3, 6}, // Left column
		{1, 4, 7}, // Middle column
		{2, 5, 8}, // Right column
		{0, 4, 8}, // Diagonal top-left to bottom-right
		{2, 4, 6}, // Diagonal top-right to bottom-left
	}

	player := b.currentPlayer
	for _, pattern := range winPatterns {
		if b.cells[pattern[0]] == player &&
			b.cells[pattern[1]] == player &&
			b.cells[pattern[2]] == player {
			return true
		}
	}
	return false
}

// GetWinningCells returns the indices of winning cells (empty if no winner)
func (b *Board) GetWinningCells() []int {
	if b.winner == Empty {
		return nil
	}

	winPatterns := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, pattern := range winPatterns {
		if b.cells[pattern[0]] == b.winner &&
			b.cells[pattern[1]] == b.winner &&
			b.cells[pattern[2]] == b.winner {
			return []int{pattern[0], pattern[1], pattern[2]}
		}
	}
	return nil
}

// Clone creates a copy of the board for AI simulation
func (b *Board) Clone() *Board {
	return &Board{
		cells:         b.cells,
		currentPlayer: b.currentPlayer,
		winner:        b.winner,
		gameOver:      b.gameOver,
	}
}

// GetEmptyCells returns indices (0-8) of all empty cells
func (b *Board) GetEmptyCells() []int {
	var empty []int
	for i, cell := range b.cells {
		if cell == Empty {
			empty = append(empty, i)
		}
	}
	return empty
}

// simulateMove places a mark at position (0-8) without validation (for AI)
func (b *Board) simulateMove(pos int, player Cell) {
	b.cells[pos] = player
}

// undoMove removes a mark at position (0-8) (for AI)
func (b *Board) undoMove(pos int) {
	b.cells[pos] = Empty
}
