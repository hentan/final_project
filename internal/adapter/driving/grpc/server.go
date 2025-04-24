package grpc

import (
	"context"
	v1 "github.com/hentan/final_project/internal/adapter/driving/grpc/v1"
	"github.com/hentan/final_project/internal/logger"
	"github.com/hentan/internal_api_books/gen/go/books"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
)

type Server struct {
	server     *grpc.Server
	port       string
	errors     chan error
	bookServer books.BookServiceServer
}

func NewGRPCServer(ctx context.Context, port string) *Server {
	grpcServer := grpc.NewServer()
	service := v1.NewBookstoreServer()
	reflection.Register(grpcServer)

	return &Server{
		server:     grpcServer,
		port:       port,
		errors:     make(chan error),
		bookServer: service,
	}
}

func (s *Server) Start() {
	newLogger := logger.GetLogger()

	lis, err := net.Listen("tcp", "0.0.0.0:"+s.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	books.RegisterBookServiceServer(s.server, s.bookServer)

	newLogger.Info("Listening on port 50051")
	if err := s.server.Serve(lis); err != io.EOF {
		log.Fatalf("failed to serve: %v", err)
	}
}

type BookServiceServer struct {
	books.UnimplementedBookServiceServer
}
