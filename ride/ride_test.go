package ride_test

import (
	"testing"
	"time"

	"github.com/caevv/fare-estimation/ride"
	"github.com/stretchr/testify/assert"
)

func TestNewRide(t *testing.T) {
	assert.Equal(
		t,
		&ride.Ride{
			RideID: "1",
			Fare:   1.30,
			Segments: []ride.Segment{{
				Tuples: []ride.Tuple{{
					RideID: "1",
					Lat:    10,
					Lon:    20,
					Time:   time.Unix(100, 0),
				}},
				InitialTime: time.Unix(100, 0),
			},
			},
		},
		ride.NewRide(ride.Tuple{
			RideID: "1",
			Lat:    10,
			Lon:    20,
			Time:   time.Unix(100, 0),
		}),
	)

}

func TestAddTuple(t *testing.T) {
	r := &ride.Ride{
		RideID: "1",
		Fare:   1.30,
		Segments: []ride.Segment{{
			Tuples: []ride.Tuple{{
				RideID: "1",
				Lat:    10,
				Lon:    20,
				Time:   time.Unix(100, 0),
			}},
			InitialTime: time.Unix(100, 0),
		},
		},
	}

	myTuple := ride.Tuple{
		RideID: "1",
		Lat:    11,
		Lon:    12,
		Time:   time.Unix(200, 0),
	}

	r.AddTuple(myTuple)

	expectedRide := &ride.Ride{
		RideID: "1",
		Fare:   1.30,
		Segments: []ride.Segment{{
			Tuples: []ride.Tuple{{
				RideID: "1",
				Lat:    10,
				Lon:    20,
				Time:   time.Unix(100, 0),
			}, {
				RideID: "1",
				Lat:    11,
				Lon:    12,
				Time:   time.Unix(200, 0),
			}},
			InitialTime: time.Unix(100, 0),
		}},
	}

	// less than 1 KM
	assert.Equal(t, expectedRide, r)

	// More than 1KM
	myTuple = ride.Tuple{
		RideID: "1",
		Lat:    10000,
		Lon:    10000,
		Time:   time.Unix(300, 0),
	}

	r.AddTuple(myTuple)
	expectedRide.Fare = 2.60

	assert.Equal(t, expectedRide.Fare, r.Fare)
}
