package mazeSolver

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func drawPath(data []string, path []Point) []string {
	data2 := make([][]string, len(data))
	for i := 0; i < len(data); i++ {
		data2[i] = strings.Split(data[i], "")
	}

	for i := 0; i < len(path); i++ {
		if len(data2[path[i].y]) > 0 && len(data2[path[i].y][path[i].x]) > 0 {
			data2[path[i].y][path[i].x] = "*"
		}
	}

	out := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		out[i] = strings.Join(data2[i], "")
	}
	return out
}

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expected := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	result := MazeSolver(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})
	drawResult := drawPath(maze, result)
	expectedResult := drawPath(maze, expected)

	fmt.Println(result)
	fmt.Println(expected)

	if !reflect.DeepEqual(drawResult, expectedResult) {
		t.Fail()
	}
}
