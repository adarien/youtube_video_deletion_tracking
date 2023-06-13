package youtube

import (
	"testing"
)

func TestPlaylistDataGet(t *testing.T) {

	client, err := Connect("../../.env/client_secret.json")
	if err != nil {
		t.Errorf("Connect() failed, expected nil error but got %v", err)
	}
	// Positive test case
	response, err := client.PlaylistDataGet("UCSPd93is2UQsd_jZ6yHBfqQ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if response == nil {
		t.Errorf("Expected non-nil response")
	}

	// Negative test case
	response, err = client.PlaylistDataGet("")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	if response != nil {
		t.Errorf("Expected nil response")
	}
}
