package ride_test

import (
	"testing"
	"time"

	"github.com/caevv/fare-estimation/ride"
	"github.com/stretchr/testify/assert"
)

func TestSegmentRate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected bool
	}{
		{"2019-01-01T14:00:00Z", false},
		{"2019-01-01T00:00:00Z", true},
		{"2019-01-01T01:00:00Z", true},
		{"2019-01-01T06:00:00Z", false},
	}

	var seg ride.Segment

	for _, test := range tests {
		initialTime, err := time.Parse(time.RFC3339, test.input)
		assert.NoError(err)

		seg = ride.Segment{
			InitialTime: initialTime,
		}
		assert.Equal(seg.IsHighRate(), test.expected)
	}
}
