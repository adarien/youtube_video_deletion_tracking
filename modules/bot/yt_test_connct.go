package tgbot

import (
	"fmt"
	"github.com/adarien/youtube_video_deletion_tracking/misc"
	"github.com/adarien/youtube_video_deletion_tracking/modules/youtube"
	tg "github.com/nixys/nxs-go-telegram"
)

func testCmd(_ *tg.Telegram, _ *tg.Session, _ string, _ string) (tg.CommandHandlerRes, error) {
	return tg.CommandHandlerRes{
		NextState: stateYouTube,
	}, nil
}

func youTubeState(t *tg.Telegram, _ *tg.Session) (tg.StateHandlerRes, error) {

	var m string

	bCtx, b := t.UsrCtxGet().(botCtx)
	if b == false {
		return tg.StateHandlerRes{}, misc.ErrUserCtxExtract
	}

	yt, err := youtube.Init(
		youtube.Settings{
			Secret: bCtx.s,
		},
	)
	if err != nil {
		return tg.StateHandlerRes{}, err
	}

	pls, err := yt.PlayListsGet(bCtx.u)
	if err != nil {
		return tg.StateHandlerRes{}, fmt.Errorf("playlists pls get: %w", err)
	}

	for _, l := range pls {
		m = m + fmt.Sprintf("%s\n", l.Title)
	}

	return tg.StateHandlerRes{
		Message:      m,
		StickMessage: true,
	}, nil
}
