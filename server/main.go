package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "org/tik/hyper-queue-service/.gen/agent/queue/circular/agent"
	cll "org/tik/hyper-queue-service/circularLinkedList"
	"org/tik/hyper-queue-service/utils"
)

var (
	port = flag.Int("port", 5050, "The server port")
)

type server struct {
	pb.UnimplementedAgentQueueServer
	list cll.CircularLinkedList[pb.GAgent]
}

func (s *server) Insert(ctx context.Context, agent *pb.GAgent) (*pb.Empty, error) {
	log.Printf("add agent: %v", agent)
	s.list.Insert(*agent)
	return &pb.Empty{}, nil
}

func (s *server) Poll(ctx context.Context, empty *pb.Empty) (*pb.GAgent, error) {
	a := s.list.Poll()
	log.Printf("poll agent: %v", &a)
	return &a, nil
}
func (s *server) List(empty *pb.Empty, stream pb.AgentQueue_ListServer) error {
	arr := s.list.ToArray()
	for _, agent := range arr {
		if err := stream.Send(&agent); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) DeleteById(ctx context.Context, request *pb.DeleteByIdRequest) (*pb.Empty, error) {
	log.Printf("delete agent: %v", request.Id)
	s.list.DeleteByComparator(func(value any) bool {
		return utils.IsEqual(value, request.Id)
	})
	return &pb.Empty{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAgentQueueServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
