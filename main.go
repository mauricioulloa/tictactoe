package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"tictactoe/game"
)

const (
	modeTwoPlayer = 1
	modeVsAI      = 2
)

func main() {
	board := game.NewBoard()
	renderer := game.NewRenderer()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Show menu and get game mode
		renderer.Clear()
		renderer.RenderMenu()

		if !scanner.Scan() {
			return
		}
		input := strings.TrimSpace(scanner.Text())

		// Handle quit from menu
		if input == "q" || input == "quit" {
			renderer.RenderGoodbye()
			return
		}

		mode, err := strconv.Atoi(input)
		if err != nil || (mode != modeTwoPlayer && mode != modeVsAI) {
			continue // Invalid input, show menu again
		}

		// Play the game
		board.Reset()
		if mode == modeTwoPlayer {
			playTwoPlayer(board, renderer, scanner)
		} else {
			playVsAI(board, renderer, scanner)
		}
	}
}

func playTwoPlayer(board *game.Board, renderer *game.Renderer, scanner *bufio.Scanner) {
	for {
		// Main game loop
		for !board.IsGameOver() {
			renderer.Clear()
			renderer.RenderBoard(board)
			renderer.RenderPrompt(board.CurrentPlayer())

			// Get player input
			if !scanner.Scan() {
				return
			}
			input := strings.TrimSpace(scanner.Text())

			// Handle quit
			if input == "q" || input == "quit" {
				return
			}

			// Parse move
			pos, err := strconv.Atoi(input)
			if err != nil {
				renderer.Clear()
				renderer.RenderBoard(board)
				renderer.RenderError("Please enter a number between 1-9")
				continue
			}

			// Try to make the move
			if !board.MakeMove(pos) {
				renderer.Clear()
				renderer.RenderBoard(board)
				if pos < 1 || pos > 9 {
					renderer.RenderError("Invalid position! Use 1-9")
				} else {
					renderer.RenderError("That cell is already taken!")
				}
				continue
			}
		}

		// Game over - show final state
		renderer.Clear()
		renderer.RenderBoard(board)

		if board.Winner() != game.Empty {
			renderer.RenderWinner(board.Winner())
		} else {
			renderer.RenderDraw()
		}

		// Ask to play again
		if !askPlayAgain(renderer, scanner) {
			return
		}
		board.Reset()
	}
}

func playVsAI(board *game.Board, renderer *game.Renderer, scanner *bufio.Scanner) {
	ai := game.NewAI(game.O, game.X) // AI plays as O, human plays as X

	for {
		// Main game loop
		for !board.IsGameOver() {
			renderer.Clear()
			renderer.RenderBoard(board)

			if board.CurrentPlayer() == game.X {
				// Human's turn
				renderer.RenderPrompt(game.X)

				if !scanner.Scan() {
					return
				}
				input := strings.TrimSpace(scanner.Text())

				// Handle quit
				if input == "q" || input == "quit" {
					return
				}

				// Parse move
				pos, err := strconv.Atoi(input)
				if err != nil {
					renderer.Clear()
					renderer.RenderBoard(board)
					renderer.RenderError("Please enter a number between 1-9")
					continue
				}

				// Try to make the move
				if !board.MakeMove(pos) {
					renderer.Clear()
					renderer.RenderBoard(board)
					if pos < 1 || pos > 9 {
						renderer.RenderError("Invalid position! Use 1-9")
					} else {
						renderer.RenderError("That cell is already taken!")
					}
					continue
				}
			} else {
				// AI's turn
				renderer.RenderAIThinking()
				time.Sleep(500 * time.Millisecond) // Small delay for better UX

				pos := ai.GetBestMove(board)
				board.MakeMove(pos)

				renderer.Clear()
				renderer.RenderBoard(board)
				renderer.RenderAIMove(pos)
				time.Sleep(300 * time.Millisecond) // Let player see the move
			}
		}

		// Game over - show final state
		renderer.Clear()
		renderer.RenderBoard(board)

		if board.Winner() != game.Empty {
			isHumanWinner := board.Winner() == game.X
			renderer.RenderWinnerVsAI(board.Winner(), isHumanWinner)
		} else {
			renderer.RenderDraw()
		}

		// Ask to play again
		if !askPlayAgain(renderer, scanner) {
			return
		}
		board.Reset()
	}
}

func askPlayAgain(renderer *game.Renderer, scanner *bufio.Scanner) bool {
	renderer.RenderPlayAgain()
	if !scanner.Scan() {
		return false
	}
	response := strings.ToLower(strings.TrimSpace(scanner.Text()))
	return response == "y" || response == "yes"
}
