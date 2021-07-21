package bizerror

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMergeErr(t *testing.T) {
	err := mergeErr(Silence(errors.New("")), errors.New(""))
	assert.True(t, IsSilence(err))
}
