package errorx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCf(t *testing.T) {

	tests := []struct {
		name               string
		code               int
		format             string
		args               interface{}
		expectedErrErrType ErrType
	}{
		{
			name:               "ERRTYPE_INTERNAL_SERVER",
			code:               CODE_INTERNALSERVER,
			expectedErrErrType: ERRTYPE_INTERNAL_SERVER,
			format:             "internal server error: %s",
			args:               "database failure",
		},
		{
			name:               "CODE_NOT_FOUND",
			code:               CODE_NOT_FOUND,
			expectedErrErrType: ERRTYPE_NOT_FOUND,
			format:             "not_found error: %s",
			args:               "path not found",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d: %s", i, test.name), func(t *testing.T) {
			formattedErr := Cf(test.code, test.format, test.args)
			assert.NotNil(t, formattedErr)
			assert.Contains(t, fmt.Sprintf("Formatted Error: %+v\n", formattedErr), test.expectedErrErrType)
			assert.Equal(t, test.expectedErrErrType, formattedErr.Type())
			assert.Equal(t, test.code, formattedErr.Code())
		})

	}
}
