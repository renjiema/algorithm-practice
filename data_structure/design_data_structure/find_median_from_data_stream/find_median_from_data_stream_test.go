package find_median_from_data_stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	t.Log((-1) / 2)
}

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian())
	mf.AddNum(3)
	assert.Equal(t, float64(2), mf.FindMedian())
	mf.AddNum(5)
	assert.Equal(t, float64(2.5), mf.FindMedian())
	mf.AddNum(4)
	assert.Equal(t, float64(3), mf.FindMedian())

	mf = Constructor()
	mf.AddNum(-1)
	assert.Equal(t, float64(-1), mf.FindMedian())
	mf.AddNum(-2)
	assert.Equal(t, -1.5, mf.FindMedian())
	mf.AddNum(-3)
	assert.Equal(t, -2.0, mf.FindMedian())
	mf.AddNum(-4)
	assert.Equal(t, -2.5, mf.FindMedian())
	mf.AddNum(-5)
	assert.Equal(t, -3.0, mf.FindMedian())

	mf = Constructor()
	mf.AddNum(6)
	assert.Equal(t, 6.0, mf.FindMedian())
	mf.AddNum(10)
	assert.Equal(t, 8.0, mf.FindMedian())
	mf.AddNum(2)
	assert.Equal(t, 6.0, mf.FindMedian())
	mf.AddNum(6)
	assert.Equal(t, 6.0, mf.FindMedian())
	mf.AddNum(5)
	assert.Equal(t, 6.0, mf.FindMedian())
	mf.AddNum(0)
	assert.Equal(t, 5.5, mf.FindMedian())
	mf.AddNum(6)
	assert.Equal(t, 6.0, mf.FindMedian())
}
