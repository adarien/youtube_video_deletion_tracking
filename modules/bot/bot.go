package tgbot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adarien/youtube_video_deletion_tracking/ds/primedb"
	"github.com/adarien/youtube_video_deletion_tracking/misc"
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	"github.com/adarien/youtube_video_deletion_tracking/modules/users"

	tg "github.com/nixys/nxs-go-telegram"
	"github.com/sirupsen/logrus"
)

type Settings struct {
	APIToken  string
	Log       *logrus.Logger
	RedisHost string
	Users     []string
	PrimeDB   primedb.DB
	User      string
	Secret    string
}

type Bot struct {
	bot tg.Telegram
}

type botCtx struct {
	log   *logrus.Logger
	m     members
	users users.Users
	lb    localization.Bundle
	d     primedb.DB
	u     string
	s     string
}

func Init(settings Settings) (*Bot, error) {

	lb, err := localization.Init()
	if err != nil {
		return nil, err
	}

	// Set up bot
	bot, err := tg.Init(
		tg.Settings{
			BotSettings: tg.SettingsBot{
				BotAPI: settings.APIToken,
			},
			RedisHost: settings.RedisHost,
		},

		tg.Description{

			Commands: []tg.Command{
				{
					Command:     "settings",
					Description: "Settings",
					Handler:     settingsCmd,
				},
				{
					Command:     "test command",
					Description: "YouTube test",
					Handler:     testCmd,
				},
			},

			InitHandler:  initHandler,
			ErrorHandler: errorHandler,
			PrimeHandler: primeHandler,

			States: map[tg.SessionState]tg.State{

				/*
					Common
				*/
				stateHello: {
					StateHandler: helloState,
				},
				stateBye: {
					StateHandler: byeState,
				},
				stateSessionConflict: {
					StateHandler: sessionConflictState,
				},
				stateUserNotRegistered: {
					StateHandler: userNotRegisteredState,
				},

				/*
					Init settings
				*/
				stateInitLang: {
					StateHandler:    initLangState,
					CallbackHandler: initLangCallback,
				},
				stateInitEnd: {
					StateHandler: initEndState,
				},

				/*
					Settings
				*/
				stateSettings: {
					StateHandler:    settingsState,
					CallbackHandler: settingsCallback,
				},
				// Language
				stateSettingsLangSelect: {
					StateHandler:    settingsLangSelectState,
					CallbackHandler: settingsLangSelectCallback,
				},
				// YouTube test connection
				stateYouTube: {
					StateHandler: youTubeState,
				},
			},
		},
		botCtx{
			log: settings.Log,
			m:   settings.Users,
			users: users.Init(
				users.Settings{
					DB: settings.PrimeDB,
				}),
			lb: lb,
			d:  settings.PrimeDB,
			u:  settings.User,
			s:  settings.Secret,
		})
	if err != nil {
		return nil, fmt.Errorf("bot setup error: %v", err)
	}

	return &Bot{
		bot: bot,
	}, nil
}

// UpdatesGet runtimeBotUpdates checks updates at Telegram and put it into queue
func (b *Bot) UpdatesGet(ctx context.Context, ch chan error) {
	if err := b.bot.GetUpdates(ctx); err != nil {
		if err == tg.ErrUpdatesChanClosed {
			ch <- nil
		} else {
			ch <- err
		}
	} else {
		ch <- nil
	}
}

// Queue runtimeBotQueue processes an updates from queue
func (b *Bot) Queue(ctx context.Context, ch chan error) {
	timer := time.NewTimer(time.Millisecond * 200)
	for {
		select {
		case <-timer.C:
			if err := b.bot.Processing(); err != nil {
				ch <- err
			}
			timer.Reset(time.Millisecond * 200)
		case <-ctx.Done():
			return
		}
	}
}

func (b *Bot) SendMessage(chatID int64, message string) error {
	_, err := b.bot.SendMessage(chatID, 0, tg.SendMessageData{
		Message: message,
	})

	return err
}

// initHandler is a first handler in bots life cycle
// args not used in func, but required for interface compatibility in handler
func initHandler(_ *tg.Telegram, _ *tg.Session) (tg.InitHandlerRes, error) {
	return tg.InitHandlerRes{
		NextState: stateHello,
	}, nil
}

func errorHandler(t *tg.Telegram, s *tg.Session, e error) (tg.ErrorHandlerRes, error) {

	ss, _, _ := s.StateGet()

	_, err := t.SendMessage(s.UserIDGet(), 0, tg.SendMessageData{
		Message: "bot error: `" + e.Error() + "` (state: `" + ss.String() + "`)",
	})
	if err != nil {
		return tg.ErrorHandlerRes{}, err
	}

	return tg.ErrorHandlerRes{}, nil
}

// primeHandler is a handler that checks user's auth
// arg 'tg.HandlerSource' not used in func, but required for interface compatibility in handler
func primeHandler(t *tg.Telegram, sess *tg.Session, _ tg.HandlerSource) (tg.PrimeHandlerRes, error) {

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return tg.PrimeHandlerRes{}, misc.ErrUserCtxExtract
	}

	if bCtx.m.userAuthCheck(sess.UserNameGet()) == false {
		return tg.PrimeHandlerRes{
			NextState: stateUserNotRegistered,
		}, nil
	}

	// Check account for current user exist
	if _, err := bCtx.users.Get(sess.UserIDGet()); err != nil {

		if errors.Is(err, misc.ErrNotFound) == true {

			s, b, err := sess.StateGet()
			if err != nil {
				return tg.PrimeHandlerRes{}, err
			}
			if b == true {
				switch s {
				case
					// If user account not created in DB
					// allowed only init states to interact with bot
					stateInitLang,
					stateInitEnd:
					return tg.PrimeHandlerRes{
						NextState: tg.SessStateContinue(),
					}, nil
				}
			}

			// For all other states (including no state) bot will be switched to
			// init state
			return tg.PrimeHandlerRes{
				NextState: stateInitLang,
			}, nil
		}

		return tg.PrimeHandlerRes{}, err
	}

	return tg.PrimeHandlerRes{
		NextState: tg.SessStateContinue(),
	}, nil
}
