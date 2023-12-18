package delivery

import (
	"fmt"
	"log"
	"roomate/config"
	"roomate/delivery/controller"
	"roomate/delivery/middleware"
	"roomate/manager"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager     manager.UseCaseManager
	engine        *gin.Engine
	host          string
	logService    common.MyLogger
	auth          usecase.AuthUseCase
	jwtService    common.JwtToken
	gSheetUc      usecase.GSheetUseCase
	gSheetService common.GSheet
	gDriveService common.GDrive
}

func (s *Server) setupController() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/api/v1")
	// Register all controller in here
	controller.NewUserController(s.ucManager.UserUsecase(), rg).Route()
	controller.NewRoleController(s.ucManager.RoleUsecase(), rg).Route()
	controller.NewCustomerController(s.ucManager.CustomerUseCase(), rg).Route()
	controller.NewRoomController(s.ucManager.RoomUseCase(), rg).Route()
	controller.NewServiceController(s.ucManager.ServiceUseCase(), rg).Route()
	controller.NewBookingController(s.ucManager.BookingUseCase(), rg, authMiddleware).Route()
	controller.NewAuthController(s.auth, rg, s.jwtService).Route()
	controller.NewGSheetController(s.gSheetUc, rg, authMiddleware).Route()
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
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.FileConfig)
	jwtService := common.NewJwtToken(cfg.TokenConfig)
	gSheetService := common.NewGSheet(cfg.SheetConfig)
	gDriveService := common.NewGDrive(cfg.SheetConfig)

	return &Server{
		ucManager:     useCaseManager,
		engine:        engine,
		host:          host,
		logService:    logService,
		auth:          usecase.NewAuthUseCase(useCaseManager.UserUsecase(), jwtService),
		jwtService:    jwtService,
		gSheetUc:      usecase.NewGSheetUseCase(repoManager.BookingRepo(), useCaseManager.UserUsecase(), useCaseManager.CustomerUseCase(), gDriveService, gSheetService),
		gSheetService: gSheetService,
		gDriveService: gDriveService,
	}
}
