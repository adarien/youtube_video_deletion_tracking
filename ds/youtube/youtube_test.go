package youtube

import (
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {

	// Positive test case
	client, err := Connect("../../.env/client_secret.json")
	if err != nil {
		t.Errorf("Connect() failed, expected nil error but got %v", err)
	}
	if client.Service == nil {
		t.Errorf("Connect() failed, expected non-nil service but got nil")
	}

	// Negative test case
	_, err = Connect("nonexistent_file.json")
	if err == nil {
		t.Errorf("Connect() failed, expected non-nil error but got nil")
	}
}

func TestGetToken(t *testing.T) {

	// positive test case
	config := &oauth2.Config{}
	token, err := getToken(config)
	if err != nil {
		t.Errorf("getToken() returned an error: %v", err)
	}
	if token == nil {
		t.Errorf("getToken() returned nil token")
	}
	// negative test case - invalid cache file path
	invalidPath := "/path/to/invalid/file"
	_, err = getTokenFromFile(invalidPath)
	if err == nil {
		t.Errorf("getTokenFromFile() should return an error for an invalid path")
	}
	// negative test case - invalid token cache file
	invalidFile := "invalid_file"
	_, err = getTokenFromFile(invalidFile)
	if err == nil {
		t.Errorf("getTokenFromFile() should return an error for an invalid file")
	}
	// negative test case - failed to get token from web
	config = &oauth2.Config{
		Endpoint: oauth2.Endpoint{},
	}
	_, err = getTokenFromWeb(config)
	if err == nil {
		t.Errorf("getTokenFromWeb() should return an error for an invalid config")
	}
	// negative test case - failed to save token to cache file
	cacheFile := "/path/to/invalid/file"
	err = saveToken(cacheFile, token)
	if err == nil {
		t.Errorf("saveToken() should return an error for an invalid file path")
	}
}

func TestGetPathTokenCacheFile(t *testing.T) {
	// Positive test case
	expectedPath := filepath.Join(os.Getenv("HOME"), ".credentials", "youtube-go-quickstart.json")
	actualPath, err := getPathTokenCacheFile()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actualPath != expectedPath {
		t.Errorf("Expected path %s, but got %s", expectedPath, actualPath)
	}
	// Negative test case
	_, err = os.Stat("/path/to/invalid/file")
	if !os.IsNotExist(err) {
		t.Errorf("Expected file %s to not exist, but it does", expectedPath)
	}
}

// TestGetTokenFromFile tests the getTokenFromFile function.
// TODO: repair this test
func TestGetTokenFromFile(t *testing.T) {
	// Positive test case
	expectedToken := &oauth2.Token{AccessToken: "token123", TokenType: "Bearer"}
	file := "testdata/token.json"
	actualToken, err := getTokenFromFile(file)
	if err != nil {
		t.Errorf("getTokenFromFile(%v) failed with error %v", file, err)
	}
	if !reflect.DeepEqual(actualToken, expectedToken) {
		t.Errorf("getTokenFromFile(%v) = %v, expected %v", file, actualToken, expectedToken)
	}
	// Negative test case
	file = "testdata/non-existent-file.json"
	_, err = getTokenFromFile(file)
	if err == nil {
		t.Errorf("getTokenFromFile(%v) should have failed but did not", file)
	}
}
