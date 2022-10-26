package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"order-service/models"

	"order-service/pkg/kafka"
	"order-service/pkg/order/pb"
	"order-service/pkg/service"
)

func (s *Server) init() {
	checkErr := func(name string, err error) {
		if err != nil {
			log.Fatalf("init %s failed: %+v", name, err)
		}
	}

	checkErr("database", s.initDB())
	checkErr("kafka.producer", s.initKafkaProducer())
	s.initService()

	checkErr("kafka.consumer", s.initKafkaConsumer())

	s.initRouter()
	s.initGRPC()

}

func (s *Server) initDB() error {
	log.Println("Init DB")
	cfg := s.config.Database
	ds := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=%s",
		cfg.User, cfg.Pass, cfg.Addr, cfg.Name, "Asia%2fHo_Chi_Minh")

	db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("connect db: %v", err)
	}

	db.AutoMigrate(models.Order{})

	s.gormDB = db

	return nil
}

func (s *Server) initService() {
	log.Println("Init Service")

	handleDB := service.New(s.gormDB)
	s.handle.Service = handleDB

	// init service publisher
	publisherService := service.NewPublisher(s.producer, s.config.Kafka.Orders.Topic)
	s.handle.PublisherService = publisherService

	// init service implement GRPC
	grpcSvc := service.NewGRPCService(handleDB)
	s.grpcService = grpcSvc

}

func (s *Server) initGRPC() error {
	s.grpc = grpc.NewServer()
	pb.RegisterDisbursementServer(s.grpc, s.grpcService)
	return nil
}

func (s *Server) initRouter() error {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/orders", s.handle.CreateOrders)
		api.GET("/orders", s.handle.GetOrders)
		api.GET("/orders/:id", s.handle.GetOrderById)
	}
	s.gin = router

	return nil
}

func (s *Server) initKafkaProducer() error {

	{
		log.Println("Init Kafka Producer")

		cfg := s.config.Kafka.Orders
		p, err := kafka.NewProducer(cfg.Brokers)
		if err != nil {
			return err
		}

		s.producer = p

	}
	return nil
}

func (s *Server) initKafkaConsumer() error {
	log.Println("Init Kafka Consumer")

	cfg := s.config.Kafka.Payment
	cg, err := kafka.NewConsumer(cfg.Brokers, s.config.Kafka.GroupID)
	if err != nil {
		return fmt.Errorf("connect to PE: brokers=%v, topic=%v", cfg.Brokers, cfg.Topic)
	}

	// consume message
	cg.Consume(cfg.Topic, service.NewMessageService(s.handle.Service, s.handle.PublisherService).HandleKafkaMessage)
	s.consumer = cg
	return nil
}
