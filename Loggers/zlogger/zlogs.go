package zlogger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// 日志是否静默状态
var silent = true

func init() {

	hook := lumberjack.Logger{
		//Filename:   path.Join("logs", time.Now().Format("2006-01-02")+".log"), // 日志文件路径
		Filename:   path.Join("logs", "app.log"), // 日志文件路径
		MaxSize:    100,                          // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 50,                           // 日志文件最多保存多少个备份
		MaxAge:     7,                            // 文件最多保存多少天
		Compress:   false,                        // 是否压缩
		LocalTime:  true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "zlogger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields()
	//filed := zap.Fields(zap.String("serverName", "Server1"))

	// 构造日志
	//log = zap.New(core, caller, development, filed)
	log = zap.New(core, caller, development, filed, zap.AddCallerSkip(1))

	//fmt.Println("log 初始化成功")

}

func Debug(f interface{}, v ...interface{}) {
	if silent == true {
		return
	}
	log.Debug(getFormatMsg(f, v...))
}

func Error(f interface{}, v ...interface{}) {
	if silent == true {
		return
	}
	log.Error(getFormatMsg(f, v...))
}

func Panic(f interface{}, v ...interface{}) {
	log.Panic(getFormatMsg(f, v...))
}

func Warn(f interface{}, v ...interface{}) {
	if silent == true {
		return
	}
	log.Warn(getFormatMsg(f, v...))
}

func Info(f interface{}, v ...interface{}) {
	if silent == true {
		return
	}
	log.Info(getFormatMsg(f, v...))
}

func Debugs(args ...interface{}) {
	if silent == true {
		return
	}
	log.Debug(fmt.Sprint(args...))
	//log.Sugar().Debug(args...)
}

func getFormatMsg(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") {
			return fmt.Sprintf(msg, v...)
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
	}
	msg += strings.Repeat(" %v", len(v))
	return fmt.Sprintf(msg, v...)
}

func DebugZ(msg string, fields ...zap.Field) {
	if silent == true {
		return
	}
	log.Debug(msg, fields...)
}

func Errors(msg string, e error) {
	if silent == true {
		return
	}
	log.Error(msg, zap.Error(e))
}

func ErrorZ(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func InfoZ(msg string, fields ...zap.Field) {
	if silent == true {
		return
	}
	log.Info(msg, fields...)
}

func Sync() error {
	// flushes buffer
	return log.Sync()
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func ChangeSilent(status bool) {
	//改变日志的静默状态
	silent = status
}
