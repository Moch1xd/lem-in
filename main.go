package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file>")
		return
	}

	farm, err := ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	fmt.Println("Number of ants:", farm.NumberOfAnts)
	fmt.Println("Start room:", farm.StartRoom)
	fmt.Println("End room:", farm.EndRoom)
	fmt.Println("Total rooms:", len(farm.Rooms))
	fmt.Println("Total tunnels:", len(farm.Tunnels))
}
