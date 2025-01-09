package main

import (
	"Loggers/zlogger"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println(aurora.BrightGreen("Loggers_Demo"))
	fmt.Println(aurora.BrightRed("Loggers_Demo"))
	fmt.Println(aurora.BrightBlue("Loggers_Demo"))
	fmt.Println(aurora.BrightCyan("Loggers_Demo"))
	fmt.Println(aurora.BrightYellow("Loggers_Demo"))
	fmt.Println(aurora.BrightMagenta("Loggers_Demo"))

	logger_demo1()

	zap_logger_demo1()

	zap_logger_demo2()

	zap_logger_demo3()

	zap_logger_demo4()

	zap_logger_demo5()
}

func logger_demo1() {
	var acolor = aurora.NewAurora(true)

	fmt.Println(aurora.BrightGreen("hello"))

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

func zap_logger_demo1() {

	fmt.Println("zap_logger_demo1")
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

	fmt.Println("----------------")

	defer zlogger.Sync()
}

func zap_logger_demo2() {
	fmt.Println("zap_logger_demo2")
	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-file outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to Kafka and writing
	// them to the console. We'd like to encode the console output and the Kafka
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	topicDebugging := zapcore.AddSync(io.Discard)
	topicErrors := zapcore.AddSync(io.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()

	logger.Debug("constructed a logger")
	logger.Info("hi boys!")
	logger.Error("err:", zap.Error(errors.New("something wrong")))

	fmt.Println("----------------")
}

func zap_logger_demo3() {
	fmt.Println("zap_logger_demo3")
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}

	type foo struct {
		One string
		Two string
	}

	logger.Debug("zap_logger_demo3", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("zap_logger_demo3", zap.Any("foo", foo{One: "one", Two: "two"}))
	//logger.Warn("zap_logger_demo3", zap.Any("foo", foo{One: "one", Two: "two"}))
	//logger.Error("zap_logger_demo3", zap.Any("foo", foo{One: "one", Two: "two"}))
	fmt.Println("----------------")
}

func zap_logger_demo4() {
	fmt.Println("zap_logger_demo4")

	type foo struct {
		One string
		Two string
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	logger.Debug("zap_logger_demo4", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("zap_logger_demo4", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Warn("zap_logger_demo4", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Error("zap_logger_demo4", zap.Any("foo", foo{One: "one", Two: "two"}))
	fmt.Println("----------------")
}

func zap_logger_demo5() {
	fmt.Println("zap_logger_demo5")

	type foo struct {
		One string
		Two string
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logger, _ := config.Build()
	logger.Debug("zap_logger_demo5", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Info("zap_logger_demo5", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Warn("zap_logger_demo5", zap.Any("foo", foo{One: "one", Two: "two"}))
	logger.Error("zap_logger_demo5", zap.Any("foo", foo{One: "one", Two: "two"}))
	fmt.Println("----------------")
}
