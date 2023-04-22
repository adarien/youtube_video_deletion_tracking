package tgbot

import (
	"github.com/adarien/youtube_video_deletion_tracking/misc"
	"github.com/adarien/youtube_video_deletion_tracking/modules/localization"
	"github.com/adarien/youtube_video_deletion_tracking/modules/users"
	tg "github.com/nixys/nxs-go-telegram"
)

type lang struct {
	Tag string
}

type members []string

func helloState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgHello.String(), nil)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: false,
	}, nil
}

// byeState represents StateHandler for `bye` state
func byeState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgBye.String(), nil)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: true,
		NextState:    tg.SessStateDestroy(),
	}, nil
}

// sessionConflictState represents StateHandler for session conflict state
func sessionConflictState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(localization.MsgSessionConflict.String(), nil)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: true,
		NextState:    tg.SessStateDestroy(),
	}, nil
}

func userNotRegisteredState(t *tg.Telegram, sess *tg.Session) (tg.StateHandlerRes, error) {

	_, l, err := userEnv(t, sess)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	m, err := l.MessageCreate(
		localization.MsgUserNotRegistered.String(),
		map[string]string{
			"UserName": sess.UserNameGet(),
		},
	)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: false,
		NextState:    tg.SessStateDestroy(),
	}, nil
}

func (ms members) userAuthCheck(user string) bool {

	for _, u := range ms {
		if u == user {
			return true
		}
	}

	return false
}

func userEnv(t *tg.Telegram, sess *tg.Session) (users.User, localization.Lang, error) {

	var user users.User

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return users.User{}, localization.Lang{}, misc.ErrUserCtxExtract
	}

	// Initialize default language
	defLang, err := bCtx.lb.LangSwitch("")
	if err != nil {
		return users.User{}, localization.Lang{}, err
	}

	b, err = sess.SlotGet(usersSlotName, &user)
	if err != nil {

		if err != tg.ErrSessionNotExist {
			return users.User{}, defLang, err
		}

		// If session was not started

		user, l, err := userEnvInit(bCtx, sess.UserIDGet())
		if err != nil {
			return users.User{}, defLang, err
		}

		return user, l, nil
	}

	if b == true {
		l, err := bCtx.lb.LangSwitch(user.Lang)
		if err != nil {
			return users.User{}, defLang, err
		}
		return user, l, nil
	}

	// If session was started, but slot was not created

	user, l, err := userEnvInit(bCtx, sess.UserIDGet())
	if err != nil {
		return users.User{}, defLang, err
	}

	if err = sess.SlotSave(usersSlotName, user); err != nil {
		return users.User{}, defLang, err
	}

	return user, l, nil
}

func userEnvInit(bCtx botCtx, userID int64) (users.User, localization.Lang, error) {

	// Create user env slot if not exist
	user, err := bCtx.users.Get(userID)
	if err != nil {
		return users.User{}, localization.Lang{}, err
	}

	l, err := bCtx.lb.LangSwitch(user.Lang)
	if err != nil {
		return users.User{}, localization.Lang{}, err
	}

	return user, l, nil
}
