package matrix

//import ("fmt")

func Create(content [][]string) map[string]map[string][]string {
	table := make(map[string]map[string][]string)
    for i := range content {
    	record := content[i]
		if _, ok := table[record[0]]; !ok {
			table[record[0]] = make(map[string][]string)
		}
		if _, ok := table[record[0]][record[1]]; !ok {
			table[record[0]][record[1]] = []string{}
		}
		table[record[0]][record[1]] = append(table[record[0]][record[1]], record[2])
	}
	return table
}

func CreateGraph(mapContent map[string]map[string][]string, guid string) map[string]map[string][]string {
	results := map[string]map[string][]string {
		guid: mapContent[guid],
	}
	for k1, _ := range mapContent {
		if val, ok := mapContent[k1][guid]; ok{
			results[k1] = map[string][]string {
				guid: val,
			}
		}
	}
	return results
}

func CreateCompleteGraph(content map[string]map[string][]string, guid string) map[string]map[string][]string {
	degree1 := CreateGraph(content, guid)
	var originGuids = make(map[string]bool)
	for k1, _ := range degree1 {
		if _, ok := degree1[k1]; ok{
			originGuids[k1] = true
		}
		for k2, _ := range degree1[k1] {
			if _, ok := degree1[k1][k2]; ok{
				originGuids[k2] = true
			}
		}
	}
	var childs = make(map[string]map[string][]string)
	for k3, _ := range originGuids {
		results := CreateGraph(content, k3)
		for k4, _ := range results {
			if _, ok := originGuids[k4]; ok {
				if _, ok := childs[k4]; !ok {
					childs[k4] = make(map[string][]string)
				}
				for k5, _ := range results[k4] {
					if _, ok := originGuids[k5]; ok {
						if _, ok := childs[k4][k5]; !ok {
							childs[k4][k5] = []string{}
						}
						elems := []string{}
						for i := 0; i < len(results[k4][k5]); i++{
							isDuplicate := false
							for j := 0; j < len(childs[k4][k5]); j++{
								if results[k4][k5][i] == childs[k4][k5][j] {
									isDuplicate = true
									break
								}
							}
							if !isDuplicate {
								elems = append(elems, results[k4][k5][i])
							}
						}
						childs[k4][k5] = append(childs[k4][k5], elems...)
					}
				}
			}
		}
	}
	return childs;
}