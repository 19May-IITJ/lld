package main

type Vehicle struct {
	LicensePlate string
	Color        string
}

func NewVehicle(plate, color string) Vehicle {
	return Vehicle{
		LicensePlate: plate,
		Color:        color,
	}
}
