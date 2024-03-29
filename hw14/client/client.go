package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-search/hw14/messages_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	client := messages_proto.NewMessengerClient(conn)

	input := bufio.NewScanner(os.Stdin)

	go func() {
		fmt.Println("Please enter a message:")
		for input.Scan() {
			message := &messages_proto.Message{Id: int64(uuid.New().ID()), Text: input.Text(), CreatedAt: timestamppb.Now()}
			_, err = client.Send(context.Background(), message)
			if err != nil {
				log.Fatal("failed to send message: ", err)
			}

			err = printAllMessagesOnServer(client)
			if err != nil {
				log.Fatal("failed to receive messages: ", err)
				return
			}

			fmt.Println("Please enter a message:")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")
}

func printAllMessagesOnServer(client messages_proto.MessengerClient) error {
	fmt.Println("I'm requesting messages from a gRPC server.")

	stream, err := client.Messages(context.Background(), &messages_proto.Empty{})
	if err != nil {
		return fmt.Errorf("failed to request messages: %v", err)
	}

	for {
		book, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive message: %v", err)
		}
		fmt.Printf("Book: %+v\n", book)
	}
	return nil
}
