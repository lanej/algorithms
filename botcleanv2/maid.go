package botcleanv2

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const (
	clean  = '-'
	dirty  = 'd'
	bot    = 'b'
	hidden = 'o'

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
	w.Write([]byte(direction))
}

func readBoard(rd io.Reader) (int, int, [][]byte) {
	var x, y int
	h := 5
	w := 5

	reader := bufio.NewReader(rd)

	if line, _, err := reader.ReadLine(); err != nil {
		panic(err)
	} else {
		x, _ = strconv.Atoi(string(line[0]))
		y, _ = strconv.Atoi(string(line[2]))
	}

	board := make([][]byte, h)

	for i := 0; i < h; i++ {
		buffer := make([]byte, w+1)
		reader.Read(buffer)

		board[i] = buffer[0:w]
	}

	return x, y, board
}

func nextMove(x, y int, board [][]byte) string {
	// Use nearest neighbor for now
	if board[x][y] == dirty {
		board[x][y] = clean
		return "CLEAN"
	}

	i, j, direction := nearestNeighbor(x, y, board)

	board[x][y] = clean

	if board[i][j] == clean {
		board[i][j] = bot
	}

	return direction
}

func nearestNeighbor(x, y int, board [][]byte) (int, int, string) {
	var nx, ny, distance int
	var direction string

	mark := func(i, j int) {
		xd := x - i
		if xd < 0 {
			xd = -1 * xd
		}

		yd := y - j
		if yd < 0 {
			yd = -1 * yd
		}

		nd := xd + yd

		if distance == 0 || nd < distance {
			distance = nd
			nx = i
			ny = j
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == dirty {
				mark(i, j)
			}
		}
	}

	if distance == 0 {
		nx, ny = search(x, y, board)
	}

	if nx < x {
		direction = up
	} else if nx > x {
		direction = down
	} else if ny < y {
		direction = left
	} else if ny > y {
		direction = right
	}

	return nx, ny, direction
}

func search(x, y int, board [][]byte) (int, int) {
	searchPattern := [][]int{
		[]int{1, 1},
		[]int{2, 1},
		[]int{3, 1},
		[]int{3, 2},
		[]int{3, 3},
		[]int{2, 3},
		[]int{1, 3},
		[]int{1, 2},
		[]int{1, 1},
	}

	// move to search pattern
	if x == 0 {
		return 1, y
	} else if y == 0 {
		return x, 1
	} else if y == 4 {
		return x, 3
	} else if x == 4 {
		return 3, y
	}

	// progress search pattern
	var index int
	for i, point := range searchPattern {
		if point[0] == x && point[1] == y {
			index = i
			break
		}
	}
	next := searchPattern[index+1]
	return next[0], next[1]
}
