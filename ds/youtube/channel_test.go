package youtube

import (
	"testing"
)

func TestChannelIDsGet(t *testing.T) {

	c, err := Connect("../../.env/client_secret.json")
	if err != nil {
		t.Errorf("Connect() failed, expected nil error but got %v", err)
	}

	// Positive test case
	username := "test_username"
	// expectedIDs := []string{"", "id2", "id3"}
	// actualIDs, err := c.ChannelIDsGet(username)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }
	// if !reflect.DeepEqual(actualIDs, expectedIDs) {
	// 	t.Errorf("unexpected channel IDs: got %v, want %v", actualIDs, expectedIDs)
	// }

	// Negative test case: channels not found
	actualIDs, err := c.ChannelIDsGet(username)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if actualIDs != nil {
		t.Errorf("unexpected channel IDs: got %v, want nil", actualIDs)
	}
}
