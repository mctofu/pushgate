package pushgate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	tests := []struct {
		s       string
		limit   int
		trimmed string
	}{
		{"below limit", 12, "below limit"},
		{"at limit", 8, "at limit"},
		{"over limit", 9, "over l..."},
		{"below with 謝", 13, "below with 謝"},
		{"at with 謝", 9, "at with 謝"},
		{"over with 謝", 10, "over wi..."},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.trimmed, trim(test.s, test.limit))
		})
	}
}
