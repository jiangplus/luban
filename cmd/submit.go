package cmd

import (
    "fmt"
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	pb "github.com/jiangplus/luban/rpc"

)

func init() {
	rootCmd.AddCommand(submitCmd)
}

const (
	address = "localhost:50051"
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
		c := pb.NewLubanClient(conn)

		// Contact the server and print out its response.

		fmt.Println(args)
		filename := ""
		if len(args) < 1 {
			log.Fatalln("filename is needed")
		}
		filename = args[0]

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}
		if len(data) == 0 {
			log.Fatalln("input file is empty")
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Submit(ctx, &pb.SubmitRequest{Data: string(data)})
		if err != nil {
			log.Fatalf("could not submit job: %v", err)
		}
		log.Printf("Data: %s", r.GetData())
	},
}
