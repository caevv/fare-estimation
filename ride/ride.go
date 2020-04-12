package ride

type Ride struct {
	RideID   string
	Segments []Segment
	Fare     float64
}

func NewRide(t Tuple) *Ride {
	return &Ride{
		RideID: t.RideID,
		Fare:   1.30, // initial fare
		Segments: []Segment{{
			Tuples:      []Tuple{t},
			InitialTime: t.Time,
		}},
	}
}

func (r *Ride) AddTuple(t Tuple) {

	// TODO: split every 1km instead of just more than 1km on every tuple
	if r.Segments[len(r.Segments)-1].DistanceInKm() < 1 {
		r.Segments[len(r.Segments)-1].Tuples = append(r.Segments[len(r.Segments)-1].Tuples, t)

		return
	}

	// Calculate Fare
	if r.Segments[len(r.Segments)-1].IsHighRate() {
		r.Fare += 1.30
	} else {
		r.Fare += 0.74
	}

	// TODO: idle hour

	r.Segments = append(r.Segments, Segment{
		Tuples:      []Tuple{t},
		InitialTime: t.Time,
	})
}
