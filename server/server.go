package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	gral "github.com/jamoreno22/lab2_dist/pkg/proto"
	"google.golang.org/grpc"
)

type nameServer struct {
	gral.UnimplementedNameNodeServer
}

var path = "Log"

func main() {

	// create a listener on TCP port 8000
	namelis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	ns := nameServer{}                               // create a gRPC server object
	grpcNameServer := grpc.NewServer()               // attach the Ping service to the server
	gral.RegisterNameNodeServer(grpcNameServer, &ns) // start the server

	log.Println("NameServer running ...")
	if err := grpcNameServer.Serve(namelis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// - - - - - - - - - - NameNode Server functions - - - - - - - - - - -

// Writelog server side
func (s *nameServer) WriteLog(wls gral.NameNode_WriteLogServer) error {
	log.Printf("Stream WriteLogServer")
	// create log
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// saved Proposals array
	sP := []gral.Proposal{}

	for {
		prop, err := wls.Recv()
		if err == io.EOF {
			log.Printf("EOF ------------")
			return (wls.SendAndClose(&gral.Message{Text: "Oh no... EOF"}))
		}
		if err != nil {
			return err
		}

		sP = append(sP, *prop)

		// Aquí va el código para guardar el log
		_, err2 := f.WriteString("ip " + prop.Ip)

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
