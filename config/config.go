package config

var conf *Config

type Config struct {
	Env string
}

// Config情報を返す。InitConf実行後に使用できる想定
func GetConf() *Config {
	if conf == nil {
		panic(`config is not set. please execute "InitConf"`)
	}
	return conf
}

// 環境情報をconfigにセット
func SetEnv(env string) {
	switch env {
	case `local`:
		conf = &Config{
			Env: env,
		}
	case `dev`:
		conf = &Config{
			Env: env,
		}
	case `stg`:
		conf = &Config{
			Env: env,
		}
	case `prd`:
		conf = &Config{
			Env: env,
		}
	default:
		panic(``)
	}
}
