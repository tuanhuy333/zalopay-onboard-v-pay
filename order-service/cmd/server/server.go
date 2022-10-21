package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import underlying database driver
	"github.com/jmoiron/sqlx"
	"google.golang.org/appengine/log"
	"gorm.io/gorm"

	_ "gorm.io/driver/mysql"

	"V_Pay_Onboard_Program/pkg/handlers"
	"V_Pay_Onboard_Program/pkg/kafka"
)

type Server struct {
	config   Config
	http     *http.Server
	db       *sqlx.DB
	gormDB   *gorm.DB
	handle   handlers.Handler
	consumer kafka.Consumer
	producer kafka.Producer
	gin      *gin.Engine
}

type Config struct {
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
	// start consumer
	go func() {
		err := s.consumer.Start()
		if err != nil && !errors.Is(err, kafka.ErrConsumerGroupClosed) {
			log.Errorf(nil, "stop pe consumer:%v", err)
		}
	}()
}
func (s *Server) startHTTP() {
	go func() {
		s.gin.Run(":8099")
		return
	}()
}
func (s *Server) start() {
	// start consumer
	s.startConsumers()

	// start http
	s.startHTTP()

}

func (s *Server) Shutdown() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
