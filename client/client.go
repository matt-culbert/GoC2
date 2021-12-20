package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blackhat-go/bhg/ch-14/grpcapi"
	"google.golang.org/grpc"
)

func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client grpcapi.AdminClient
	)

	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", 9090), opts...); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = grpcapi.NewAdminClient(conn)

	var cmd = new(grpcapi.Command)
	cmd.In = os.Args[1]
	// gotta add logic here to determine if cmd is RunCommand or ListBeacons
	// if strings.hasSuffix(cmd, "1") then strings.Strip(cmd, "1") ListBeacons
	ctx := context.Background()
	cmd, err = client.RunCommand(ctx, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Out)
}
