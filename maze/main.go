package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func readMazeFromFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	var maze [][]int = make([][]int, row)
	for i := 0; i < row; i++ {
		maze[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

var directions = []point{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func (cur point) add(p point) point {
	return point{p.i + cur.i, p.j + cur.j}
}
func (p point) at(grid *[][]int) (int, bool) {
	if p.i < 0 || p.j < 0 {
		return 0, false
	}
	if p.i >= len(*grid) {
		return 0, false
	}
	if p.j >= len((*grid)[p.i]) {
		return 0, false
	}
	return (*grid)[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}
	for len(queue) > 0 {
		curPoint := queue[0]
		if curPoint == end {
			break
		}
		queue = queue[1:]
		for _, direPoint := range directions {
			nextPoint := curPoint.add(direPoint)
			if nextPoint == start {
				continue
			}
			if value, ok := nextPoint.at(&maze); value == 1 || !ok {
				continue
			}
			if value, ok := nextPoint.at(&steps); value != 0 || !ok {
				continue
			}
			step, _ := curPoint.at(&steps)
			steps[nextPoint.i][nextPoint.j] = step + 1
			queue = append(queue, nextPoint)
		}
	}
	return steps
}
func main() {
	maze := readMazeFromFile("./maze/maze.in")
	steps := walk(maze, point{
		i: 0,
		j: 0,
	}, point{
		i: 5,
		j: 4,
	})
	for _, row := range steps {
		for _, value := range row {
			fmt.Printf("%3d", value)
		}
		fmt.Println()
	}
}
