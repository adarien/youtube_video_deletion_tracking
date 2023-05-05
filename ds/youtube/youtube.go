package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

type Client struct {
	Service *youtube.Service
}

func Connect(s string) (Client, error) {

	cs, err := os.ReadFile(s)
	if err != nil {
		return Client{}, fmt.Errorf("client secret read: %w", err)
	}

	config, err := google.ConfigFromJSON(cs, youtube.YoutubeReadonlyScope)
	if err != nil {
		return Client{}, fmt.Errorf("client secret parse: %w", err)
	}

	token, err := getToken(config)
	if err != nil {
		return Client{}, fmt.Errorf("token get: %w", err)
	}

	service, err := youtube.NewService(
		context.Background(),
		option.WithTokenSource(
			config.TokenSource(
				context.Background(),
				token,
			),
		),
	)
	if err != nil {
		return Client{}, fmt.Errorf("service create: %w", err)
	}

	return Client{
		Service: service,
	}, nil
}

// getToken uses a Context and Config to retrieve a Token.
// It returns the retrieved Token and any error encountered.
func getToken(config *oauth2.Config) (*oauth2.Token, error) {

	cacheFile, err := getPathTokenCacheFile()
	if err != nil {
		err := fmt.Errorf("credential path get: %w", err)
		return nil, err
	}

	token, err := getTokenFromFile(cacheFile)
	if err != nil {

		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}

		if err = saveToken(cacheFile, token); err != nil {
			return nil, err
		}
	}

	return token, nil
}

// getPathTokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename and any error encountered.
func getPathTokenCacheFile() (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")

	err = os.MkdirAll(tokenCacheDir, 0700)
	if err != nil {
		return "", err
	}

	return filepath.Join(
		tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json"),
	), err
}

// getTokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func getTokenFromFile(file string) (t *oauth2.Token, err error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		if e := f.Close(); e != nil {
			if err != nil {
				err = fmt.Errorf("close file: %w <- %w", e, err)
			} else {
				err = fmt.Errorf("close file: %w", e)
			}
		}
	}()

	t = &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	if err != nil {
		err = fmt.Errorf("decode json: %w", err)
		return nil, err
	}

	return t, err
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token and any error encountered.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {

	var code string

	fmt.Printf("%s: \n%v\n",
		"Go to the following link in your browser then type the authorization code",
		config.AuthCodeURL("state-token", oauth2.AccessTypeOffline),
	)

	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("authorization code read: %w", err)
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("retrieve token from web: %w", err)
	}

	return token, nil
}

// saveToken uses a file path to create a file and store the token in it.
// It returns any error encountered.
func saveToken(file string, token *oauth2.Token) (err error) {

	fmt.Printf("Saving credential file to: %s\n", file)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("oauth token open: %w", err)
	}
	defer func() {
		if e := f.Close(); e != nil {
			if err != nil {
				err = fmt.Errorf("close file: %w <- %w", e, err)
			} else {
				err = fmt.Errorf("close file: %w", e)
			}
		}
	}()

	if err = json.NewEncoder(f).Encode(token); err != nil {
		return fmt.Errorf("oauth token encode: %w", err)
	}

	return err
}
