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

	//ANSI转义码 原生代码打印颜色

	//步骤1：改变当前打印颜色，数字92是颜色代码，还有其它常用颜色代码 30~37、90~97 等
	fmt.Print("\x1b[92m")
	//或者
	//fmt.Sprintf("\x1b[%dm", "")

	//步骤2：打印
	fmt.Println("color changed 1.")
	fmt.Println("color changed 2.")
	fmt.Println("color changed 3.")

	//步骤3：重置颜色为默认
	fmt.Print("\x1b[0m")
	fmt.Println("color reset.")

	//语句话 完成颜色改变和打印并重置
	fmt.Print("\x1b[93m", "color changed", "\x1b[0m", "\n")
	fmt.Print("\x1b[95m", "color changed", "\x1b[0m", "\n")

	fmt.Println(aurora.BrightGreen("Loggers_Demo"))
	fmt.Println(aurora.BrightRed("Loggers_Demo"))
	fmt.Println(aurora.BrightBlue("Loggers_Demo"))
	fmt.Println(aurora.BrightCyan("Loggers_Demo"))
	fmt.Println(aurora.BrightYellow("Loggers_Demo"))
	fmt.Println(aurora.BrightMagenta("Loggers_Demo"))

	fmt.Println(aurora.White("Loggers_Demo"))
	fmt.Println(aurora.Gray(0, "Loggers_Demo"))
	fmt.Println(aurora.Gray(10, "Loggers_Demo"))
	fmt.Println(aurora.Gray(24, "Loggers_Demo"))
	fmt.Println(aurora.BrightWhite("Loggers_Demo"))

	logger_demo1()

	zap_logger_demo1()

	zap_logger_demo2()

	zap_logger_demo3()

	zap_logger_demo4()

	zap_logger_demo5()

	zap_logger_demo6()

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
	//logger := zap.New(core)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	logger.Debug("constructed a logger")
	logger.Info("hi boys!")
	logger.Warn("constructed a logger")
	logger.Error("constructed a logger")
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

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	//atomicLevel := zap.NewAtomicLevel()
	//atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		atomicLevel,                                             // 日志级别
	)

	// 记录 对应源码文件及行号
	caller := zap.AddCaller()

	// 开启开发模式，会记录 DPanic-level
	//development := zap.Development()

	trace := zap.AddStacktrace(zap.ErrorLevel)

	// 设置初始化字段
	filed := zap.Fields()
	//filed := zap.Fields(zap.String("serverName", "Server1"))

	// 构造日志对象
	logger := zap.New(core, caller, trace, filed)

	logger.Debug("zap_logger_demo5 constructed a logger")
	logger.Info("zap_logger_demo5 hi boys!")
	logger.Warn("zap_logger_demo5 constructed a logger")
	logger.Error("zap_logger_demo5 constructed a logger")
	logger.Error("zap_logger_demo5 err:", zap.Error(errors.New("something wrong")))
	fmt.Println("----------------")
}

func zap_logger_demo6() {
	fmt.Println("zap_logger_demo6")

	//输出每行日志的格式，可以看 NewConsoleEncoder 或 NewJSONEncoder 对象的 EncodeEntry 函数

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	colorFomat := func(color uint8, text string) string { return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(color), text) }
	//使用map的目的是减少字符拼接
	colorMap := make(map[zapcore.Level]string, 4)
	colorMap[zapcore.DebugLevel] = colorFomat(92, zapcore.DebugLevel.CapitalString())
	colorMap[zapcore.InfoLevel] = colorFomat(94, zapcore.InfoLevel.CapitalString())
	colorMap[zapcore.WarnLevel] = colorFomat(93, zapcore.WarnLevel.CapitalString())
	colorMap[zapcore.ErrorLevel] = colorFomat(91, zapcore.ErrorLevel.CapitalString())

	customColorLevelEncoder := func(_level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {

		s, ok := colorMap[_level]
		if !ok {
			s = colorFomat(31, _level.CapitalString())
		}
		enc.AppendString(s)

	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	//atomicLevel := zap.NewAtomicLevel()
	//atomicLevel.SetLevel(zap.DebugLevel)

	lowLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		lowLevel,                                                // 日志级别
	)

	// 记录 对应源码文件及行号
	caller := zap.AddCaller()

	// 开启开发模式，会记录 DPanic-level
	//development := zap.Development()

	trace := zap.AddStacktrace(zap.ErrorLevel)

	// 设置初始化字段
	filed := zap.Fields()

	// 构造日志对象
	logger := zap.New(core, caller, trace, filed)

	logger.Debug("zap_logger_demo6 constructed a logger")
	logger.Info("zap_logger_demo6 hi boys!")
	logger.Warn("zap_logger_demo6 constructed a logger")
	logger.Error("zap_logger_demo6 constructed a logger")
	logger.Error("zap_logger_demo6 err:", zap.Error(errors.New("something wrong")))
	fmt.Println("----------------")
}
