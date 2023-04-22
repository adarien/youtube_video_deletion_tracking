package tgbot

import (
	"github.com/adarien/youtube_video_deletion_tracking/misc"
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	tg "github.com/nixys/nxs-go-telegram"
)

// initLangState represents StateHandler for `initLang` state
// *tg.Session not used in func, but required for interface compatibility in handler
func initLangState(t *tg.Telegram, _ *tg.Session) (tg.StateHandlerRes, error) {

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return tg.StateHandlerRes{}, misc.ErrUserCtxExtract
	}

	l, err := bCtx.lb.LangSwitch("")
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgInitLang.String(), nil)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message: m,
		Buttons: [][]tg.Button{
			{
				{
					Text:       l.BotButton(localization.ButtonEN),
					Identifier: buttonIDEN,
				},
			},
			{
				{
					Text:       l.BotButton(localization.ButtonRU),
					Identifier: buttonIDRU,
				},
			},
		},
		StickMessage: true,
	}, nil
}

func initLangCallback(t *tg.Telegram, sess *tg.Session, identifier string) (tg.CallbackHandlerRes, error) {

	var l string

	switch identifier {
	case buttonIDEN, buttonIDRU:
		l = identifier
	default:
		l = buttonIDEN
	}

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return tg.CallbackHandlerRes{}, misc.ErrUserCtxExtract
	}

	user, err := bCtx.users.UserCreate(sess.UserIDGet(), l)
	if err != nil {
		return tg.CallbackHandlerRes{}, err
	}

	if err = sess.SlotSave(usersSlotName, user); err != nil {
		return tg.CallbackHandlerRes{}, err
	}

	return tg.CallbackHandlerRes{
		NextState: stateInitEnd,
	}, nil
}
