package game

import "math"

// AI represents the computer opponent
type AI struct {
	player   Cell // The AI's mark (usually O)
	opponent Cell // The human's mark (usually X)
}

// NewAI creates a new AI player
func NewAI(aiPlayer, humanPlayer Cell) *AI {
	return &AI{
		player:   aiPlayer,
		opponent: humanPlayer,
	}
}

// GetBestMove returns the optimal move (1-9) using minimax with alpha-beta pruning
func (ai *AI) GetBestMove(board *Board) int {
	// Clone the board for simulation
	simBoard := board.Clone()

	bestScore := math.MinInt32
	bestMove := -1

	// Try each empty cell
	for _, pos := range simBoard.GetEmptyCells() {
		// Make the move on simulation board
		simBoard.simulateMove(pos, ai.player)

		// Get score for this move
		score := ai.minimax(simBoard, 0, false, math.MinInt32, math.MaxInt32)

		// Undo the move
		simBoard.undoMove(pos)

		// Update best move
		if score > bestScore {
			bestScore = score
			bestMove = pos
		}
	}

	// Convert from 0-8 index to 1-9 position
	return bestMove + 1
}

// minimax implements the minimax algorithm with alpha-beta pruning
// isMaximizing: true when it's AI's turn, false when it's opponent's turn
// alpha: best score the maximizer can guarantee
// beta: best score the minimizer can guarantee
func (ai *AI) minimax(board *Board, depth int, isMaximizing bool, alpha, beta int) int {
	// Check terminal states
	winner := ai.checkWinnerFor(board)
	if winner == ai.player {
		return 10 - depth // Prefer faster wins
	}
	if winner == ai.opponent {
		return depth - 10 // Prefer slower losses
	}
	if len(board.GetEmptyCells()) == 0 {
		return 0 // Draw
	}

	if isMaximizing {
		// AI's turn - maximize score
		maxScore := math.MinInt32
		for _, pos := range board.GetEmptyCells() {
			board.simulateMove(pos, ai.player)
			score := ai.minimax(board, depth+1, false, alpha, beta)
			board.undoMove(pos)

			maxScore = max(maxScore, score)
			alpha = max(alpha, score)
			if beta <= alpha {
				break // Beta cutoff
			}
		}
		return maxScore
	} else {
		// Opponent's turn - minimize score
		minScore := math.MaxInt32
		for _, pos := range board.GetEmptyCells() {
			board.simulateMove(pos, ai.opponent)
			score := ai.minimax(board, depth+1, true, alpha, beta)
			board.undoMove(pos)

			minScore = min(minScore, score)
			beta = min(beta, score)
			if beta <= alpha {
				break // Alpha cutoff
			}
		}
		return minScore
	}
}

// checkWinnerFor checks if there's a winner on the board
func (ai *AI) checkWinnerFor(board *Board) Cell {
	winPatterns := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, pattern := range winPatterns {
		a, b, c := board.cells[pattern[0]], board.cells[pattern[1]], board.cells[pattern[2]]
		if a != Empty && a == b && b == c {
			return a
		}
	}
	return Empty
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
