package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/caevv/fare-estimation/ride"
)

func Start(fileName string) {
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	tupleA, err := ride.NewTuple(record[0], record[1], record[2], record[3])
	if err != nil {
		log.Fatalln(err)
	}

	currentRide := ride.NewRide(*tupleA)

	var tupleB *ride.Tuple

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		tupleB, err = ride.NewTuple(record[0], record[1], record[2], record[3])
		if err != nil {
			log.Fatalln(err)
		}

		// New Ride?
		if tupleB.RideID != tupleA.RideID {
			tupleA = tupleB
			currentRide = ride.NewRide(*tupleA)

			currentRide.AddTuple(*tupleA)

			continue
		}

		if tupleB.IsInvalid(tupleA) {
			// Do not add invalid tuple
			continue
		}

		currentRide.AddTuple(*tupleB)

		fmt.Printf("Ride ID: %s Ride Fare: %v \n", currentRide.RideID, currentRide.Fare)
	}
}
