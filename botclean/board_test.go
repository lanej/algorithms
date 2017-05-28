package botclean

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBoard(t *testing.T) {
	input := "0 1\n-b--d\n-d--d\n--dd-\n--d--\n----d"

	r := strings.NewReader(input)

	x, y, board := readBoard(r)
	expectedBoard := [][]byte{
		[]byte{clean, bot, clean, clean, dirty},
		[]byte{clean, dirty, clean, clean, dirty},
		[]byte{clean, clean, dirty, dirty, clean},
		[]byte{clean, clean, dirty, clean, clean},
		[]byte{clean, clean, clean, clean, dirty},
	}

	assert.Equal(t, 0, x)
	assert.Equal(t, 1, y)
	assert.Equal(t, expectedBoard, board)
}

func TestMove1(t *testing.T) {
	input := [][]byte{
		[]byte{bot, clean, clean, clean, dirty},
		[]byte{clean, dirty, clean, clean, dirty},
		[]byte{clean, clean, dirty, dirty, clean},
		[]byte{clean, clean, dirty, clean, clean},
		[]byte{clean, clean, clean, clean, dirty},
	}

	direction := nextMove(0, 0, input)
	assert.Equal(t, "DOWN", direction)

	direction = nextMove(1, 0, input)
	assert.Equal(t, "RIGHT", direction)

	direction = nextMove(1, 1, input)
	assert.Equal(t, "CLEAN", direction)
}

func TestMove2(t *testing.T) {
	input := [][]byte{
		[]byte{bot, dirty, clean, clean, clean},
		[]byte{clean, dirty, clean, clean, clean},
		[]byte{clean, clean, dirty, clean, clean},
		[]byte{clean, clean, clean, dirty, clean},
		[]byte{clean, clean, dirty, clean, dirty},
	}

	direction := nextMove(0, 0, input)
	assert.Equal(t, "RIGHT", direction)

	direction = nextMove(0, 1, input)
	assert.Equal(t, "CLEAN", direction)

	direction = nextMove(0, 1, input)
	assert.Equal(t, "DOWN", direction)

	direction = nextMove(1, 1, input)
	assert.Equal(t, "CLEAN", direction)
}

func TestMove3(t *testing.T) {
	input := [][]byte{
		[]byte{clean, dirty, clean, clean, clean},
		[]byte{clean, dirty, clean, clean, clean},
		[]byte{clean, clean, dirty, clean, clean},
		[]byte{clean, clean, clean, dirty, clean},
		[]byte{clean, clean, dirty, clean, dirty},
	}

	direction := nextMove(0, 1, input)
	assert.Equal(t, "CLEAN", direction)
}

func TestSolve(t *testing.T) {
	input := "0 1\n-b--d\n-d--d\n--dd-\n--d--\n----d"

	r, in := io.Pipe()
	out, w := io.Pipe()

	go func() {
		in.Write([]byte(input))
	}()

	go func() {
		buffer := make([]byte, 6)
		reader := bufio.NewReader(out)
		n, err := reader.Read(buffer)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "RIGHT", string(buffer[0:n]))
	}()

	solve(r, w)
}

func print(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf(string(board[i][j]))
		}
		fmt.Println()
	}
}
