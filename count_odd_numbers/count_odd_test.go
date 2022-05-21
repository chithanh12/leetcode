package count_odd_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOdd(t *testing.T) {
	count := countOdds(0, 9)
	assert.Equal(t, 5, count)

	assert.Equal(t, 3, countOdds(3, 8))
	assert.Equal(t, 1, countOdds(8, 10))
}
