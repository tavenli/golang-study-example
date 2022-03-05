package main

import (
	"github.com/kardianos/service"
	"log"
)

func Service_Main() {

	//Service 接口的主要方法 Run、Start、Stop、Restart、Install、Uninstall

	//配置服务的显示信息
	svcConfig := &service.Config{
		Name:        "Go Service Example Simple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
		logger.Error(err)
	}

	err = service.Control(s, "install")
	if err != nil {
		log.Fatal(err)
		logger.Error(err)
	}

}

type program struct{}

func (p *program) Start(s service.Service) error {
	//启动服务时触发，建议不要有阻塞代码，执行真实的启动逻辑时要用异步执行
	go p.run()
	return nil
}
func (p *program) run() {
	//启动服务执行的逻辑
}

func (p *program) Stop(s service.Service) error {
	//停止服务时触发，不要有阻塞代码
	return nil
}
