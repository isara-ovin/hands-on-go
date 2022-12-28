package sample

import "fmt"

type bmw struct {
	distance float64
	fuel     float64
}

type benz struct {
	distance float64
	fuel     float64
	speed    int
}

// Recivers for each mileage
func (b bmw) mileage() float64 {
	return b.distance / b.fuel
}

func (z benz) mileage() float64 {
	return z.distance / z.fuel
}

type vehicle interface {
	mileage() float64
}

func main() {

	b := bmw{
		distance: 1000.0,
		fuel:     33.3,
	}

	z := benz{
		distance: 1000.0,
		fuel:     33.3,
		speed:    86,
	}

	mileages := []vehicle{b, z}
	total := 0.0
	for _, m := range mileages {
		total += m.mileage()
	}

	fmt.Println("Total Mileage", total)
}
