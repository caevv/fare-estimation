package ride

import (
	"strconv"
	"time"

	"github.com/umahmood/haversine"
)

type Tuple struct {
	RideID string
	Lat    float64
	Lon    float64
	Time   time.Time
}

func (t *Tuple) DistanceInKm(lastTuple *Tuple) float64 {
	pointA := haversine.Coord{Lat: t.Lat, Lon: t.Lon}
	pointB := haversine.Coord{Lat: lastTuple.Lat, Lon: lastTuple.Lon}
	_, distanceInKm := haversine.Distance(pointA, pointB)

	return distanceInKm
}

func (t *Tuple) SpeedInKmh(lastTuple *Tuple) float64 {
	return t.DistanceInKm(lastTuple) / t.Time.Sub(lastTuple.Time).Hours()
}

func (t *Tuple) IsMoving(lastTuple *Tuple) bool {
	return t.SpeedInKmh(lastTuple) > 10
}

func (t *Tuple) IsInvalid(lastTuple *Tuple) bool {
	return t.SpeedInKmh(lastTuple) > 100
}

func NewTuple(rideID, lat, lon, timestamp string) (*Tuple, error) {
	lati, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, err
	}

	long, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil, err
	}

	timeInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Tuple{
		RideID: rideID,
		Lat:    lati,
		Lon:    long,
		Time:   time.Unix(timeInt, 0),
	}, nil
}
