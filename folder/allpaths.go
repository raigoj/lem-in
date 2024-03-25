package folder

var (
	possiblePath []string
	isVisited    = make(map[string]bool)
)

func (g *Graph) SearchForAllPaths() {
	g.DFSrecursice(g.StartRoom)
}

func (g *Graph) DFSrecursice(roomName string) {
	isVisited[roomName] = true
	possiblePath = append(possiblePath, roomName)
	if roomName == g.EndRoom {
		var copy []string
		copy = append(copy, possiblePath[1:]...)
		g.SolvedPaths = append(g.SolvedPaths, copy)
	} else {
		var neighbours []string = g.AdjacencyList[roomName]
		for _, val := range neighbours {
			if isVisited[val] != true {
				g.DFSrecursice(val)
			}
		}
	}
	possiblePath = possiblePath[:len(possiblePath)-1]
	isVisited[roomName] = false
}
