package main

type Room struct {
	Name string
	X    int
	Y    int
}

type Tunnel struct {
	Room1 string
	Room2 string
}

type AntFarm struct {
	NumberOfAnts int
	Rooms        []Room
	StartRoom    Room
	EndRoom      Room
	Tunnels      []Tunnel
}