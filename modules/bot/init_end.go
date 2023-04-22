package tgbot

import (
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	tg "github.com/nixys/nxs-go-telegram"
)

func initEndState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(
		localization.MsgInitEnd.String(),
		map[string]string{
			"UserName": sess.UserNameGet(),
		},
	)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: true,
		NextState:    tg.SessStateDestroy(),
	}, nil
}
