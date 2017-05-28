package botclean

import (
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
	assert.Equal(t, "RIGHT", direction)

	direction = nextMove(1, 0, input)
	assert.Equal(t, "DOWN", direction)

	direction = nextMove(1, 1, input)
	assert.Equal(t, "CLEAN", direction)
}

func TestSolve(t *testing.T) {
	input := "0 1\n-b--d\n-d--d\n--dd-\n--d--\n----d"

	r, w := io.Pipe()

	go func() {
		w.Write([]byte(input))
	}()
	solve(r, w)

	buffer := make([]byte, 5)
}
