package primedb

import "github.com/adarien/youtube_video_deletion_tracking/misc"

const UsersTableName = "users"

type User struct {
	TgID int64  `gorm:"column:tg_user_id"`
	Lang string `gorm:"column:lang"`
}

type UserUpdateData struct {
	TgID int64   `gorm:"column:tg_user_id"`
	Lang *string `gorm:"column:lang"`
}

func (User) TableName() string {
	return UsersTableName
}

func (UserUpdateData) TableName() string {
	return UsersTableName
}

func (db *DB) UserGet(tgID int64) (User, error) {

	user := User{}

	r := db.client.
		Where(
			User{
				TgID: tgID,
			},
		).
		Find(&user)
	if r.Error != nil {
		return User{}, r.Error
	}

	if r.RowsAffected == 0 {
		return User{}, misc.ErrNotFound
	}

	return user, nil
}

func (db *DB) UserUpdate(u UserUpdateData) (User, error) {

	user := User{}

	r := db.client.
		Where(
			UserUpdateData{
				TgID: u.TgID,
			},
		).
		Assign(u).
		FirstOrCreate(&user)
	if r.Error != nil {
		return User{}, r.Error
	}

	return user, nil
}
