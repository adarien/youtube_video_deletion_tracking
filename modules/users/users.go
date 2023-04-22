package users

import (
	"fmt"

	"github.com/adarien/youtube_video_deletion_tracking/ds/primedb"
)

type Settings struct {
	DB primedb.DB
}

type Users struct {
	d primedb.DB
}

type User struct {
	TgID int64
	Lang string
}

type UserUpdateData struct {
	Lang *string
}

func Init(s Settings) Users {
	return Users{
		d: s.DB,
	}
}

func (uu *Users) Get(tgID int64) (User, error) {

	u, err := uu.d.UserGet(tgID)
	if err != nil {
		return User{}, err
	}

	return User{
		TgID: u.TgID,
		Lang: u.Lang,
	}, nil
}

func (uu *Users) UserCreate(tgID int64, lang string) (User, error) {

	if _, err := uu.d.UserUpdate(primedb.UserUpdateData{
		TgID: tgID,
		Lang: &lang,
	}); err != nil {
		return User{}, fmt.Errorf("create new user record: %w", err)
	}

	return User{
		TgID: tgID,
		Lang: lang,
	}, nil
}
