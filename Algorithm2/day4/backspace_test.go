package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackSpace(t *testing.T) {
	assert.Equal(t, true, backspaceCompare("xywrrmp", "xywrrmu#p"))
}
