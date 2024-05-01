package main

import (
	"Loggers/zlogger"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func main() {

	//
	//logger_demo1()

	zlogger_demo1()
}

func logger_demo1() {
	var acolor = aurora.NewAurora(true)

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	//	gologger.DefaultLogger.SetFormatter(&formatter.JSON{})

	//logsOptions := writer.DefaultFileWithRotationOptions
	//logsOptions.Rotate = true
	//logsOptions.Compress = true
	//logsOptions.RotateEachDay = true
	//logsOptions.FileName = "app.log"
	//filewriterWithRotation, err := writer.NewFileWithRotation(&logsOptions)
	//if err != nil {
	//	panic(err)
	//}
	//gologger.DefaultLogger.SetWriter(filewriterWithRotation)

	gologger.Print().Msgf("\tgologger: sample test\t\n")
	gologger.Info().Str("user", "pdteam").Msg("running simulation program")
	for i := 0; i < 10; i++ {
		gologger.Info().Str("count", strconv.Itoa(i)).Msg("running simulation step...")
	}
	gologger.Debug().Str("state", "running").Msg("planner running")
	gologger.Warning().Str("state", "errored").Str("status", "404").Msg("could not run")

	//Fatal 会中断程序执行
	//gologger.Fatal().Msg(acolor.BrightGreen("bye bye").String())
	gologger.Info().Msgf("%v", acolor.BrightGreen("latest1"))
	//gologger.Fatal().Msg(fmt.Sprintf("(%v)", acolor.BrightGreen("latest2")))
	gologger.Info().Msgf("Current httpx version %v", fmt.Sprintf("(%v)", acolor.BrightGreen("latest3")))

	fmt.Println("----------------")
}

func zlogger_demo1() {
	action := "do loging"
	zlogger.Info("开始运行", time.Now())
	zlogger.ChangeSilent(false)

	err := errors.New("some error")

	zlogger.Debug("hi boys!")
	zlogger.Debug("some format message %d ,the number you got it?", 100)
	zlogger.DebugZ("some message", zap.String("action", action))
	zlogger.Debugs("some words,", 100, time.Now())
	zlogger.Info("---------")
	zlogger.Error("err:", zap.Error(err))
	zlogger.Errors("Logger error", err)
}
