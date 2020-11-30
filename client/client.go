package main

import (
	"context"
	"io"
	"log"

	name "github.com/jamoreno22/namenode/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("10.10.28.20:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	client := name.NewNameNodeClient(conn)

	proposals := []name.Proposal{}
	proposals = append(proposals, name.Proposal{Ip: "8000", Chunk: &name.Chunk{Name: "Chunk1", Data: []byte("ABC€")}})
	proposals = append(proposals, name.Proposal{Ip: "8001", Chunk: &name.Chunk{Name: "Chunk2", Data: []byte("ABC2")}})

	runSendProposal(client, proposals)

	bookName := "Mujercitas-Alcott_Louisa_May.pdf"

	runGetChunkDistribution(client, &name.Message{Text: bookName})

	resp, err := client.GetBookInfo(context.Background(), &name.Book{Name: bookName, Parts: 3})
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	log.Println(resp)

}

func runGetChunkDistribution(nc name.NameNodeClient, bookName *name.Message) ([]name.Proposal, error) {
	stream, err := nc.GetChunkDistribution(context.Background(), bookName)
	if err != nil {
		log.Printf("%v", err)
	}
	proposals := []name.Proposal{}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			return proposals, nil
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", nc, err)
		}
		proposals = append(proposals, *feature)
	}
}

func runSendProposal(nc name.NameNodeClient, proposals []name.Proposal) error {

	stream, err := nc.SendProposal(context.Background())
	if err != nil {
		log.Println("Error de stream send proposal")
	}

	log.Println("ki voy")
	a := 1
	for _, prop := range proposals {

		if err := stream.Send(&prop); err != nil {
			log.Println("error al enviar chunk")
			log.Fatalf("%v.Send(%d) = %v", stream, a, err)
		}
		a = a + 1
	}
	for {
		log.Println("ki voy")

		in, err := stream.Recv()
		if err == io.EOF {
			// read done.
			//DistributeChunks()
			log.Printf("weno")
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive a proposal : %v", err)
		}
		log.Printf("Got a proposal ip :%s ", in.Ip)
	}
}
