package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"shop/pkg/config"
	"shop/pkg/server"
)

var (
	appConfig = flag.String("config", "config.yaml", "application config path")
)

func main() {
	//解析命令行参数
	flag.Parse()

	//创建日志实例
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	//获取系统配置
	conf, err := config.Parse(*appConfig)
	if err != nil {
		logger.Fatalf("[INIT CONFIG]:" + err.Error())
	}

	//创建服务
	s, err := server.New(conf, logger)
	if err != nil {
		logger.Fatalf("[INIT SERVER]:" + err.Error())
	}

	//开启服务
	if err := s.Run(); err != nil {
		logger.Fatalf("[RUN SERVER ERROR]:%v", err)
	}
}
