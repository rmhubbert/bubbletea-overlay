package overlay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clamp(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		min      int
		max      int
		expected int
	}{
		{"val 0, min 0, max 100", 0, 0, 100, 0},
		{"val 100, min 0, max 100", 100, 0, 100, 100},
		{"val -1, min 0, max 100", -1, 0, 100, 0},
		{"val 101, min 0, max 100", 101, 0, 100, 100},
		{"val -1, min 0, max -100", -1, 0, -100, -1},
		{"val 0, min 100, max 0", 0, 100, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := clamp(test.val, test.min, test.max)
			assert.Equal(t, test.expected, res, "they should be equal")
		})
	}
}

func Test_line(t *testing.T) {
	tests := []struct {
		name     string
		val      string
		expected int
	}{
		{"3 lines, no unexpected line endings", "aaa\nbbb\nccc", 3},
		{"3 lines, one unexpected line ending", "aaa\r\nbbb\nccc", 3},
		{"1 line, no line ending", "aaabbbccc", 1},
		{"empty string", "", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := lines(test.val)
			assert.Len(t, res, test.expected, "they should be equal")
		})
	}
}

// Note on this function. The foreground will be pushed 1 column left and / or up when the
// position is center, but the centering calculation does not have mod 0.
func Test_offsets(t *testing.T) {
	tests := []struct {
		name      string
		fg        string
		bg        string
		xPos      Position
		yPos      Position
		xOff      int
		yOff      int
		expectedX int
		expectedY int
	}{
		{
			"centered, odd fg height and width, no offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Center,
			Center,
			0,
			0,
			2,
			2,
		},
		{
			"centered, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Center,
			Center,
			0,
			0,
			2,
			3,
		},
		{
			"centered, odd fg height and width, with offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Center,
			Center,
			1,
			1,
			3,
			3,
		},
		{
			"centered, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Center,
			Center,
			0,
			0,
			2,
			3,
		},
		{
			"centered, even fg height and width, with offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Center,
			Center,
			1,
			1,
			3,
			4,
		},
		{
			"top left, odd fg height and width, no offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Left,
			Top,
			0,
			0,
			0,
			0,
		},
		{
			"top left, odd fg height and width, with offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Left,
			Top,
			1,
			1,
			1,
			1,
		},
		{
			"top left, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Left,
			Top,
			0,
			0,
			0,
			0,
		},
		{
			"top left, even fg height and width, with offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Left,
			Top,
			1,
			1,
			1,
			1,
		},
		{
			"top right, odd fg height and width, no offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Right,
			Top,
			0,
			0,
			4,
			0,
		},
		{
			"top right, odd fg height and width, with offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Right,
			Top,
			1,
			1,
			5,
			1,
		},
		{
			"top right, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Right,
			Top,
			0,
			0,
			5,
			0,
		},
		{
			"top right, even fg height and width, with offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Right,
			Top,
			1,
			1,
			6,
			1,
		},
		{
			"bottom left, odd fg height and width, no offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Left,
			Bottom,
			0,
			0,
			0,
			4,
		},
		{
			"bottom left, odd fg height and width, with offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Left,
			Bottom,
			1,
			1,
			1,
			5,
		},
		{
			"bottom left, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Left,
			Bottom,
			0,
			0,
			0,
			5,
		},
		{
			"bottom left, even fg height and width, with offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Left,
			Bottom,
			1,
			1,
			1,
			6,
		},
		{
			"bottom right, odd fg height and width, no offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Right,
			Bottom,
			0,
			0,
			4,
			4,
		},
		{
			"bottom right, odd fg height and width, with offset",
			strings.Repeat("abcde\n", 5),
			strings.Repeat("123456789\n", 9),
			Right,
			Bottom,
			1,
			1,
			5,
			5,
		},
		{
			"bottom right, even fg height and width, no offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Right,
			Bottom,
			0,
			0,
			5,
			5,
		},
		{
			"bottom right, even fg height and width, with offset",
			strings.Repeat("abcd\n", 4),
			strings.Repeat("123456789\n", 9),
			Right,
			Bottom,
			1,
			1,
			6,
			6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			x, y := offsets(test.fg, test.bg, test.xPos, test.yPos, test.xOff, test.yOff)
			assert.Equal(t, test.expectedX, x, "x should be equal")
			assert.Equal(t, test.expectedY, y, "y should be equal")
		})
	}
}

