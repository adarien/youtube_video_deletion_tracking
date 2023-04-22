package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	appctx "github.com/nixys/nxs-go-appctx/v2"
	"github.com/sirupsen/logrus"

	"github.com/adarien/youtube_video_deletion_tracking/ctx"
	"github.com/adarien/youtube_video_deletion_tracking/routines/bot"
)

func main() {

	// Read command line arguments
	args := ctx.ArgsRead()

	appCtx, err := appctx.ContextInit(appctx.Settings{
		CustomContext:    &ctx.Ctx{},
		Args:             &args,
		CfgPath:          args.ConfigPath,
		TermSignals:      []os.Signal{syscall.SIGTERM, syscall.SIGINT},
		ReloadSignals:    []os.Signal{syscall.SIGHUP},
		LogrotateSignals: []os.Signal{syscall.SIGUSR1},
		LogFormatter:     &logrus.JSONFormatter{},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appCtx.Log().Info("program started")

	// main() body function
	defer appCtx.MainBodyGeneric()

	// Create main context
	c := context.Background()

	appCtx.RoutineCreate(c, bot.Runtime)
}
