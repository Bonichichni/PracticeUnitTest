package DoFunc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	testingData := map[string]struct {
		char   string
		num    int
		b      bool
		exp    string
		expErr string
	}{
		"success":               {"a", 8, true, "[8A]", ""},
		"invalid int":           {"f", 9, false, "", "invalid s"},
		"invalid char":          {"g", 13, true, "", "invalid s"},
		"success but not upper": {"a", 13, false, "a", ""},
	}
	for name, tt := range testingData {
		t.Run(name, func(t *testing.T) {
			actual, err := Do(tt.char, tt.num, tt.b)
			assert.Equal(t, tt.exp, actual)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
