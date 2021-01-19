package games

import (
	"math/rand"
	"strings"
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
	player  []int
	targets [][]int
	items   [][]int
	isWin   bool
}

// Constructor
func NewSokoban(width int, height int) *Sokoban {
	var game = &Sokoban{
		state:   nil,
		player:  nil,
		targets: nil,
		items:   nil,
		isWin:   false,
	}

	var numObj = rand.Intn(2) + 3

	// Setup targets and item keys
	for i := 0; i < numObj; i++ {
		var target = []int{rand.Intn(height), rand.Intn(width)}
		var item = []int{rand.Intn(height-2) + 1, rand.Intn(width-2) + 1}
		game.targets = append(game.targets, target)

		for isIn(game.items, item) || isIn(game.targets, item) {
			item = []int{rand.Intn(height-2) + 1, rand.Intn(width-2) + 1}
		}

		game.items = append(game.items, item)
	}

	// Instantiate the player
	var player = []int{rand.Intn(height), rand.Intn(width)}
	for isIn(game.items, player) || isIn(game.targets, player) {
		player = []int{rand.Intn(height), rand.Intn(width)}
	}
	game.player = []int{player[0], player[1]}

	game.state = fillGrid(width, height, player, game.targets, game.items)

	return game
}

// Methods
func (game *Sokoban) Move(direction string) {
	if game.isWin {
		return
	}

	var forecast = []int{game.player[0], game.player[1]}

	// Made the changes aka the prediction
	switch strings.ToLower(direction) {
	case "w":
		forecast[0] -= 1
	case "s":
		forecast[0] += 1
	case "a":
		forecast[1] -= 1
	case "d":
		forecast[1] += 1
	}

	// Boundaries
	if forecast[0] < 0 || forecast[0] > len(game.state)-1 {
		return
	}
	if forecast[1] < 0 || forecast[1] > len(game.state[forecast[0]])-1 {
		return
	}

	// Check next item in the forecast
	var nextSpot = game.state[forecast[0]][forecast[1]]

	if nextSpot == dead || nextSpot == target {
		return
	}

	// Recursive movement
	if nextSpot == bomb {
		var hasMoved = game.MoveObject(forecast, direction)
		if !hasMoved {
			return
		}
	}

	// Switch the new location with the previous, basically move the character
	var temp = game.state[game.player[0]][game.player[1]]
	game.state[game.player[0]][game.player[1]] = empty
	game.state[forecast[0]][forecast[1]] = temp

	// Update the coordinates
	game.player = []int{forecast[0], forecast[1]}

	game.isWin = len(game.targets) == 0
}

func (game *Sokoban) MoveObject(location []int, direction string) bool {
	var forecast = []int{location[0], location[1]}

	// Made the changes aka prediction
	switch strings.ToLower(direction) {
	case "w":
		forecast[0] -= 1
	case "s":
		forecast[0] += 1
	case "a":
		forecast[1] -= 1
	case "d":
		forecast[1] += 1
	}

	// Boundaries
	if forecast[0] < 0 || forecast[0] > len(game.state)-1 {
		return false
	}
	if forecast[1] < 0 || forecast[1] > len(game.state[forecast[0]])-1 {
		return false
	}

	// Check next item in the forecast
	var nextSpot = game.state[forecast[0]][forecast[1]]
	var object = bomb

	// Not stack on top finished target
	if nextSpot == dead {
		return false
	}

	// Recursive movement
	if nextSpot == bomb {
		var hasMoved = game.MoveObject(forecast, direction)
		if !hasMoved {
			return false
		}
	}

	if nextSpot == target {
		object = dead
		game.targets = remove(game.targets, forecast)
	}

	// Made the movement
	game.state[forecast[0]][forecast[1]] = object
	game.state[location[0]][location[1]] = empty

	for i := 0; i < len(game.items); i++ {
		var curr = game.items[i]
		if curr[0] == location[0] && curr[1] == location[1] {
			game.items[i] = []int{forecast[0], forecast[1]}
		}
	}

	return true
}

func (game *Sokoban) Show() string {
	return show(game.state, game.isWin)
}

// Helper function
func fillGrid(width int, height int, player []int, targets [][]int, items [][]int) [][]string {
	// Create an empty 2D Grid
	var grid [][]string
	for i := 0; i < height; i++ {
		var row []string // Create an empty row

		// For each column per row, insert an emoji depending on the coordinates given
		for j := 0; j < width; j++ {
			var coords = []int{i, j}
			if player[0] == i && player[1] == j {
				row = append(row, fox)
			} else if isIn(targets, coords) {
				row = append(row, target)
			} else if isIn(items, coords) {
				row = append(row, bomb)
			} else {
				row = append(row, empty)
			}
		}

		// Append the row to the grid
		grid = append(grid, row)
	}
	return grid
}

func isIn(array [][]int, item []int) bool {
	for i := 0; i < len(array); i++ {
		if item[0] == array[i][0] && item[1] == array[i][1] {
			return true
		}
	}
	return false
}

func remove(array [][]int, item []int) [][]int {
	for i := 0; i < len(array); i++ {
		if array[i][0] == item[0] && array[i][1] == item[1] {
			return append(array[:i], array[i+1:]...)
		}
	}
	return array
}

func show(grid [][]string, isWin bool) string {
	// Choose a random color unless game has ended
	var color = []string{"🟧", "🟪", "🟥"}[rand.Intn(3)]
	if isWin {
		color = "🟩"
	}

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
