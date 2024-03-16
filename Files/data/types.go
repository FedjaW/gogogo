package data

import "fmt"

type location string

func (origin location) DistanceTo(destination location) distanceInKm {
	// some calc goes here...
 	return 10
}

func locationTest() {
	nyc := location("33.33, 44.44")
	nyc.DistanceTo(location("12.2, 43."))
}

type distanceInMiles float32 
type distanceInKm float32

// not a function but a method
func (miles distanceInMiles) toKm() distanceInKm {
	return distanceInKm(miles * 1.60)
}

// function
func Test() {
	const miles distanceInMiles = 99.3

	km := miles.toKm()
	fmt.Println(km)
}

