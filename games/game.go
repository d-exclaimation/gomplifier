package games

import (
	"fmt"
	"math/rand"
)

var (
	fox    = "🦊"
	target = "🐍"
	bomb   = "💣"
	dead   = "🔥"
	empty  = "➖"
)

// Sokoban Structure
type Sokoban struct {
	state   [][]string
	player  Vector2
	targets []Vector2
	items   []Vector2
	IsWin   bool
}

// Constructor
func NewSokoban(width int, height int) *Sokoban {
	var game = &Sokoban{
		state:   nil,
		player:  Vector2{
			x: 0,
			y: 0,
		},
		targets: nil,
		items:   nil,
		IsWin:   false,
	}

	var numObj = rand.Intn(2) + 3

	// Setup targets and item keys
	for i := 0; i < numObj; i++ {
		var (
			target = Vector2{rand.Intn(width), rand.Intn(height)}
			item = Vector2{rand.Intn(width-2) + 1, rand.Intn(height-2) + 1}
		)

		game.targets = append(game.targets, target)
		for IsIn(game.items, item) || IsIn(game.targets, item) {
			item = Vector2{rand.Intn(width-2) + 1, rand.Intn(height-2) + 1}
		}
		game.items = append(game.items, item)
	}

	// Instantiate the player
	var player = Vector2{rand.Intn(width), rand.Intn(height)}
	for IsIn(game.items, player) || IsIn(game.targets, player) {
		player = Vector2{rand.Intn(width), rand.Intn(height)}
	}
	game.player = player
	game.state = fillGrid(width, height, player, game.targets, game.items)

	return game
}

// Methods
func (game *Sokoban) Move(direction string) {
	if game.IsWin {
		return
	}

	var forecast = game.player.ApplyMovement(DirectionKeyCode[direction])
	fmt.Println(fmt.Sprintf("%d, %d", forecast.x, forecast.y))

	// Boundaries
	if forecast.y < 0 || forecast.y > len(game.state)-1 {
		return
	}
	if forecast.x < 0 || forecast.x > len(game.state[forecast.y])-1 {
		return
	}

	// Check next item in the forecast
	var nextSpot = game.state[forecast.y][forecast.x]

	// Not stack on top finished target
	if nextSpot == dead || nextSpot == target {
		return
	}

	if nextSpot == bomb {
		var next = game.MoveObject(forecast, direction)
		if !next {
			return
		}
	}

	// Switch the new location with the previous, basically move the character
	var temp = game.state[game.player.y][game.player.x]
	game.state[game.player.y][game.player.x] = empty
	game.state[forecast.y][forecast.x] = temp

	// Update the coordinates
	game.player = Vector2{forecast.x, forecast.y}

	game.IsWin = len(game.targets) == 0
}

func (game *Sokoban) MoveObject(location Vector2, direction string) bool {
	var forecast = location.ApplyMovement(DirectionKeyCode[direction])

	// Boundaries
	if forecast.y < 0 || forecast.y > len(game.state)-1 {
		return false
	}
	if forecast.x < 0 || forecast.x > len(game.state[forecast.y])-1 {
		return false
	}

	// Check next item in the forecast
	var (
		nextSpot = game.state[forecast.y][forecast.x]
		object = bomb
	)
	fmt.Println(fmt.Sprintf("%d, %d", forecast.x, forecast.y))
	// Not stack on top finished target
	if nextSpot == dead {
		return false
	}

	if nextSpot == bomb {
		var next = game.MoveObject(forecast, direction)
		if !next {
			return false
		}
	}

	if nextSpot == target {
		object = dead
		game.targets = Remove(game.targets, forecast)
	}

	// Made the movement
	game.state[forecast.y][forecast.x] = object
	game.state[location.y][location.x] = empty

	for i := 0; i < len(game.items); i++ {
		if game.items[i].Equal(location) {
			game.items[i] = Vector2{ forecast.x, forecast.y }
		}
	}

	return true
}

func (game *Sokoban) Show() string {
	return show(game.state, game.IsWin)
}

// Helper function
func fillGrid(width int, height int, player Vector2, targets []Vector2, items []Vector2) [][]string {
	// Create an empty 2D Grid
	var grid [][]string
	for i := 0; i < height; i++ {

		// For each column per row, insert an emoji depending on the coordinates given
		var row []string
		for j := 0; j < width; j++ {
			var coords = Vector2{j, i}
			if player.Equal(coords) {
				row = append(row, fox)
			} else if IsIn(items, coords) {
				row = append(row, bomb)
			} else if IsIn(targets, coords) {
				row = append(row, target)
			} else {
				row = append(row, empty)
			}
		}

		// Append the row to the grid
		grid = append(grid, row)
	}
	return grid
}

// View Functions
func show(grid [][]string, isWin bool) string {
	// Choose a random color unless game has ended
	var color = []string{"🟧", "🟪", "🟥"}[rand.Intn(3)]
	if isWin { color = "🟩" }

	// Draw the grid state
	var res = wall(len(grid[0]), color) + "\n"
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res += grid[i][j]
		}
		res += "\n"
	}
	res += wall(len(grid[0]), color) + "\n"

	// Add footer to explain available commands
	if isWin {
		res += "restart call command (!start)"
	} else {
		res += "To move, send W A S D"
	}
	return res
}

func wall(count int, word string) string {
	res := ""
	for i := 0; i < count; i++ {
		res += word
	}
	return res
}
