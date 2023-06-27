package primedb

// func TestUserGet(t *testing.T) {
// 	// Positive test case
// 	db := setupDB()
// 	defer db.client.Close()
// 	// Add a user to the database
// 	user := User{TgID: 123}
// 	db.client.Create(&user)
// 	// Get the user from the database
// 	result, err := db.UserGet(123)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if result.TgID != 123 {
// 		t.Errorf("Expected TgID to be 123, but got %v", result.TgID)
// 	}
// 	// Negative test case
// 	result, err = db.UserGet(456)
// 	if err != misc.ErrNotFound {
// 		t.Errorf("Expected error %v, but got %v", misc.ErrNotFound, err)
// 	}
// 	if result != (User{}) {
// 		t.Errorf("Expected empty User, but got %v", result)
// 	}
// }
//
// func TestUserUpdate(t *testing.T) {
// 	// Positive test case
// 	db := setupDB()
// 	defer db.client.Close()
// 	// Add a user to the database
// 	user := User{TgID: 123}
// 	db.client.Create(&user)
// 	// Update the user's name
// 	update := UserUpdateData{TgID: 123, Name: "John"}
// 	result, err := db.UserUpdate(update)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if result.Name != "John" {
// 		t.Errorf("Expected Name to be John, but got %v", result.Name)
// 	}
// 	// Negative test case
// 	update = UserUpdateData{TgID: 456, Name: "Jane"}
// 	result, err = db.UserUpdate(update)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if result.TgID != 456 {
// 		t.Errorf("Expected TgID to be 456, but got %v", result.TgID)
// 	}
// 	if result.Name != "Jane" {
// 		t.Errorf("Expected Name to be Jane, but got %v", result.Name)
// 	}
// }
