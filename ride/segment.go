package ride

import (
	"time"

	"github.com/umahmood/haversine"
)

type Segment struct {
	Tuples      []Tuple
	InitialTime time.Time
}

func (s Segment) IsHighRate() bool {
	return s.InitialTime.Hour() < 5
}

func (s Segment) DistanceInKm() float64 {
	pointA := haversine.Coord{Lat: s.Tuples[0].Lat, Lon: s.Tuples[0].Lon}
	pointB := haversine.Coord{Lat: s.Tuples[len(s.Tuples)-1].Lat, Lon: s.Tuples[len(s.Tuples)-1].Lon}
	_, distanceInKm := haversine.Distance(pointA, pointB)

	return distanceInKm
}
