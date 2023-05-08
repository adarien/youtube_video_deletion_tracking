package youtube

import "fmt"

func (c *Client) ChannelIDsGet(username string) ([]string, error) {

	var ids []string

	call := c.Service.Channels.List([]string{"snippet", "contentDetails"})

	response, err := call.ForUsername(username).Do()
	if err != nil {
		return nil, fmt.Errorf("ch not call: %w", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("incorrect username")
	}

	for _, ch := range response.Items {
		ids = append(ids, ch.Id)
	}

	return ids, nil
}
