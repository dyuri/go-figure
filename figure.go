package figure

import (
	"fmt"
	"io"
	"log"
	"strings"
)

const asciiOffset = 32
const firstAscii = ' '
const lastAscii = '~'

type colorizer func(cidx, row int, fragment, text string) string

type figure struct {
	font
	phrase string
	strict bool
}

func NewFigure(phrase, fontName string, strict bool) figure {
	font := newFont(fontName)
	if font.reverse {
		phrase = reverse(phrase)
	}
	return figure{font: font, phrase: phrase, strict: strict}
}

func NewFigureWithFont(phrase string, reader io.Reader, strict bool) figure {
	font := newFontFromReader(reader)
	if font.reverse {
		phrase = reverse(phrase)
	}
	return figure{font: font, phrase: phrase, strict: strict}
}

func (figure figure) Slicify(cfun colorizer) (rows []string) {
	for r := 0; r < figure.height; r++ {
		printRow := ""
		for idx, char := range figure.phrase {
			if char < firstAscii || char > lastAscii {
				if figure.strict {
					// TODO err
					log.Fatal("invalid input.")
				} else {
					char = '?'
				}
			}
			fontIndex := char - asciiOffset
			charRowText := scrub(figure.letters[fontIndex][r], figure.hardblank)
			if cfun != nil && strings.ReplaceAll(charRowText, " ", "") != "" {
				charRowText = cfun(idx, r, charRowText, figure.phrase)
			}
			printRow += charRowText
		}
		if r < figure.baseline || len(strings.TrimSpace(printRow)) > 0 {
			rows = append(rows, strings.TrimRight(printRow, " "))
		}
	}
	return rows
}

func (fig figure) Print() {
	for _, printRow := range fig.Slicify(nil) {
		fmt.Println(printRow)
	}
}

func (fig figure) ColorString(cfun colorizer) string {
	s := ""
	for _, printRow := range fig.Slicify(cfun) {
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

func (fig figure) String() string {
	s := ""
	for _, printRow := range fig.Slicify(nil) {
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

func scrub(text string, char byte) string {
	return strings.ReplaceAll(text, string(char), " ")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
