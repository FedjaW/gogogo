package data

import "fmt"

type distanceInMiles float32 
type distanceInKm float32

func (miles distanceInMiles) toKm() distanceInKm {
	return distanceInKm(miles * 1.60)
}

func Test() {
	const miles distanceInMiles = 99.3

	km := miles.toKm()
	fmt.Println(km)
}

