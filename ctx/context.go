package ctx

import (
	"github.com/adarien/youtube_video_deletion_tracking/ds/primedb"
	tgbot "github.com/adarien/youtube_video_deletion_tracking/modules/bot"
	appctx "github.com/nixys/nxs-go-appctx/v2"
)

// Ctx defines application custom context
type Ctx struct {
	Conf confOpts
	Bot  *tgbot.Bot
}

// Init initiates application custom context
func (c *Ctx) Init(opts appctx.CustomContextFuncOpts) (appctx.CfgData, error) {

	// Read config file
	conf, err := confRead(opts.Config)
	if err != nil {
		return appctx.CfgData{}, err
	}

	// Set application context
	c.Conf = conf

	// Connect to PgSQL
	primeDB, err := primedb.Connect(primedb.Settings{
		Host:     c.Conf.PgSQL.Host,
		Port:     c.Conf.PgSQL.Port,
		Database: c.Conf.PgSQL.DB,
		User:     c.Conf.PgSQL.User,
		Password: c.Conf.PgSQL.Password,
	})
	if err != nil {
		return appctx.CfgData{}, err
	}

	c.Bot, err = tgbot.Init(tgbot.Settings{
		APIToken:  c.Conf.Telegram.APIToken,
		Log:       opts.Log,
		RedisHost: c.Conf.Telegram.RedisHost,
		Users:     c.Conf.Telegram.Users,
		PrimeDB:   primeDB,
	})
	if err != nil {
		return appctx.CfgData{}, err
	}

	return appctx.CfgData{
		LogFile:  c.Conf.LogFile,
		LogLevel: c.Conf.LogLevel,
		PidFile:  c.Conf.PidFile,
	}, nil
}

// Reload reloads application custom context
func (c *Ctx) Reload(opts appctx.CustomContextFuncOpts) (appctx.CfgData, error) {

	opts.Log.Debug("reloading context")

	return c.Init(opts)
}

// Free frees application custom context
func (c *Ctx) Free(opts appctx.CustomContextFuncOpts) int {

	opts.Log.Debug("freeing context")

	return 0
}
