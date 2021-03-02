package word_test

import (
	"testing"

	"github.com/edipermadi/arabic/pkg/word"
	"github.com/edipermadi/unicode"
	"github.com/stretchr/testify/require"
)

func TestWord_Parse(t *testing.T) {
	src := word.New("كَتَبَ")
	dst := src.Cleanup()
	require.Equal(t, []rune{unicode.ArabicLetterKaf, unicode.ArabicLetterTeh, unicode.ArabicLetterBeh}, dst.Runes())
	t.Logf("cleaned = %s", dst.String())
}
