package cli

import (
	"flag"
	"fmt"
)

type flags []string

// バッチのパラメータ情報
var batchParams flags

func (s *flags) String() string {
	return fmt.Sprintf("%v", batchParams)
}
func (s *flags) Set(v string) error {
	*s = append(*s, v)
	return nil
}

// バッチ実行コマンドの情報を取得する
func GetCommandLineBatchInput() (string, string, string, []string) {
	batchName := flag.String("batch", "", "exec batch name")
	env := flag.String("env", "dev", "exec environment name")
	saveTarget := flag.String("target", "gouf", "save target name")
	// 複数指定できるflag。バッチのパラメータに使用する。
	flag.Var(&batchParams, "p", "Enter the required param for the batch")
	flag.Parse()

	return *batchName, *env, *saveTarget, batchParams
}
