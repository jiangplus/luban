package cmd

import (
	"fmt"
	"context"
    yamlutil "gopkg.in/yaml.v2"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	pb "github.com/jiangplus/luban/rpc"

	. "github.com/jiangplus/luban/core"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedLubanServer
}

func parseSpec(data string) FlowSpec {
	var flowspec FlowSpec

	err := yamlutil.Unmarshal([]byte(data), &flowspec)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
	return flowspec
}

func sort_tasks(tasks []TaskSpec) (bool, []TaskSpec) {
	// toposort
	graph := NewGraph(len(tasks))
	for _, task := range tasks {
		graph.AddNode(task.Name)
	}
	for _, task := range tasks {
		if task.Deps != nil {
			for _, dep_name := range task.Deps {
				graph.AddEdge(task.Name, dep_name)
			}
		}
	}
	result, ok := graph.Toposort()
	if !ok {
		panic("cycle detected")
	}
	sorted_tasks := []TaskSpec{}
	for _, task_name := range result {
		for _, task := range tasks {
			if task.Name == task_name {
				sorted_tasks = append(sorted_tasks, task)
			}
		}
	}
	return ok, sorted_tasks
}

func (s *server) Submit(ctx context.Context, in *pb.SubmitRequest) (*pb.SubmitReply, error) {
	log.Printf("Received: %v", in.GetData())
	flowspec := parseSpec(in.GetData())
	fmt.Println(flowspec)

	tasks := flowspec.Tasks
	ok, sorted_tasks := sort_tasks(tasks)
	if !ok {
		panic("cycle detected")
	}
	fmt.Println(sorted_tasks)

	return &pb.SubmitReply{Data: "OK"}, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start luban server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("luban server started")

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterLubanServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
