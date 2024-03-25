package folder

func (g *Graph) GetBestComb() {
	minLvl := 999999
	var bestComb [][]string
	for _, comb := range g.PathGroups {
		freeSpace := 0
		totalCombLen := 0
		level := 0
		LongestPath := len(comb[len(comb)-1]) // as we got it sorted by lenght , we take longest one from the end of combination
		for i := 0; i < len(comb); i++ {
			freeSpace = freeSpace + LongestPath - len(comb[i])
			totalCombLen = totalCombLen + len(comb[i])
		}
		level = (g.Ant-freeSpace)/len(comb) + LongestPath // (antsnumber - freespace) / amountofpathsincomb + longhestpath
		if level < minLvl {
			minLvl = level
			bestComb = comb
		}
	}
	g.BestCombination = bestComb
}