func Test_composite(t *testing.T) {
	fg := strings.Repeat("abc\n", 2) + "abc"
	bg := strings.Repeat("1234567\n", 6) + "1234567"

	tests := []struct {
		name     string
		xPos     Position
		yPos     Position
		xOff     int
		yOff     int
		fg       string
		bg       string
		expected string
	}{
		{
			"centered, no offset",
			Center,
			Center,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				2,
			) + strings.Repeat(
				"12abc67\n",
				3,
			) + "1234567\n1234567",
		},
		{
			"centered, with offset",
			Center,
			Center,
			1,
			1,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				3,
			) + strings.Repeat(
				"123abc7\n",
				3,
			) + "1234567",
		},
		{
			"top left, no offset",
			Left,
			Top,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"abc4567\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				3,
			) + "1234567",
		},
		{
			"top left, with offset",
			Left,
			Top,
			1,
			1,
			fg,
			bg,
			"1234567\n" + strings.Repeat(
				"1abc567\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				2,
			) + "1234567",
		},
		{
			"top center, no offset",
			Center,
			Top,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"12abc67\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				3,
			) + "1234567",
		},
		{
			"top center, with offset",
			Center,
			Top,
			1,
			1,
			fg,
			bg,
			"1234567\n" + strings.Repeat(
				"123abc7\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				2,
			) + "1234567",
		},
		{
			"top right, no offset",
			Right,
			Top,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234abc\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				3,
			) + "1234567",
		},
		{
			"top right, with offset",
			Right,
			Top,
			-1,
			1,
			fg,
			bg,
			"1234567\n" + strings.Repeat(
				"123abc7\n",
				3,
			) + strings.Repeat(
				"1234567\n",
				2,
			) + "1234567",
		},
		{
			"center left, no offset",
			Left,
			Center,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				2,
			) + strings.Repeat(
				"abc4567\n",
				3,
			) + "1234567\n1234567",
		},
		{
			"center left, with offset",
			Left,
			Center,
			1,
			1,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				3,
			) + strings.Repeat(
				"1abc567\n",
				3,
			) + "1234567",
		},
		{
			"center right, no offset",
			Right,
			Center,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				2,
			) + strings.Repeat(
				"1234abc\n",
				3,
			) + "1234567\n1234567",
		},
		{
			"center right, with offset",
			Right,
			Center,
			-1,
			1,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				3,
			) + strings.Repeat(
				"123abc7\n",
				3,
			) + "1234567",
		},
		{
			"bottom right, no offset",
			Right,
			Bottom,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				4,
			) + strings.Repeat(
				"1234abc\n",
				2,
			) + "1234abc",
		},
		{
			"bottom right, with offset",
			Right,
			Bottom,
			-1,
			-1,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				3,
			) + strings.Repeat(
				"123abc7\n",
				3,
			) + "1234567",
		},
		{
			"bottom left, no offset",
			Left,
			Bottom,
			0,
			0,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				4,
			) + strings.Repeat(
				"abc4567\n",
				2,
			) + "abc4567",
		},
		{
			"bottom left, with offset",
			Left,
			Bottom,
			1,
			-1,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				3,
			) + strings.Repeat(
				"1abc567\n",
				3,
			) + "1234567",
		},
		{
			"bottom left, out of bounds offset",
			Left,
			Bottom,
			10,
			10,
			fg,
			bg,
			strings.Repeat(
				"1234567\n",
				4,
			) + strings.Repeat(
				"1234abc\n",
				2,
			) + "1234abc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			x := composite(test.fg, test.bg, test.xPos, test.yPos, test.xOff, test.yOff)
			assert.Equal(t, test.expected, x, "x should be equal")
		})
	}
}
