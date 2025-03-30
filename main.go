package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

const (
		mars        = 62100000 // (km) The distance between earth and Mars at the time of the Departure Date on each ticket is October 13, 2020. 
		firstString = "Spaceline      |Days  |Trip type  | Price"
)

var (
	Spaceline, Triptype, Price string
	Days                       float64
)

func main() {
  
  SearchForTicketsToMars()

}

func SearchForTicketsToMars() {
    
    secondString := "="
    for i := 0; i < len(firstString)-1; i++ {
        secondString += "="
    }
    
    fmt.Printf("%s\n%s\n", firstString, secondString)

	for i := 0; i < 10; i++ {
		s := rand.Intn(3)
		switch s {
		case 0:
			Spaceline = "Virgin Galactic| "
		case 1:
			Spaceline = "SpaceX         | "
		case 2:
			Spaceline = "Space Adventure| "
		}

		w := rand.Intn(30) + 30 // The spacecraft's speed is randomly selected from a range of 30 to 60 km/s.
		Days = math.RoundToEven(float64(mars / (w * 3600) / 24))

		t := rand.Intn(2) 
		switch t {
		case 0:
			Triptype = " |One-way   "
		case 1:
			Triptype = " |Round-trip"
		}

		t++
		p := strconv.Itoa(t * (w + 20)) // The price for a round trip is doubled.
		if len(p) == 3 {
			Price = "| $ " + p
		} else {
			Price = "| $  " + p
		}

		fmt.Println(Spaceline, Days, Triptype, Price)

	}

}





