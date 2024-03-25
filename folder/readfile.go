package folder

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// File reading, data check.
var allInputData string

func (g *Graph) ReadFile() {
	var start bool = false
	var end bool = false
	var allRooms []string
	var allRoomLinks [][]string
	if len(os.Args) != 2 {
		fmt.Println("Error - type: go run main.go maps/example01.txt")
		os.Exit(0)
	}
	inputfile := os.Args[1]
	data, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatalf("Unable to read file %v", err)
	}
	data = []byte(strings.ReplaceAll(string(data), "\r", ""))

	allInputData = string(data)

	lines := strings.Split(string(data), "\n")

	for i, v := range lines {
		if i == 0 {
			g.Ant, _ = strconv.Atoi(v)
			if g.Ant < 1 {
				fmt.Println("Error - no ants in given file")
				os.Exit(1)
			}
			if g.Ant > 9999 {
				ErrHandler()
			}
		}
		if v == "##start" {
			start = true
			g.StartRoom = getRoomName(lines[i+1])
		}
		if v == "##end" {
			end = true
			g.EndRoom = getRoomName(lines[i+1])
		}
		if strings.Contains(v, " ") {
			var roomName string
			roomData := strings.Split(v, " ")
			roomName = roomData[0]

			match := regexp.MustCompile(`^[^L#]\s*`)
			result := match.MatchString(roomName)

			if result {
				allRooms = append(allRooms, roomName)
			} else {
				ErrHandler()
			}
		}
		if strings.Contains(v, "-") {
			linkData := strings.Split(v, "-")
			for _, roomInLink := range linkData {
				if !matched(roomInLink, allRooms) {
					ErrHandler()
				}
			}
			allRoomLinks = append(allRoomLinks, linkData)
		}
	}
	if start != true {
		ErrHandler()
	}
	if end != true {
		ErrHandler()
	}
	if len(allRoomLinks) < 1 {
		ErrHandler()
	}
	for i, _ := range allRooms {
		for j := i + 1; j < len(allRooms); j++ {
			if allRooms[i] == allRooms[j] {
				ErrHandler()
			}
		}
	}
	g.GetAdjList(allRooms, allRoomLinks)
}

func matched(roomInLink string, allRooms []string) bool {
	for _, roomData := range allRooms {
		if roomData == roomInLink {
			return true
		}
	}
	return false
}

func getRoomName(ln string) string {
	splittedLine := strings.Split(ln, " ")
	return splittedLine[0]
}

func (g *Graph) GetAdjList(allRooms []string, allRoomLinks [][]string) {
	if g.AdjacencyList == nil {
		g.AdjacencyList = make(map[string][]string)
	}
	for _, v := range allRooms {
		g.AdjacencyList[v] = []string{}
	}
	for _, y := range allRoomLinks {
		first := y[0]
		second := y[1]
		g.AdjacencyList[first] = append(g.AdjacencyList[first], second)
		g.AdjacencyList[second] = append(g.AdjacencyList[second], first)
	}
}
