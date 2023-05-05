package tgbot

import tg "github.com/nixys/nxs-go-telegram"

var (
	// Init
	stateInitLang = tg.SessState("initLang")
	stateInitEnd  = tg.SessState("initEnd")

	// Common
	stateBye               = tg.SessState("bye")
	stateHello             = tg.SessState("hello")
	stateSessionConflict   = tg.SessState("sessionConflict")
	stateUserNotRegistered = tg.SessState("userNotRegistered")

	// Settings
	stateSettings           = tg.SessState("settings")
	stateSettingsLangSelect = tg.SessState("settLangSelect")

	// YouTube
	stateYouTube = tg.SessState("youTube")
)
