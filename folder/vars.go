package folder

type Graph struct {
	Ant             int
	Ants            []Ant
	StartRoom       string
	EndRoom         string
	AdjacencyList   map[string][]string
	SolvedPaths     [][]string
	PathGroups      map[int][][]string
	BestCombination [][]string
}

type Ant struct {
	Id       int
	Skip     bool
	Path     []string
	RoomId   int
	Previous string
}
