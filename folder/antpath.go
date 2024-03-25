package folder

import (
	"fmt"
	"os"
)

func (g *Graph) SendAnts(n int, comb [][]string) {
	fmt.Println(allInputData)
	fmt.Println()
	for _, path := range comb {
		if len(path) == 1 {
			for i := 1; i <= n; i++ {
				fmt.Print("L", i, "-", g.EndRoom, " ")
			}
			fmt.Println()
			os.Exit(0)
		}
	}

	var combsAnts [][]string
	combsAnts = append(combsAnts, comb...)
	ants := make([]Ant, n)
	i := 0
	for i < n {
		for j := range combsAnts {
			var combToCompareWith []string
			if j == len(combsAnts)-1 {
				combToCompareWith = combsAnts[0]
			} else {
				combToCompareWith = combsAnts[j+1]
			}
			for len(combsAnts[j]) <= len(combToCompareWith) {
				if i == n {
					break
				}
				combsAnts[j] = append(combsAnts[j], "Ant")
				ants[i].Path = comb[j]
				ants[i].Id = i
				i++
			}
		}
	}

	exit := false
	busy := make(map[string]bool)
	for !exit {
		for id, ant := range ants {
			if ant.Skip {
				continue
			}
			room := ant.Path[ant.RoomId]
			if busy[room] {
				if id == n-1 {
					fmt.Println()
				}
				continue
			}
			fmt.Print("L", ant.Id+1, "-", room, " ")
			if id == n-1 {
				fmt.Println()
				if room == g.EndRoom {
					exit = true
				}
			}
			ants[id].RoomId++
			busy[ants[id].Previous] = false
			if room != g.EndRoom {
				busy[room] = true
				ants[id].Previous = room
			}
			if room == g.EndRoom {
				ants[id].Skip = true
			}
		}
	}
}
