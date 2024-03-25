package main

import (
	"lem-in/folder"
)

var g folder.Graph

func main() {
	g.ReadFile()
	g.SearchForAllPaths()
	g.CombineAndSort()
	g.GetBestComb()
	g.SendAnts(g.Ant, g.BestCombination)
}
