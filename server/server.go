package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
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
	namelis, err := net.Listen("tcp", "10.10.28.20:9000")
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

// SendProposal
func (s *nameNodeServer) SendProposal(srv name.NameNode_SendProposalServer) error {
	for {
		prop, err := srv.Recv()
		if err == io.EOF {
			log.Printf("EOF")

			// Agregar la función para chequear la propuesta para la distribucion centralizada
			props, err2 := generateproposal(receivedProposal)
			if err2 != nil {
				log.Printf("Oh no!: %v", err2)
			}
			s.WriteLog(props, len(props), infoBook.Name)
			for _, p := range props {
				if err3 := srv.Send(&p); err3 != nil {
					log.Printf("%v", err3)
				}
				log.Printf("Enviando propuesta al datanode")
			}
			return io.EOF
		}
		if err != nil {
			return err
		}
		receivedProposal = append(receivedProposal, *prop)
	}
}

// GetChunkDistribution
func (s *nameNodeServer) GetChunkDistribution(req *name.Message, srv name.NameNode_GetChunkDistributionServer) error {
	file, err := os.Open("Log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberOfLines := 0
	proposals := []name.Proposal{}

	for scanner.Scan() {
		if numberOfLines > 0 {
			line := strings.Split(scanner.Text(), " ")
			proposals = append(proposals, name.Proposal{Ip: line[1], Chunk: &name.Chunk{Name: line[0]}})
		}
		if scanner.Text()[:len(req.GetText())-1] == req.GetText() {
			numberOfLines, err = strconv.Atoi(scanner.Text()[len(scanner.Text()) : len(scanner.Text())-1])
			if err != nil {
				log.Printf("%v", err)
			}
		}

	}

	//stream
	for _, prop := range proposals {
		if err := srv.Send(&prop); err != nil {
			return err
		}
	}
	return nil
}

// GetBookInfo
func (s *nameNodeServer) GetBookInfo(ctx context.Context, req *name.Book) (*name.Message, error) {
	infoBook = *req
	return &name.Message{Text: "holi"}, nil
}

// Writelog
func (s *nameNodeServer) WriteLog(sP []name.Proposal, parts int, nameBook string) error {
	// create log
	f, err := os.Create("Log.txt")
	if err != nil {
		log.Fatal(err)
	}
	// saved Proposals array

	// Crear Log
	f.WriteString(nameBook + " " + strconv.Itoa(parts) + "\n")
	for _, prop := range sP {
		_, err2 := f.WriteString(prop.Chunk.Name + " " + prop.Ip + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	defer f.Close()
	return nil
}

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
			propResponse = append(propResponse, name.Proposal{Ip: gIps[rand.Intn(len(gIps))], Chunk: prop.Chunk})
		} else {
			propResponse = append(propResponse, prop)
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
