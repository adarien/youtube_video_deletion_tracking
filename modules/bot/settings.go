package tgbot

import (
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	tg "github.com/nixys/nxs-go-telegram"
)

// settingsCmd represents CommandHandler for `/settings` command
// args not used in func, but required for interface compatibility in handler
func settingsCmd(_ *tg.Telegram, _ *tg.Session, _ string, _ string) (tg.CommandHandlerRes, error) {
	return tg.CommandHandlerRes{
		NextState: stateSettings,
	}, nil
}

func settingsState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgSettings.String(), nil)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	buttons := [][]tg.Button{
		{
			{
				Text:       l.BotButton(localization.ButtonSettingsLang),
				Identifier: buttonIDSettingsLang,
			},
		},
		{
			{
				Text:       l.BotButton(localization.ButtonBack),
				Identifier: buttonIDBack,
			},
		},
	}

	return tg.StateHandlerRes{
		Message:      m,
		Buttons:      buttons,
		StickMessage: true,
	}, nil
}

// settingsCallback represents CallbackHandler for settings state
// args not used in func, but required for interface compatibility in handler
func settingsCallback(_ *tg.Telegram, _ *tg.Session, identifier string) (tg.CallbackHandlerRes, error) {

	switch identifier {
	case buttonIDSettingsLang:
		return tg.CallbackHandlerRes{
			NextState: stateSettingsLangSelect,
		}, nil
	case buttonIDBack:
		return tg.CallbackHandlerRes{
			NextState: stateBye,
		}, nil
	}

	return tg.CallbackHandlerRes{}, nil
}
