package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func roomExists(rooms []Room, name string) bool {
	for _, r := range rooms {
		if r.Name == name {
			return true
		}
	}
	return false
}

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
	if num <= 0 {
		return farm, fmt.Errorf("invalid number of ants: %d", num)
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
		} else if strings.HasPrefix(line, "#") {
			continue
		} else if strings.Contains(line, "-") {
			pieces := strings.Split(line, "-")
			if len(pieces) != 2 {
				return farm, fmt.Errorf("invalid tunnel definition: %s", line)
			}
			if pieces[0] == pieces[1] {
				return farm, fmt.Errorf("tunnel cannot connect a room to itself: %s", line)
			}
			tunnel := Tunnel{Room1: pieces[0], Room2: pieces[1]}
			for _, existingTunnel := range farm.Tunnels {
				if (existingTunnel.Room1 == tunnel.Room1 && existingTunnel.Room2 == tunnel.Room2) ||
					(existingTunnel.Room1 == tunnel.Room2 && existingTunnel.Room2 == tunnel.Room1) {
					return farm, fmt.Errorf("duplicate tunnel: %s", line)
				}
			}
			farm.Tunnels = append(farm.Tunnels, tunnel)
		} else {
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return farm, fmt.Errorf("invalid room definition: %s", line)
			}
			if strings.HasPrefix(parts[0], "L") {
				return farm, fmt.Errorf("room name cannot start with 'L': %s", line)
			}
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return farm, fmt.Errorf("invalid room coordinates: %s", line)
			}
			room := Room{Name: parts[0], X: x, Y: y}
			for _, existingRoom := range farm.Rooms {
				if existingRoom.Name == room.Name {
					return farm, fmt.Errorf("duplicate room name: %s", room.Name)
				}
			}
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

	for _, tunnel := range farm.Tunnels {
		if !roomExists(farm.Rooms, tunnel.Room1) || !roomExists(farm.Rooms, tunnel.Room2) {
			return farm, fmt.Errorf("tunnel connects to non-existent room: %s-%s", tunnel.Room1, tunnel.Room2)
		}
	}

	if farm.StartRoom.Name == "" {
		return farm, fmt.Errorf("start room not defined")
	}

	if farm.EndRoom.Name == "" {
		return farm, fmt.Errorf("end room not defined")
	}

	return farm, nil
}
