// Package cistercian contains the logic for generating Cistercian numbers.
package cistercian

import (
	"fmt"
	"strings"
)

const (
	CistercianNumberWidth  int = 9    // CistercianNumberWidth is the width of a Cistercian number.
	CistercianNumberHeight int = 13   // CistercianNumberHeight is the height of a Cistercian number.
	MaxCistercianNumber    int = 9999 // MaxCistercianNumber is the maximum Cistercian number.
)

// boxChars contains the characters used to draw the Cistercian number.
var boxChars = [...]rune{'─', '│', '┌', '┐', '└', '┘', '├', '┤', '┬', '┴', '┼', '╱', '╲', '╴', '╵', '╶', '╷'}

// GenerateCistercianNumber generates a Cistercian number.
func GenerateCistercianNumber(number int, isSVG bool) (string, error) {
	if isSVG {
		return generateCistercianNumberSVG(number)
	}

	return generateCistercianNumberASCII(number)
}

func generateCistercianNumberASCII(number int) (string, error) {
	if number > MaxCistercianNumber {
		return "", fmt.Errorf("number too large: %d", number)
	}

	// for each digit, get glyph and put in glyph array, then combine all glyphs into final glyph
	glyph := make([]rune, CistercianNumberHeight*CistercianNumberWidth)

	// Add center stave, which is always there
	addCenterStave(glyph)

	addOnesPlace(number%10, glyph) // Ones place

	sb := strings.Builder{}

	for row := 0; row < CistercianNumberHeight; row++ {
		if row > 0 {
			sb.WriteString("\n")
		}

		for col := 0; col < CistercianNumberWidth; col++ {
			sb.WriteRune(glyph[row*CistercianNumberWidth+col])
		}
	}

	return sb.String(), nil
}

func addOnesPlace(digit int, glyph []rune) {
	getDigit(digit, glyph)
}

func getDigit(digit int, glyph []rune) {
	switch digit {
	case 1:
		offset := 4

		glyph[offset] = boxChars[2]

		for i := 1; i <= 4; i++ {
			glyph[offset+i] = boxChars[0]
		}
	case 2:
		offset := 4 + CistercianNumberWidth*4

		glyph[offset] = boxChars[6]
		for i := 1; i <= 4; i++ {
			glyph[offset+i] = boxChars[0]
		}
	case 3:
		offset := 5

		glyph[offset+CistercianNumberWidth] = boxChars[12]
		glyph[offset+CistercianNumberWidth*2+1] = boxChars[12]
		glyph[offset+CistercianNumberWidth*3+2] = boxChars[12]
	case 4:
		offset := 8

		glyph[offset+CistercianNumberWidth-1] = boxChars[11]
		glyph[offset+CistercianNumberWidth*2-2] = boxChars[11]
		glyph[offset+CistercianNumberWidth*3-3] = boxChars[11]
	case 5:
		offset := 4

		glyph[offset] = boxChars[2]
		for i := 1; i <= 4; i++ {
			glyph[offset+i] = boxChars[0]
		}

		offset = 8
		glyph[offset+CistercianNumberWidth-1] = boxChars[11]
		glyph[offset+CistercianNumberWidth*2-2] = boxChars[11]
		glyph[offset+CistercianNumberWidth*3-3] = boxChars[11]
	case 6:
		offset := 8

		glyph[offset] = boxChars[16]
		for i := 1; i < 4; i++ {
			glyph[offset+CistercianNumberWidth*i] = boxChars[1]
		}

		glyph[offset+CistercianNumberWidth*4] = boxChars[14]
	case 7:
		offset := 4

		glyph[offset] = boxChars[2]
		for i := 1; i < 4; i++ {
			glyph[offset+i] = boxChars[0]
		}

		glyph[offset+4] = boxChars[3]

		offset = 8

		glyph[offset+CistercianNumberWidth] = boxChars[1]
		glyph[offset+CistercianNumberWidth*2] = boxChars[1]
		glyph[offset+CistercianNumberWidth*3] = boxChars[1]
		glyph[offset+CistercianNumberWidth*4] = boxChars[14]
	case 8:
		offset := 8

		glyph[offset] = boxChars[16]
		glyph[offset+CistercianNumberWidth] = boxChars[1]
		glyph[offset+CistercianNumberWidth*2] = boxChars[1]
		glyph[offset+CistercianNumberWidth*3] = boxChars[1]

		offset = 4 + CistercianNumberWidth*4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	case 9:
		offset := 4

		glyph[offset] = boxChars[2]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[3]

		offset = 8

		glyph[offset+CistercianNumberWidth] = boxChars[1]
		glyph[offset+CistercianNumberWidth*2] = boxChars[1]
		glyph[offset+CistercianNumberWidth*3] = boxChars[1]

		offset = 4 + CistercianNumberWidth*4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	default:

	}
}

func addCenterStave(glyph []rune) {
	for row := 0; row < CistercianNumberHeight; row++ {
		char := boxChars[1]

		if row == 0 {
			char = boxChars[16]
		}

		if row == CistercianNumberHeight-1 {
			char = boxChars[14]
		}

		for col := 0; col < CistercianNumberWidth; col++ {
			if col == 4 {
				glyph[row*CistercianNumberWidth+col] = char
			} else {
				glyph[row*CistercianNumberWidth+col] = ' '
			}
		}
	}
}

func generateCistercianNumberSVG(number int) (string, error) {
	panic("unimplemented")
}
