package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"pay-service/pkg/disbursement/pb"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pay-service/pkg/service"

	"pay-service/models"
	"pay-service/pkg/kafka"
)

func (s *Server) init() {
	checkErr := func(name string, err error) {
		if err != nil {
			log.Fatalf("init %s failed: %+v", name, err)
		}
	}

	checkErr("database", s.initDB())
	checkErr("kafka.producer", s.initKafkaProducer())

	s.initGRPC()

	s.initService()

	s.initRouter()

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
	publisherService := service.NewPublisher(s.producer)
	s.handle.PublisherService = publisherService

	// init GRPC client
	clientService := service.NewClient(s.client)
	s.handle.Client = clientService

}

func (s *Server) initRouter() error {
	router := gin.Default()

	// Serve frontend static files
	router.LoadHTMLGlob("views/*.html")
	router.Static("static", "static")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	if sf := static.LocalFile("views", true); sf != nil {
		router.Use(static.Serve("/", sf))
	}

	api := router.Group("/api")
	{
		api.GET("/orders/:id", s.handle.GetOrderById)
		api.POST("/confirm", s.handle.Confirm)
	}
	s.gin = router

	return nil
}

func (s *Server) initGRPC() error {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error when dial %v", err)
	}
	//defer cc.Close()
	client := pb.NewDisbursementClient(cc)
	s.client = client
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
