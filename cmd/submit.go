package cmd

import (
    "fmt"
	"context"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	pb "github.com/jiangplus/luban/rpc"

)

func init() {
	rootCmd.AddCommand(submitCmd)
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)


var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit luban job to server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("luban job")


		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())

		r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	},
}
