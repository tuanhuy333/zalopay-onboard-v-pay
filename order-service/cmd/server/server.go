package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import underlying database driver
	"github.com/jmoiron/sqlx"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	_ "gorm.io/driver/mysql"

	"order-service/pkg/handlers"
	"order-service/pkg/kafka"
	"order-service/pkg/service"
)

type Server struct {
	config      Config
	http        *http.Server
	grpc        *grpc.Server
	grpcService *service.GRPCService
	db          *sqlx.DB
	gormDB      *gorm.DB
	handle      handlers.Handler
	consumer    kafka.Consumer
	producer    kafka.Producer
	gin         *gin.Engine
}

type Config struct {
	Env      string
	Database struct {
		Addr string
		Name string
		User string
		Pass string
	}
	Kafka struct {
		GroupID string

		Orders struct {
			Brokers []string
			Topic   string
		}
		Payment struct {
			Brokers []string
			Topic   string
		}
	}
}

func NewServer(c Config) *Server {
	return &Server{config: c}
}

func (s *Server) Start() {
	s.init()
	s.start()
}

func (s *Server) startConsumers() {
	fmt.Println("Start Consumer")
	// start consumer
	go func() {
		err := s.consumer.Start()
		if err != nil && !errors.Is(err, kafka.ErrConsumerGroupClosed) {
			log.Errorf(nil, "stop consumer:%v", err)
		}
	}()
}
func (s *Server) startHTTP() {
	fmt.Println("Start HTTP")
	go func() {
		s.gin.Run(":8099")
		return
	}()
}
func (s *Server) startGRPC() {
	fmt.Println("Start GRPC")

	addr := "localhost:9090"
	if s.config.Env != "local" {
		addr = ":9090"
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf(nil, "listen grpc: %v", err)
	}
	if err = s.grpc.Serve(l); err != nil {
		log.Errorf(nil, "serve grpc: %v", err)
	}
}
func (s *Server) start() {
	s.startConsumers()
	s.startHTTP()
	s.startGRPC()
}

func (s *Server) Shutdown() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
