package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xenirio/graphite/matrix"
)

func main() {
	f, _ := os.Open("D:\\Projects\\Golang\\src\\github.com\\xenirio\\graphite\\sample\\relationships.csv")
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	relationMap := matrix.Create(result)
	var origin string
	var degree int
	fmt.Print("Entity Guid : ")
	fmt.Scan(&origin)
	fmt.Print("Degree : ")
	fmt.Scan(&degree)
	now1 := time.Now()
	/*edges, lastNodes := matrix.CreateSimpleGraph(relationMap, degree, strings.ToUpper(origin))
	edges = append(edges, matrix.FindInterconnectedEdges(relationMap, lastNodes)...)*/
	mapOrigins := make(map[string]bool)
	for _, v := range strings.Split(origin, ",") {
		mapOrigins[strings.ToUpper(v)] = true
	}
	edges := matrix.CreateInterconnectionGraph(relationMap, degree, mapOrigins)
	now2 := time.Now()
	fmt.Println(edges)
	fmt.Printf("Total : %d\n", len(edges))
	fmt.Printf("Time : %s\n", now2.Sub(now1))
}
