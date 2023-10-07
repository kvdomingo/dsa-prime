package mazeSolver

import "strings"

type Point struct {
	x, y int
}

func walk(maze []string, wall string, current Point, end Point, seen [][]bool, path []Point) bool {
	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	if current.x < 0 || current.x >= len(maze[0]) || current.y < 0 || current.y >= len(maze) {
		return false
	}

	if strings.Split(maze[current.y], "")[current.x] == wall {
		return false
	}

	if current.x == end.x && current.y == end.y {
		path = append(path, end)
		return true
	}

	if seen[current.y][current.x] {
		return false
	}

	seen[current.y][current.x] = true
	path = append(path, current)

	for i := 0; i < len(dir); i++ {
		d := dir[i]
		x, y := d[0], d[1]
		if walk(maze, wall, Point{x: current.x + x, y: current.y + y}, end, seen, path) {
			return true
		}
	}

	path = path[:len(path)-1]

	return false
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	var seen [][]bool
	var path []Point

	for i := 0; i < len(maze); i++ {
		seen = append(seen, make([]bool, len(maze[i])))
	}

	walk(maze, wall, start, end, seen, path)
	return path
}
