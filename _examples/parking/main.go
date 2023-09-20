package main

import (
	"fmt"
)

func main() {
	parkingLot := CreateParkingLot(10)

	vehicle1 := NewVehicle(
		"ABC123",
		"Red")
		
	ticket1, err := parkingLot.ParkVehicle(vehicle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Parked Vehicle 1:", ticket1)

		err := parkingLot.ExitParkingLot(ticket1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Vehicle 1 has exited")
		}
	}

	vehicle2 := Vehicle{LicensePlate: "XYZ789", Color: "Blue"}
	ticket2, err := parkingLot.ParkVehicle(vehicle2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Parked Vehicle 2:", ticket2)
	}
}
