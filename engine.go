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
	f, _ := os.Open("D:\\Projects\\Golang\\bin\\sample\\relationships.csv")
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
	results := matrix.CreateGraph(relationMap, degree, strings.ToUpper(origin))
	now2 := time.Now()
	fmt.Println(results)
	fmt.Printf("Total : %d\n", len(results))
	fmt.Printf("Time : %s\n", now2.Sub(now1))
}
