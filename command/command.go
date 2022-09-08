package command

import (
	"github.com/miyamoto-jo/go-batch-sample/batch"
	"github.com/miyamoto-jo/go-batch-sample/cli"
	"github.com/miyamoto-jo/go-batch-sample/config"
	"github.com/miyamoto-jo/go-batch-sample/logger"
)

var batchList = map[string]batch.Batch{
	"userPaymentLog": batch.NewUserPaymentLog(),
}

// コマンドができること
type Command interface {
	Run() error
	SetEnv()
	getBatch() batch.Batch
}

type command struct {
	BatchName  string   // バッチコマンド名
	ExeEnvName string   // 実行環境名
	SaveTarget string   // 保存先
	Params     []string // バッチのパラメータ
}

func Create() *command {
	b, e, t, p := cli.GetCommandLineBatchInput()
	return &command{
		BatchName:  b,
		ExeEnvName: e,
		SaveTarget: t,
		Params:     p,
	}
}

// コマンド実行
func (c *command) Run() error {
	// 開始ログ
	logger.Infoln(`start ` + c.BatchName)
	// 保存先の表示
	logger.Infoln(`SaveTarget ` + c.SaveTarget)

	// バッチを取得
	batch := c.getBatch()

	err := batch.Execute(c.Params)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	//終了ログ
	logger.Infoln(`finish ` + c.BatchName)

	return nil
}

// 環境情報をセットする
func (c *command) SetEnv() {
	// config設定
	config.SetEnv(c.ExeEnvName)

	// 環境依存の設定を行う
	// graphql, mysql, slack通知, などなど
}

// バッチを取得する
func (c *command) getBatch() batch.Batch {
	for k, f := range batchList {
		if k == c.BatchName {
			return f
		}
	}
	logger.Errorln(`invalid command`)
	// コマンドが間違っていた場合はhelpを表示する
	return batch.NewHelp()
}
