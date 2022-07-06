package test_practice_test

import (
	tp "go-practice/test_practice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEstimateValue(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		assert.Equal(t, "small", tp.EstimateValue(9))
	})

	t.Run("Medium", func(t *testing.T) {
		assert.Equal(t, "medium", tp.EstimateValue(99))
	})

	t.Run("Big", func(t *testing.T) {
		assert.Equal(t, "big", tp.EstimateValue(100))
	})
}
