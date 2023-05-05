// TODO: it is necessary to transfer part of functionality to ds/youtube

package youtube

import (
	"fmt"
	"github.com/adarien/youtube_video_deletion_tracking/ds/youtube"
	yt "google.golang.org/api/youtube/v3"
)

type Settings struct {
	Secret string
	User   string
}

type YT struct {
	s yt.Service
}

type playlistMeta struct {
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
		s: *client.Service,
	}, nil
}

func (y *YT) ChannelListsGet(username string) (*yt.ChannelListResponse, error) {

	call := y.s.Channels.List([]string{"snippet", "contentDetails"})
	call = call.ForUsername(username)

	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("channel not call: %v", err)
	}
	if len(response.Items) == 0 {
		return nil, fmt.Errorf("incorrect userName")
	}

	return response, nil
}

func (y *YT) PlayListsGet(response *yt.ChannelListResponse) ([]playlistMeta, error) {

	var playlists []playlistMeta

	channelID := response.Items[0].Id

	response2, err := y.PlaylistDataGet(channelID)
	if err != nil {
		return nil, err
	}

	for _, playlist := range response2.Items {
		if playlist.Snippet.Title != "Favorites" {
			meta := playlistMeta{}
			meta.ID = playlist.Id
			meta.Title = playlist.Snippet.Title
			meta.Count = playlist.ContentDetails.ItemCount
			playlists = append(playlists, meta)
		}
	}

	return playlists, nil
}

func (y *YT) PlaylistDataGet(channelId string) (*yt.PlaylistListResponse, error) {

	part := []string{"snippet", "contentDetails"}
	call := y.s.Playlists.List(part)
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	call = call.MaxResults(25)

	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("getPlaylistsInfo not call: %v", err)
	}

	return response, nil
}
