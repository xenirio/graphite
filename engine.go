package main

import (
	"fmt"
	"bufio"
	"encoding/csv"
	"os"
	"strings"
	"github.com/xenirio/graphite/matrix"
	)

func main() {
	f, _ := os.Open("D:\\Projects\\Golang\\bin\\sample\\relationships.csv");
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	x := matrix.Create(result)
	var origin string
	fmt.Printf("Entity Guid : ")
	fmt.Scanf("%s", &origin)
	output := matrix.CreateCompleteGraph(x, strings.ToUpper(origin))
	for k1, v1 := range output {
		for k2, v2 := range v1 {
			fmt.Printf("[%s][%s] = %s\n", k1, k2, v2)
		}
	}
}