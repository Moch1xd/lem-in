package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) (AntFarm, error) {
	var farm AntFarm

	file, err := os.Open(filename)
	if err != nil {
		return farm, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	num, err := strconv.Atoi(firstLine)
	if err != nil {
		return farm, err
	}
	farm.NumberOfAnts = num

	startRoom := false
	endRoom := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" {
			startRoom = true
		} else if line == "##end" {
			endRoom = true
		} else if strings.Contains(line, "-") {
			pieces := strings.Split(line, "-")
			if len(pieces) != 2 {
				return farm, fmt.Errorf("invalid tunnel definition: %s", line)
			}
			tunnel := Tunnel{Room1: pieces[0], Room2: pieces[1]}
			farm.Tunnels = append(farm.Tunnels, tunnel)
		} else {
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return farm, fmt.Errorf("invalid room definition: %s", line)
			}
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return farm, fmt.Errorf("invalid room coordinates: %s", line)
			}
			room := Room{Name: parts[0], X: x, Y: y}
			farm.Rooms = append(farm.Rooms, room)
			if startRoom {
				farm.StartRoom = room
				startRoom = false
			} else if endRoom {
				farm.EndRoom = room
				endRoom = false
			}
		}
	}

	return farm, nil
}