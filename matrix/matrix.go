package matrix

type Node struct {
	guid   string
	degree int
}

type Edge struct {
	guid string
	from Node
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

func CreateGraph(content map[string]map[string]*[]string, degree int, guid string) []Edge {
	nodes := make(map[string]Node)
	edges := []Edge{}
	queue := make([]Node, 0)
	inQueue := make(map[string]bool)
	visited := make(map[string]bool)
	lastDegreeNodes := make(map[string]bool)
	nodes[guid] = Node{guid: guid, degree: 0}
	queue = append(queue, nodes[guid])
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		visited[elem.guid] = true
		nextDegree := elem.degree + 1

		for k, _ := range content[elem.guid] {
			if _, ok := nodes[k]; !ok {
				nodes[k] = Node{guid: k}
			}
			if _, ok := visited[k]; !ok {
				for i := range *content[elem.guid][k] {
					edges = append(edges, Edge{guid: (*content[elem.guid][k])[i], from: nodes[elem.guid], to: nodes[k]})
				}
				if _, ok := inQueue[k]; !ok {
					if nextDegree == degree {
						lastDegreeNodes[k] = true
					} else {
						node := nodes[k]
						node.degree = nextDegree
						queue = append(queue, node)
						inQueue[k] = true
					}
				}
			}
		}
	}
	tmpLastDegreeNodes := lastDegreeNodes
	for keyLastNode, _ := range tmpLastDegreeNodes {
		for k, _ := range content[keyLastNode] {
			if _, ok := nodes[k]; !ok {
				nodes[k] = Node{guid: k}
			}
			if _, ok := lastDegreeNodes[k]; ok {
				for i := range *content[keyLastNode][k] {
					edges = append(edges, Edge{guid: (*content[keyLastNode][k])[i], from: nodes[keyLastNode], to: nodes[k]})
				}
			}
		}
		delete(lastDegreeNodes, keyLastNode)
	}
	return edges
}
