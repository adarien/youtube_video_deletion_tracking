package tgbot

import (
	"github.com/adarien/youtube_video_deletion_tracking/ds/primedb"
	"github.com/adarien/youtube_video_deletion_tracking/misc"
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	tg "github.com/nixys/nxs-go-telegram"
)

func settingsLangSelectState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgLangSelect.String(), nil)
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
			{
				{
					Text:       l.BotButton(localization.ButtonBack),
					Identifier: buttonIDBack,
				},
			},
		},
		StickMessage: true,
	}, nil
}

func settingsLangSelectCallback(t *tg.Telegram, sess *tg.Session, identifier string) (tg.CallbackHandlerRes, error) {

	slotL := lang{}

	b, err := sess.SlotGet(langSlotName, &slotL)
	if err != nil {
		return tg.CallbackHandlerRes{}, err
	}
	if b == false {
		u, _, err := userEnv(t, sess)
		if err != nil {
			return tg.CallbackHandlerRes{}, err
		}

		slotL.Tag = u.Lang
	}

	switch identifier {
	case buttonIDEN:

		if slotL.Tag == buttonIDEN {
			return tg.CallbackHandlerRes{}, nil
		}

	case buttonIDRU:

		if slotL.Tag == buttonIDRU {
			return tg.CallbackHandlerRes{}, nil
		}

	case buttonIDBack:

		return tg.CallbackHandlerRes{
			NextState: stateSettings,
		}, nil
	}

	slotL.Tag = identifier

	err = sess.SlotSave(langSlotName, slotL)
	if err != nil {
		return tg.CallbackHandlerRes{}, err
	}

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return tg.CallbackHandlerRes{}, misc.ErrUserCtxExtract
	}

	user, err := bCtx.d.UserUpdate(primedb.UserUpdateData{
		TgID: sess.UserIDGet(),
		Lang: &identifier,
	})
	if err != nil {
		return tg.CallbackHandlerRes{}, err
	}

	if err = sess.SlotSave(usersSlotName, user); err != nil {
		return tg.CallbackHandlerRes{}, err
	}

	return tg.CallbackHandlerRes{
		NextState: stateSettingsLangSelect,
	}, nil
}
