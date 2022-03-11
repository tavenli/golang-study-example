package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"runtime"
)

func Service_Main() {

	fmt.Println("---------------------------------------")
	fmt.Println("runtime.GOOS", runtime.GOOS)
	fmt.Println("service.Interactive()", service.Interactive())
	fmt.Println("service.Platform()", service.Platform())

	fmt.Println("---------------------------------------")

	//Service 接口的主要方法 Run、Start、Stop、Restart、Install、Uninstall

	//配置服务的显示信息
	svcConfig := &service.Config{
		Name:        "Go Service Example Simple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	//执行程序的路径，如果不设置，则为当前程序
	//svcConfig.Executable = "/opt/GoBase.sh"
	svcConfig.Arguments = []string{"-action", "run"}

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

	switch *action {
	case "run":
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	default:
		//要用管理员权限运行命令，才能成功注册为服务
		//支持的参数有 "start", "stop", "restart", "install", "uninstall"
		err = service.Control(s, *action)
		if err != nil {
			log.Fatal(err)
			logger.Error(err)
		}
	}

}

type program struct{}

func (p *program) Start(s service.Service) error {
	//启动服务时触发，建议不要有阻塞代码，执行真实的启动逻辑时要用异步执行
	fmt.Println("execute Start()")
	go p.run()
	return nil
}
func (p *program) run() {
	//启动服务执行的逻辑，需要是阻塞的
	fmt.Println("execute run()")
	for {
	}
}

func (p *program) Stop(s service.Service) error {
	//停止服务时触发，不要有阻塞代码
	fmt.Println("execute Stop()")
	return nil
}
