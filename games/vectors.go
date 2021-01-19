package games

type Vector2 struct {
	x, y int
}


func (vec Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		x: vec.x + other.x,
		y: vec.y + other.y,
	}
}

func (vec Vector2) Equal(other Vector2) bool {
	return vec.x == other.x && vec.y == other.y
}
