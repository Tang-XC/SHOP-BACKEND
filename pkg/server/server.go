package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"shop/pkg/common"
	"shop/pkg/config"
	"shop/pkg/controller"
	database "shop/pkg/db"
	"shop/pkg/middleware"
	"shop/pkg/repository"
	"shop/pkg/service"
	"shop/pkg/utils/set"
	"syscall"
	"time"
)

type Server struct {
	engine      *gin.Engine
	config      *config.Config
	logger      *logrus.Logger
	repository  repository.Repository
	controllers []controller.Controller
}

// 关闭服务
func (s *Server) Close() {
	//关闭其他服务
}

// 运行服务
func (s *Server) Run() error {
	defer s.Close()

	s.initRouter()
	addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}
	s.logger.Infof("Start server on %s", addr)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatalf("Failed to start server,%v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Server.GracefulShutdownPeriod)*time.Second)
	defer cancel()

	ch := <-sig
	s.logger.Infof("Receive signal:%s", ch)
	return server.Shutdown(ctx)
}

// 初始化路由
func (s *Server) initRouter() {
	root := s.engine
	root.GET("/", common.WrapFunc(s.getRoutes))
	root.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := root.Group("/api/v1")
	controllers := make([]string, 0, len(s.controllers))
	for _, router := range s.controllers {
		router.RegisterRoute(api)
		controllers = append(controllers, router.Name())
	}
	logrus.Infof("Server enabled controllers: %v", controllers)
}

// 获取所有路由
func (s *Server) getRoutes() []string {
	paths := set.NewString()
	for _, r := range s.engine.Routes() {
		if r.Path != "" {
			paths.Insert(r.Path)
		}
	}
	return paths.Slice()
}

// 创建服务
func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	//建立数据库连接
	db, err := database.NewDB(&conf.DB)
	if err != nil {
		return nil, err
	}

	//创建数据访问层
	repository := repository.NewRepository(db)
	//迁移数据库，创建表
	if err := repository.Migrate(); err != nil {
		return nil, err
	}
	//初始化数据访问层
	if err := repository.Init(); err != nil {
		return nil, err
	}

	//创建业务逻辑层
	userService := service.NewUserService(repository.User())
	authService := service.NewAuthService(repository.User())
	roleService := service.NewRoleService(repository.Role())
	permissionService := service.NewPermissionService(repository.Permission())
	//创建表示层
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService, authService)
	roleController := controller.NewRoleController(roleService)
	permissionController := controller.NewPermissionController(permissionService)

	controllers := []controller.Controller{userController, authController, roleController, permissionController}
	e := gin.Default()
	e.Use(
		middleware.CORSMiddleware(),
	)
	return &Server{
		engine:      e,
		config:      conf,
		logger:      logger,
		controllers: controllers,
		repository:  repository,
	}, nil
}
