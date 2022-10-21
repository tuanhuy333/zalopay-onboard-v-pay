package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"V_Pay_Onboard_Program/pkg/service"

	"V_Pay_Onboard_Program/models"
	"V_Pay_Onboard_Program/pkg/kafka"
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
