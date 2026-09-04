package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Room struct {
	Name string
	X int
	Y int
}

type Tunnel struct {
	Room1 string
	Room2 string
}


type AntFarm struct {
	NumberOfAnts int
	Rooms []Room
	StartRoom Room
	EndRoom Room
	Tunnels []Tunnel
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
	
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	var farm AntFarm

	scanner.Scan()
	line := scanner.Text()

	num, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Error converting number of ants:", err)
		return
	}

	farm.NumberOfAnts = num
	fmt.Println("Number of ants:", farm.NumberOfAnts)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}