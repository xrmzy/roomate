package delivery

import (
	"fmt"
	"log"
	"roomate/config"
	"roomate/delivery/controller"
	"roomate/delivery/middleware"
	"roomate/manager"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager  manager.UseCaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
}

func (s *Server) setupController() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	rg := s.engine.Group("/api/v1")
	// Register all controller in here
	controller.NewUserController(s.ucManager.UserUsecase(), rg).Route()
	controller.NewRoleController(s.ucManager.RoleUsecase(), rg).Route()
	controller.NewCustomerController(s.ucManager.CustomerUseCase(), rg).Route()
	controller.NewRoomController(s.ucManager.RoomUseCase(), rg).Route()
	controller.NewServiceController(s.ucManager.ServiceUseCase(), rg).Route()
	controller.NewBookingController(s.ucManager.BookingUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.FileConfig)

	return &Server{
		ucManager:  usecaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
	}
}
