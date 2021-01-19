package games

type Vector2 struct {
	x, y int
}

type Vector2Direction string

const (
	up Vector2Direction = "up"
	down Vector2Direction = "down"
	left Vector2Direction = "left"
	right Vector2Direction = "right"
)

var (
	DirectionKeyCode = map[string] Vector2Direction {
		"w": up,
		"a": left,
		"s": down,
		"d": right,
	}
)

func (vec Vector2) Add(other Vector2) Vector2 {
	return Vector2{ vec.x + other.x,vec.y + other.y}
}

func (vec Vector2) Equal(other Vector2) bool {
	return vec.x == other.x && vec.y == other.y
}

func (vec Vector2) ApplyMovement(direction Vector2Direction) Vector2 {
	var (
		x, y = vec.x, vec.y
	)
	switch direction {
	case up:
		y -= 1
	case down:
		y += 1
	case left:
		x -= 1
	case right:
		x += 1
	}
	return Vector2{x,y}
}

func IsIn(arr []Vector2, item Vector2) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].Equal(item) {
			return true
		}
	}
	return false
}

func Remove(arr []Vector2, item Vector2) []Vector2 {
	for i := 0; i < len(arr); i++ {
		if arr[i].x == item.x && arr[i].y == item.y {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
