package word

import (
	"strings"
	"unicode"
)

// Word represents a word
type Word struct {
	runes []rune
}

// New instantiate word
func New(str string) Word {
	return Word{runes: []rune(str)}
}

// Cleanup removes not alphanumeric
func (r Word) Cleanup() Word {
	words := strings.FieldsFunc(string(r.runes), isLetterOrNumber)
	return Word{runes: []rune(strings.Join(words, ""))}
}

func (r Word) String() string {
	return string(r.runes)
}

// Runes return runes
func (r Word) Runes() []rune {
	return r.runes
}

func isLetterOrNumber(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) }
