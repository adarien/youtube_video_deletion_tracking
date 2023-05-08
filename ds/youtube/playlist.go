// placeholder for YouTube playlist related functions

package youtube

import (
	"fmt"
	yt "google.golang.org/api/youtube/v3"
)

const maxResults = 25

func (c *Client) PlaylistDataGet(channelID string) (*yt.PlaylistListResponse, error) {

	call := c.Service.Playlists.List([]string{"snippet", "contentDetails"})
	if channelID != "" {
		call = call.ChannelId(channelID)
	}

	response, err := call.MaxResults(maxResults).Do()
	if err != nil {
		return nil, fmt.Errorf("getPlaylistsInfo not call: %v", err)
	}

	return response, nil
}
