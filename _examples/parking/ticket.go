package main

import "time"

type ParkingTicket struct {
	EntryTime  time.Time
	SlotNumber int
	Vehicle    Vehicle
}
