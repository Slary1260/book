/*
 * @Author: tj
 * @Date: 2022-09-19 16:25:38
 * @LastEditors: tj
 * @LastEditTime: 2022-09-19 16:25:52
 * @FilePath: \BlackJack\logger\logrus.go
 */
package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// DefaultLogrusLogger DefaultLogrusLogger
func DefaultLogrusLogger() {
	NewLogrusLogger("", DefaultLogFilePath())
}

// NewLogrusLogger NewLogrusLogger
func NewLogrusLogger(dir, fileName string) {
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   filepath.Join(dir, fileName),
		MaxSize:    20, // Max megabytes before log is rotated
		MaxBackups: 10, // Max number of old log files to keep
		MaxAge:     30, // Max number of days to retain log files
		Compress:   true,
	}

	// logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.RFC3339})
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
	logrus.SetReportCaller(true)

	// 编译为dll文件 不能使用os.Stdout
	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
	logrus.SetOutput(logMultiWriter)
}

// DefaultLogFilePath DefaultLogFilePath
func DefaultLogFilePath() string {
	var logFilePath string
	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))

	switch runtime.GOOS {
	case "android":
		// TODO android logFilePath
		// logFilePath = "/storage/emulated/0/Android/data/com.gdh.project/files/project.log"

	case "windows":
		// TODO windows logFilePath
		// logFilePath = filepath.Join(os.Getenv("AppData"), "project/log", "project.log")
		logFilePath = filepath.Join("./log", fileName)

	case "darwin":
		// logFilePath = "~/Library/Application Support/project/project.log"
		logFilePath = filepath.Join("./log", fileName)

	default:
		// xdgCfg := os.Getenv("XDG_CONFIG_HOME")
		// if xdgCfg != "" {
		// 	logFilePath = filepath.Join(xdgCfg, "project", "project.log")
		// } else {
		// 	logFilePath = filepath.Join("~/.config/project/", "project.log")
		// }
		logFilePath = filepath.Join("./log", fileName)
	}
	return logFilePath
}
