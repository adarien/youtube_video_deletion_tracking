package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adarien/youtube_video_deletion_tracking/ctx"
	appctx "github.com/nixys/nxs-go-appctx/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

type YT struct {
	service *youtube.Service
}

// TODO: raw version from net, need improvement
// temporary unused

func Init(appCtx *appctx.AppContext) (YT, error) {

	cc := appCtx.CustomCtx().(*ctx.Ctx)

	cs, err := os.ReadFile(cc.Conf.YouTube.Secret)
	if err != nil {
		return YT{}, fmt.Errorf("read client secret: %w", err)
	}

	config, err := google.ConfigFromJSON(cs, youtube.YoutubeReadonlyScope)
	if err != nil {
		return YT{}, fmt.Errorf("parse client secret: %w", err)
	}

	token, err := getToken(config)
	if err != nil {
		return YT{}, fmt.Errorf("get token: %w", err)
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
		return YT{}, fmt.Errorf("service create: %w", err)
	}

	return YT{
		service: service,
	}, nil
}

// getToken uses a Context and Config to retrieve a Token.
// It returns the retrieved Token and any error encountered.
func getToken(config *oauth2.Config) (*oauth2.Token, error) {

	cacheFile, err := getPathTokenCacheFile()
	if err != nil {
		err := fmt.Errorf("unable to get path to cached credential file: %s", err)
		return nil, err
	}

	token, err := getTokenFromFile(cacheFile)
	if err != nil {
		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		err = saveToken(cacheFile, token)
		if err != nil {
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
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// getTokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeError := f.Close()
		if err == nil {
			err = closeError
		}
	}()
	if err != nil {
		err = fmt.Errorf("unable to close file: %s", err)
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token and any error encountered.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	instruction := "Go to the following link in your browser then type the authorization code"
	fmt.Printf("%s: \n%v\n", instruction, authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		err = fmt.Errorf("unable to read authorization code: %s", err)
		return nil, err
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		err = fmt.Errorf("unable to retrieve token from web: %s", err)
		return nil, err
	}

	return token, nil
}

// saveToken uses a file path to create a file and store the token in it.
// It returns any error encountered.
func saveToken(file string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		err = fmt.Errorf("unable to cache oauth token: %s", err)
		return err
	}

	defer func() {
		closeError := f.Close()
		if err == nil {
			err = closeError
		}
	}()
	if err != nil {
		err = fmt.Errorf("unable to close file: %s", err)
		return err
	}

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return err
	}

	return nil
}
