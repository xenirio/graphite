package matrix

// Node Type
type Node struct {
	Guid   string
	Degree int
}

// Edge Type
type Edge struct {
	Guid string
	From Node
	To   Node
}

// Create Create Relationship Matrix
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

// CreateSimpleGraph Create Simple Graph Map
func CreateSimpleGraph(content map[string]map[string]*[]string, degree int, guid string) ([]Edge, map[string]Node) {
	nodes := make(map[string]Node)
	edges := []Edge{}
	queue := make([]Node, 0)
	inQueue := make(map[string]bool)
	visited := make(map[string]bool)
	lastDegreeNodes := make(map[string]Node)
	nodes[guid] = Node{Guid: guid, Degree: 0}
	queue = append(queue, nodes[guid])
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		visited[elem.Guid] = true
		nextDegree := elem.Degree + 1

		for k := range content[elem.Guid] {
			if _, ok := nodes[k]; !ok {
				nodes[k] = Node{Guid: k}
			}
			if _, ok := visited[k]; !ok {
				for i := range *content[elem.Guid][k] {
					edges = append(edges, Edge{Guid: (*content[elem.Guid][k])[i], From: nodes[elem.Guid], To: nodes[k]})
				}
				if _, ok := inQueue[k]; !ok {
					if nextDegree == degree {
						lastDegreeNodes[k] = nodes[k]
					} else {
						node := nodes[k]
						node.Degree = nextDegree
						queue = append(queue, node)
						inQueue[k] = true
					}
				}
			}
		}
	}
	return edges, lastDegreeNodes
}

// FindInterconnectedEdges to find interconnection edge
func FindInterconnectedEdges(content map[string]map[string]*[]string, nodes map[string]Node) []Edge {
	edges := []Edge{}
	tmpNodes := nodes
	for keyLastNode := range tmpNodes {
		for k := range content[keyLastNode] {
			if _, ok := nodes[k]; ok {
				for i := range *content[keyLastNode][k] {
					edges = append(edges, Edge{Guid: (*content[keyLastNode][k])[i], From: nodes[keyLastNode], To: nodes[k]})
				}
			}
		}
		delete(nodes, keyLastNode)
	}
	return edges
}

// CreateInterconnectionGraph to create interconnection Graph
func CreateInterconnectionGraph(content map[string]map[string]*[]string, degree int, guids map[string]bool) map[string]bool {
	selectedNodes := make(map[string]bool)

	for startNode := range guids {
		guids[startNode] = false
		findPath(content, make(map[string]bool), selectedNodes, startNode, guids, degree, 0)
		guids[startNode] = true
	}

	selectedRelationships := make(map[string]bool)
	var listNode []string

	for startNode := range selectedNodes {
		listNode = append(listNode, startNode)
	}

	for i := 0; i < len(listNode)-1; i++ {
		for j := i + 1; j < len(listNode); j++ {
			if _, ok := content[listNode[i]][listNode[j]]; ok {
				for _, rel := range *content[listNode[i]][listNode[j]] {
					selectedRelationships[rel] = true
				}
			}
		}
	}

	return selectedRelationships
}

func findPath(content map[string]map[string]*[]string, visited map[string]bool, answer map[string]bool, current string, targets map[string]bool, maxDegree int, currentDegree int) {
	visited[current] = true
	if currentDegree == maxDegree {
		isTarget, exist := targets[current]
		if exist && isTarget {
			for k := range visited {
				answer[k] = true
			}
		}
	} else {
		isTarget, exist := targets[current]
		if exist && isTarget {
			for k := range visited {
				answer[k] = true
			}
		} else {
			for neighbor := range content[current] {
				if _, ok := visited[neighbor]; !ok {
					findPath(content, visited, answer, neighbor, targets, maxDegree, currentDegree+1)
				}
			}
		}
	}
	delete(visited, current)
}
