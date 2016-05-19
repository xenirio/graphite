package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/xenirio/graphite/graph"
	"github.com/xenirio/graphite/matrix"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type GraphContext struct{}

var relationMap = make(map[string]map[string]*[]string)

func (s *GraphContext) CreateGraph(origin *pb.Origin, stream pb.Graph_CreateGraphServer) error {
	edges, lastNodes := matrix.CreateGraph(relationMap, int(origin.Degree), strings.ToUpper(origin.Guid))
	edges = append(edges, matrix.FindInterconnectedEdges(relationMap, lastNodes)...)
	for _, e := range edges {
		edge := pb.Edge{Guid: e.Guid, From: e.From.Guid, To: e.To.Guid}
		if err := stream.Send(&edge); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	f, _ := os.Open("D:\\Projects\\Golang\\bin\\sample\\relationships.csv")
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	relationMap = matrix.Create(result)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGraphServer(s, &GraphContext{})
	s.Serve(lis)
	fmt.Println("service searted")
}
