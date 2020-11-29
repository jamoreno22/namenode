package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	name "github.com/jamoreno22/namenode/pkg/proto"
	"google.golang.org/grpc"
)

type nameNodeServer struct {
	name.UnimplementedNameNodeServer
}

type nameNodeSendProposalServer struct {
	grpc.ServerStream
}

var infoBook = name.Book{}

var receivedProposal []name.Proposal

func main() {

	// create a listener on TCP port 8000
	namelis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	ns := nameNodeServer{}                           // create a gRPC server object
	grpcNameServer := grpc.NewServer()               // attach the Ping service to the server
	name.RegisterNameNodeServer(grpcNameServer, &ns) // start the server

	log.Println("NameServer running ...")
	if err := grpcNameServer.Serve(namelis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// - - - - - - - - - - NameNode Server functions - - - - - - - - - - -

// Writelog server
func (s *nameNodeServer) WriteLog(sP []name.Proposal, parts int32, nameBook string) error {
	// create log
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// saved Proposals array

	for _, prop := range sP {

		// Aquí va el código para guardar el log (falta formato)
		_, err2 := f.WriteString("ip " + prop.Ip)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	return nil
}

/*func (s *nameServer) GetBookInfo(ctx context.Context, book *gral.Book) (*gral.Message, error) {

	infoBook = *book
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}*/

//Si quieres algo bien hecho tienes que hacerlo tu mismo
func generateproposal(props []name.Proposal) ([]name.Proposal, error) {
	ips := []string{"10.10.28.17:9000", "10.10.28.18:9000", "10.10.28.19:9000"}
	var propResponse []name.Proposal
	var gIps []string //ipes wenas
	for _, ip := range ips {
		if pingDataNode(ip) {
			gIps = append(gIps, ip)
		}
	}
	if len(gIps) == 0 {
		return nil, fmt.Errorf("There's no posible proposal")
	}

	for _, prop := range props {
		if !stringInSlice(prop.Ip, gIps) {
			propResponse = append(propResponse, name.Proposal{Ip: gIps[0], Chunk: prop.Chunk})
		}
	}
	return propResponse, nil
}

func pingDataNode(ip string) bool {
	timeOut := time.Duration(10 * time.Second)
	_, err := net.DialTimeout("tcp", ip, timeOut)
	if err != nil {
		return false
	}
	return true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *nameNodeServer) SendProposal(srv name.NameNode_SendProposalServer) error {
	for {
		prop, err := srv.Recv()

		if err == io.EOF {
			log.Printf("EOF")
			props, err2 := generateproposal(receivedProposal)
			if err2 != nil {
				log.Printf("Oh no!: %v", err2)
			}
			s.WriteLog(props, int32(len(props)), "inserte nombre aqui")
			for _, p := range props {
				if err3 := srv.Send(&p); err3 != nil {
					return err3
				}
			}
		}
		if err != nil {
			return err
		}
		receivedProposal = append(receivedProposal, *prop)
	}
}
