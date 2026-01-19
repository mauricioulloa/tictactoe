# Tic-Tac-Toe

A terminal-based Tic-Tac-Toe game written in Go with a colorful UI.

## Features

- **Two-player mode** - Take turns with a friend on the same terminal
- **Single-player mode** - Play against an unbeatable AI opponent
- Colorful terminal interface with Unicode box-drawing characters
- Player X in cyan, Player O (or Computer) in magenta
- Winning cells highlighted in green
- Position numbers shown for empty cells
- Play again option

## The AI

The computer opponent uses the **minimax algorithm** with alpha-beta pruning. This makes it play optimally - the best you can achieve against it is a draw with perfect play. Good luck!

## Requirements

- Go 1.21 or later

## Build & Run

```bash
# Navigate to the project directory
cd tictactoe

# Run directly
go run .

# Or build and run
go build -o tictactoe
./tictactoe
```

## How to Play

1. Select a game mode:
   - `1` - Two Players
   - `2` - vs Computer

2. The board shows positions 1-9 for empty cells:
   ```
      1 │ 2 │ 3
     ───┼───┼───
      4 │ 5 │ 6
     ───┼───┼───
      7 │ 8 │ 9
   ```

3. Player X goes first (that's you in single-player mode)
4. Enter a number (1-9) to place your mark
5. Players alternate turns until someone wins or it's a draw
6. Type `q` or `quit` to exit anytime

## Controls

- `1` / `2` - Select game mode
- `1-9` - Place mark at position
- `q` / `quit` - Exit game
- `y` / `n` - Play again prompt
