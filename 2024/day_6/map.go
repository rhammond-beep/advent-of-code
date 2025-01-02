package day6

type Point struct {
	X int
	Y int
}

type Map struct {
	ObstacleLocations map[Point]bool
	PositionsVisited  map[Point]bool
	GuardLocation     *Point
	Direction         string
	YUpperBound       int
	XUpperBound       int
}

/*
Create a lookup reference for us to refer back to
as the Guard traverses the lab
*/
func buildMap(input []string) Map {
	obstacleMap := make(map[Point]bool)
	positionsVisited := make(map[Point]bool)
	var guardLocation Point

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				obstacleMap[point] = true
			} else if input[i][j] == '^' {
				guardLocation = point
				positionsVisited[point] = true
			} else {
				positionsVisited[point] = false
				obstacleMap[point] = false
			}
		}
	}

	return Map{ObstacleLocations: obstacleMap, GuardLocation: &guardLocation, Direction: "north", YUpperBound: len(input), XUpperBound: len(input), PositionsVisited: positionsVisited}
}

/*
draw out a Ray based on where the guard is facing and return a point representing a barrier
*/
func (m *Map) WalkUntilBarrierFound() (Point, bool) {
	var closestPoint Point

	switch m.Direction {

	case "north":
		for i := m.GuardLocation.X - 1; i >= 0; i-- {
			point2Test := Point{X: i, Y: m.GuardLocation.Y}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "east":
		for i := m.GuardLocation.Y + 1; i < m.YUpperBound; i++ {
			point2Test := Point{X: m.GuardLocation.X, Y: i}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "south":
		for i := m.GuardLocation.X + 1; i < m.XUpperBound; i++ {
			point2Test := Point{X: i, Y: m.GuardLocation.Y}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "west":
		for i := m.GuardLocation.Y - 1; i >= 0; i-- {
			point2Test := Point{X: m.GuardLocation.X, Y: i}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	default:
		panic("we really shouldn't be here")
	}

	return closestPoint, true
}

func (m *Map) CountUniquePositionsVisited() (positionsVisited int) {
	for _, value := range m.PositionsVisited {
		if value {
			positionsVisited += 1
		}
	}
	return
}
