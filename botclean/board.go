package botclean

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	clean = '-'
	dirty = 'd'
	bot   = 'b'

	down  = "DOWN"
	up    = "UP"
	left  = "LEFT"
	right = "RIGHT"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	x, y, board := readBoard(r)
	direction := nextMove(x, y, board)
	go func() {
		w.Write([]byte(direction))
	}()
}

func readBoard(rd io.Reader) (int, int, [][]byte) {
	fmt.Println("Read board")
	var x, y int
	board := make([][]byte, 5)

	reader := bufio.NewReader(rd)
	fmt.Println("about to Read coords")

	line, _, err := reader.ReadLine()
	fmt.Printf("Read coords %#v\n", line)
	if err != nil {
		panic(err)
	}
	x, _ = strconv.Atoi(string(line[0]))
	y, _ = strconv.Atoi(string(line[2]))

	for i := 0; i < 5; i++ {
		buffer := make([]byte, 6)
		reader.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Read the line %#v\n", i)

		board[i] = buffer[0:5]
	}

	fmt.Println("Read the board")

	return x, y, board
}

func nextMove(x, y int, board [][]byte) string {
	// Use nearest neighbor for now
	if board[x][y] == dirty {
		board[x][y] = clean
		return "CLEAN"
	}

	i, j, direction := nearestNeighbor(x, y, board)

	if board[i][j] == clean {
		board[i][j] = bot
	}

	return direction
}

func nearestNeighbor(x, y int, board [][]byte) (int, int, string) {
	var neighborX, neighborY int
	var distance int
	var direction string

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == dirty {
				xDistance := x - i
				if xDistance < 0 {
					xDistance = -1 * xDistance
				}

				yDistance := y - j
				if yDistance < 0 {
					yDistance = -1 * yDistance
				}

				dirtyDistance := xDistance + yDistance
				if distance == 0 || dirtyDistance < distance {
					distance = dirtyDistance
					neighborX = i
					neighborY = j
				}
			}
		}
	}

	if neighborX < x {
		direction = left
	} else if neighborX > x {
		direction = right
	} else if neighborY < y {
		direction = up
	} else if neighborY > y {
		direction = down
	}

	return neighborX, neighborY, direction
}
