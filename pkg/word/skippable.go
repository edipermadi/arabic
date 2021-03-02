package word

import "github.com/edipermadi/unicode"

var skipableCharacters = map[rune]struct{}{
	// Subtending marks
	unicode.ArabicNumberSign:     {},
	unicode.ArabicSignSanah:      {},
	unicode.ArabicFootnoteMarker: {},
	unicode.ArabicSignSafha:      {},
	unicode.ArabicSignSamvat:     {},

	// Supertending mark
	unicode.ArabicNumberMarkAbove: {},

	// Radix symbols
	unicode.ArabicIndicCubeRoot:   {},
	unicode.ArabicIndicFourthRoot: {},
}
