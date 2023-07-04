package users

/*
	WIP...
*/

// // MockDatabase is a mock implementation of the database interface
// type MockDatabase struct {
// 	UserGetFn    func(tgID int64) (primedb.User, error)
// 	UserUpdateFn func(data primedb.UserUpdateData) (primedb.User, error)
// }
//
// // UserGet is a mock implementation of the UserGet method
// func (mdb *MockDatabase) UserGet(tgID int64) (primedb.User, error) {
// 	if mdb.UserGetFn != nil {
// 		return mdb.UserGetFn(tgID)
// 	}
//
// 	return primedb.User{}, nil
// }
//
// // UserUpdate is a mock implementation of the UserUpdate method
// func (mdb *MockDatabase) UserUpdate(data primedb.UserUpdateData) (primedb.User, error) {
// 	if mdb.UserUpdateFn != nil {
// 		return mdb.UserUpdateFn(data)
// 	}
//
// 	return primedb.User{}, nil
// }
//
// func TestUsers_Get(t *testing.T) {
// 	// Positive test case
// 	db := MockDatabase{}
// 	users := Users{d: db}
// 	expectedUser := User{TgID: 123, Lang: "en"}
// 	db.UserGetFn = func(tgID int64) (primedb.User, error) {
// 		return primedb.User{TgID: tgID, Lang: "en"}, nil
// 	}
// 	user, err := users.Get(123)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if user != expectedUser {
// 		t.Errorf("Expected user: %v, got: %v", expectedUser, user)
// 	}
// 	// Negative test case
// 	db.UserGetFn = func(tgID int64) (primedb.User, error) {
// 		return primedb.User{}, fmt.Errorf("error getting user")
// 	}
// 	_, err = users.Get(123)
// 	if err == nil {
// 		t.Error("Expected error, got nil")
// 	}
// }
//
// func TestUsers_UserCreate(t *testing.T) {
// 	// Positive test case
// 	db := MockDatabase{}
// 	users := Users{d: db}
// 	expectedUser := User{TgID: 123, Lang: "en"}
// 	db.UserUpdateFn = func(data primedb.UserUpdateData) (primedb.User, error) {
// 		return primedb.User{TgID: data.TgID, Lang: *data.Lang}, nil
// 	}
// 	user, err := users.UserCreate(123, "en")
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if user != expectedUser {
// 		t.Errorf("Expected user: %v, got: %v", expectedUser, user)
// 	}
// 	// Negative test case
// 	db.UserUpdateFn = func(data primedb.UserUpdateData) (primedb.User, error) {
// 		return primedb.User{}, fmt.Errorf("error updating user")
// 	}
// 	_, err = users.UserCreate(123, "en")
// 	if err == nil {
// 		t.Error("Expected error, got nil")
// 	}
// }
