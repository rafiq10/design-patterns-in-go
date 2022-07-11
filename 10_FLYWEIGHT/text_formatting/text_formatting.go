package textformatting

import (
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plTxt string) *FormattedText {
	return &FormattedText{
		plainText: plTxt, capitalize: make([]bool, len(plTxt)),
	}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

func (t *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end, Capitalize: false, Bold: false, Italic: false}
	t.formatting = append(t.formatting, r)
	return r
}

func (t *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(t.plainText); i++ {
		c := t.plainText[i]
		for _, r := range t.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}
