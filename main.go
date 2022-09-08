package main

import (
	"os"

	"github.com/miyamoto-jo/go-batch-sample/apperror"
	"github.com/miyamoto-jo/go-batch-sample/command"
	"github.com/miyamoto-jo/go-batch-sample/logger"
)

func main() {
	defer checkPanic()
	initialize()

	cmd := command.Create()
	cmd.SetEnv()

	// バッチ実行し、終了コード判定を行う
	os.Exit(apperror.ExitHandler(cmd.Run()))
}

const useLogFile = true // If this param is true, write out log to /log/application.log file instead of writing to stdout.

// 初期処理
func initialize() {
	// 独自logger設定
	logger.SetupLog(useLogFile)
}

// パニックが発生したらエラーログを出して終了する
func checkPanic() {
	if r := recover(); r != nil {
		logger.Errorln(r)
		logger.StackTrace()
		os.Exit(apperror.ExitCodeError)
	}
}
