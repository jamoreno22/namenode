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

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	client := name.NewNameNodeClient(conn)

	//Tiene que ser un stream pero no se cómo hacerlo :v
	//prop := gral.Proposal{Ip: "8000", Chunk: &gral.Chunk{Name: "Chunk1", Data: []byte("ABC€")}}
	//book := gral.Book{Name: "Libro1", Parts: 3}
	proposals := []name.Proposal{}
	proposals = append(proposals, name.Proposal{Ip: "8000", Chunk: &name.Chunk{Name: "Chunk1", Data: []byte("ABC€")}})
	proposals = append(proposals, name.Proposal{Ip: "8001", Chunk: &name.Chunk{Name: "Chunk2", Data: []byte("ABC2")}})

	runSendProposal(client, proposals)

	//resp, err := client.GetBookInfo(context.Background(), &book)
	//if err != nil {
	//	log.Fatalf("Did not connect: %s", err)
	//}
	//log.Println(resp)

}

/*

//enviar propuesta al servidor
func runWriteLog(nc gral.NameNodeClient, proposals []gral.Proposal) error {
	log.Println("Inicio de stream writeLog")
	stream, err := nc.WriteLog(context.Background())
	if err != nil {
		log.Printf("Error de stream writeLog: %v", err)
	}

	for _, prop := range proposals {
		if err := stream.Send(&prop); err != nil {
			log.Println("Error al enviar prop")
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error reception response: %v", err)
	}
	log.Printf("Route summary: %v", reply)
	return nil
}
*/

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
