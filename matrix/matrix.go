package matrix

type Node struct {
	Guid   string
	Degree int
}

type Edge struct {
	Guid string
	From Node
	To   Node
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

func CreateGraph(content map[string]map[string]*[]string, degree int, guid string) ([]Edge, map[string]Node) {
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

		for k, _ := range content[elem.Guid] {
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

func FindInterconnectedEdges(content map[string]map[string]*[]string, nodes map[string]Node) []Edge {
	edges := []Edge{}
	tmpNodes := nodes
	for keyLastNode, _ := range tmpNodes {
		for k, _ := range content[keyLastNode] {
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
