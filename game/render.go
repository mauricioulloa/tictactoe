package game

import (
	"fmt"
	"strings"
)

// ANSI color codes - Atipicus branding
const (
	colorReset = "\033[0m"
	colorBold  = "\033[1m"
	colorDim   = "\033[2m"

	// Atipicus brand colors (24-bit RGB)
	colorMint     = "\033[38;2;152;216;200m" // Mint/Cyan accent - Player X
	colorLavender = "\033[38;2;168;155;189m" // Lavender/Purple - Player O
	colorWhite    = "\033[38;2;255;255;255m" // Pure white - Winners
	colorMuted    = "\033[38;2;100;100;110m" // Muted gray - Position numbers

	// Aliases for compatibility
	colorCyan    = colorMint
	colorMagenta = colorLavender
	colorYellow  = "\033[38;2;255;200;100m" // Warm warning color
	colorGreen   = colorWhite               // Use white for winning highlight

	clearScreen = "\033[2J\033[H"
)

// Renderer handles all terminal output
type Renderer struct{}

// NewRenderer creates a new renderer
func NewRenderer() *Renderer {
	return &Renderer{}
}

// Clear clears the terminal screen
func (r *Renderer) Clear() {
	fmt.Print(clearScreen)
}

// RenderBoard displays the game board with colors
func (r *Renderer) RenderBoard(board *Board) {
	winningCells := board.GetWinningCells()
	isWinningCell := make(map[int]bool)
	for _, idx := range winningCells {
		isWinningCell[idx] = true
	}

	fmt.Println()
	fmt.Printf("   %s%sTic-Tac-Toe%s\n", colorBold, colorWhite, colorReset)
	fmt.Printf("   %sby Atipicus%s\n", colorLavender, colorReset)
	fmt.Println()

	for row := 0; row < 3; row++ {
		fmt.Print("   ")
		for col := 0; col < 3; col++ {
			idx := row*3 + col
			cell := board.GetCell(idx)

			// Add separator
			if col > 0 {
				fmt.Print(" │ ")
			} else {
				fmt.Print(" ")
			}

			// Render cell content
			r.renderCell(cell, idx, isWinningCell[idx])
		}
		fmt.Println()

		// Draw horizontal line between rows
		if row < 2 {
			fmt.Println("   ───┼───┼───")
		}
	}
	fmt.Println()
}

// renderCell renders a single cell with appropriate coloring
func (r *Renderer) renderCell(cell Cell, idx int, isWinning bool) {
	switch cell {
	case X:
		if isWinning {
			fmt.Print(colorBold + colorGreen + "X" + colorReset)
		} else {
			fmt.Print(colorCyan + "X" + colorReset)
		}
	case O:
		if isWinning {
			fmt.Print(colorBold + colorGreen + "O" + colorReset)
		} else {
			fmt.Print(colorMagenta + "O" + colorReset)
		}
	default:
		// Show position number for empty cells (1-9)
		fmt.Print(colorMuted + fmt.Sprintf("%d", idx+1) + colorReset)
	}
}

// RenderPrompt shows the current player's turn
func (r *Renderer) RenderPrompt(player Cell) {
	playerColor := colorCyan
	if player == O {
		playerColor = colorMagenta
	}
	fmt.Printf("  Player %s%c%s's turn\n", playerColor, player, colorReset)
	fmt.Print("  Enter position (1-9): ")
}

// RenderError shows an error message
func (r *Renderer) RenderError(msg string) {
	fmt.Printf("  %s%s%s\n\n", colorYellow, msg, colorReset)
}

// RenderWinner announces the winner
func (r *Renderer) RenderWinner(winner Cell) {
	playerColor := colorCyan
	if winner == O {
		playerColor = colorMagenta
	}
	fmt.Println(strings.Repeat("─", 30))
	fmt.Printf("  🎉 %sPlayer %c wins!%s 🎉\n", colorBold+playerColor, winner, colorReset)
	fmt.Println(strings.Repeat("─", 30))
	fmt.Println()
}

// RenderDraw announces a draw
func (r *Renderer) RenderDraw() {
	fmt.Println(strings.Repeat("─", 30))
	fmt.Printf("  %sIt's a draw!%s\n", colorYellow, colorReset)
	fmt.Println(strings.Repeat("─", 30))
	fmt.Println()
}

// RenderPlayAgain prompts for another game
func (r *Renderer) RenderPlayAgain() {
	fmt.Print("  Play again? (y/n): ")
}

// RenderGoodbye shows exit message
func (r *Renderer) RenderGoodbye() {
	fmt.Println()
	fmt.Println("  Thanks for playing! Goodbye!")
	fmt.Println()
}

// RenderMenu displays the game mode selection menu
func (r *Renderer) RenderMenu() {
	fmt.Println()
	fmt.Printf("   %s%sTic-Tac-Toe%s\n", colorBold, colorWhite, colorReset)
	fmt.Printf("   %sby Atipicus%s\n", colorLavender, colorReset)
	fmt.Println()
	fmt.Println("   Select game mode:")
	fmt.Println()
	fmt.Printf("   %s1%s. Two Players\n", colorMint, colorReset)
	fmt.Printf("   %s2%s. vs Computer\n", colorLavender, colorReset)
	fmt.Println()
	fmt.Print("  Enter choice (1 or 2): ")
}

// RenderAIThinking shows a message while AI computes
func (r *Renderer) RenderAIThinking() {
	fmt.Printf("  %sComputer is thinking...%s\n", colorDim, colorReset)
}

// RenderAIMove shows the AI's chosen move
func (r *Renderer) RenderAIMove(pos int) {
	fmt.Printf("  Computer plays position %s%d%s\n", colorMagenta, pos, colorReset)
}

// RenderWinnerVsAI announces the winner in single player mode
func (r *Renderer) RenderWinnerVsAI(winner Cell, isHuman bool) {
	fmt.Println(strings.Repeat("─", 30))
	if isHuman {
		fmt.Printf("  🎉 %sYou win!%s 🎉\n", colorBold+colorGreen, colorReset)
	} else {
		fmt.Printf("  %sComputer wins!%s\n", colorBold+colorMagenta, colorReset)
	}
	fmt.Println(strings.Repeat("─", 30))
	fmt.Println()
}
