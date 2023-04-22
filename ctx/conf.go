package ctx

import (
	conf "github.com/nixys/nxs-go-conf"
)

type confOpts struct {
	LogFile  string `conf:"logFile" conf_extraopts:"default=stdout"`
	LogLevel string `conf:"logLevel" conf_extraopts:"default=info"`
	PidFile  string `conf:"pidFile"`

	PgSQL    pgSQLConf    `conf:"pgSQL" conf_extraopts:"required"`
	Telegram telegramConf `conf:"telegram" conf_extraopts:"required"`
}

type pgSQLConf struct {
	Host     string `conf:"host" conf_extraopts:"required"`
	Port     int    `conf:"port" conf_extraopts:"required"`
	DB       string `conf:"db" conf_extraopts:"required"`
	User     string `conf:"user" conf_extraopts:"required"`
	Password string `conf:"password" conf_extraopts:"required"`
}

type telegramConf struct {
	APIToken  string   `conf:"apiToken" conf_extraopts:"required"`
	RedisHost string   `conf:"redisHost" conf_extraopts:"default=127.0.0.1:6379"`
	Users     []string `conf:"users" conf_extraopts:"required"`
}

func confRead(confPath string) (confOpts, error) {

	var c confOpts

	err := conf.Load(&c, conf.Settings{
		ConfPath:    confPath,
		ConfType:    conf.ConfigTypeYAML,
		UnknownDeny: true,
	})
	if err != nil {
		return c, err
	}

	return c, err
}
