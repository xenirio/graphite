package matrix

type Node struct {
	guid   string
	degree int
	edges  []Edge
}

type Edge struct {
	guid string
	to   Node
}

func Create(content [][]string) map[string]map[string]*[]string {
	table := make(map[string]map[string]*[]string)
	for i := range content {
		from := content[i][0]
		to := content[i][1]
		edge := content[i][2]

		if _, ok := table[from]; !ok {
			table[from] = make(map[string]*[]string)
		}
		var p *[]string
		var ok bool
		if p, ok = table[from][to]; !ok {
			p = &[]string{}
			table[from][to] = p

			if _, ok := table[to]; !ok {
				table[to] = make(map[string]*[]string)
			}
			table[to][from] = p
		}
		*p = append(*p, edge)
	}
	return table
}

func CreateGraph(content map[string]map[string]*[]string, degree int, guid string) []string {
	edges := []string{}
	queue := make([]Node, 0)
	inQueue := make(map[string]bool)
	visited := make(map[string]bool)
	lastDegreeNodes := make(map[string]bool)
	queue = append(queue, Node{guid: guid, degree: 0})
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		visited[elem.guid] = true
		nextDegree := elem.degree + 1

		for k, _ := range content[elem.guid] {
			if _, ok := visited[k]; !ok {
				edges = append(edges, *content[elem.guid][k]...)
				if _, ok := inQueue[k]; !ok {
					if nextDegree == degree {
						lastDegreeNodes[k] = true
					} else {
						queue = append(queue, Node{guid: k, degree: nextDegree})
						inQueue[k] = true
					}
				}
			}
		}
	}
	tmpLastDegreeNodes := make([]string, 0, len(lastDegreeNodes))
	for k, _ := range lastDegreeNodes {
		tmpLastDegreeNodes = append(tmpLastDegreeNodes, k)
	}
	for _, v := range tmpLastDegreeNodes {
		for k, _ := range content[v] {
			if _, ok := lastDegreeNodes[k]; ok {
				edges = append(edges, *content[v][k]...)
			}
		}
		delete(lastDegreeNodes, v)
	}
	return edges
}
