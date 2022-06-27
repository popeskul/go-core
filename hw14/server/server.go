package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-search/hw14/messages_proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"sync"
)

type MessengerServer struct {
	mu   sync.Mutex
	Data []*messages_proto.Message

	messages_proto.MessengerServer
}

func (s *MessengerServer) Messages(_ *messages_proto.Empty, stream messages_proto.Messenger_MessagesServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, message := range s.Data {
		err := stream.Send(message)
		if err != nil {
			return fmt.Errorf("failed to send: %v", err)
		}
	}
	return nil
}

func (s *MessengerServer) Send(_ context.Context, message *messages_proto.Message) (*messages_proto.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Data = append(s.Data, message)
	return message, nil
}

func main() {
	srv := &MessengerServer{}
	messages1 := &messages_proto.Message{Id: int64(uuid.New().ID()), Text: "The Go Programming Language", CreatedAt: timestamppb.Now()}
	messages2 := &messages_proto.Message{Id: int64(uuid.New().ID()), Text: "1984", CreatedAt: timestamppb.Now()}
	srv.Data = append(srv.Data, messages1, messages2)

	l, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	messages_proto.RegisterMessengerServer(grpcServer, srv)
	log.Fatal(grpcServer.Serve(l))
}
