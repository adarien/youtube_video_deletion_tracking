package youtube

import (
	"github.com/adarien/youtube_video_deletion_tracking/ds/youtube"
)

type Settings struct {
	Secret string
	User   string
}

type YT struct {
	c youtube.Client
}

// TODO: WIP...
type playlistData struct {
	ID    string
	Title string
	Count int64
}

func Init(s Settings) (YT, error) {

	client, err := youtube.Connect(s.Secret)
	if err != nil {
		return YT{}, err
	}

	return YT{
		c: client,
	}, nil
}

func (y *YT) PlayListsGet(username string) ([]playlistData, error) {

	var playlists []playlistData

	channelID, err := y.c.ChannelIDsGet(username)
	if err != nil {
		return nil, err
	}

	for _, id := range channelID {

		response2, err := y.c.PlaylistDataGet(id)
		if err != nil {
			return nil, err
		}

		for _, playlist := range response2.Items {

			if playlist.Snippet.Title == "Favorites" {
				continue
			}

			playlists = append(
				playlists,
				playlistData{
					ID:    playlist.Id,
					Title: playlist.Snippet.Title,
					Count: playlist.ContentDetails.ItemCount,
				},
			)
		}
	}

	return playlists, nil
}
