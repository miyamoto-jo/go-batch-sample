package batch

import "fmt"

/*
	ヘルプを表示する処理
	コマンド実行例
	./gameplayer-batch -batch help
*/

type help struct{}

func NewHelp() help {
	return help{}
}

func (h help) Execute(params []string) error {
	fmt.Println(`* help`)
	fmt.Println(`enter the command to execute in [-command]`)
	fmt.Println(`example: go run main.go -batch list -env stg -p param1 -p param2`)
	return nil
}
