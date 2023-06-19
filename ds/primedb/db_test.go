package primedb

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Open() (*gorm.DB, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*gorm.DB), args.Error(1)
}

func (m *MockDB) Close() error {
	args := m.Called()
	return args.Error(0)
}

// TODO: WIP...
// func TestConnect(t *testing.T) {
// 	// Arrange
// 	mockedDB := new(MockDB)
// 	db := DB{client: mockedDB}
// 	mockedDB.On("Open").Return(new(gorm.DB), nil)
//
// 	// Act
// 	_, err := db.client.Open()
//
// 	// Assert
// 	assert.NoError(t, err)
// 	mockedDB.AssertExpectations(t)
// }

// func TestConnect_Error(t *testing.T) {
// 	// Arrange
// 	mockedDB := new(MockDB)
// 	db := DB{client: mockedDB}
// 	mockedDB.On("Open").Return(nil, errors.New("error opening db"))
//
// 	// Act
// 	_, err := db.client.Open()
//
// 	// Assert
// 	assert.Error(t, err)
// 	mockedDB.AssertExpectations(t)
// }
//
// func TestClose(t *testing.T) {
// 	// Arrange
// 	mockedDB := new(MockDB)
// 	db := DB{client: mockedDB}
// 	mockedDB.On("Close").Return(nil)
//
// 	// Act
// 	err := db.client.Close()
//
// 	// Assert
// 	assert.NoError(t, err)
// 	mockedDB.AssertExpectations(t)
// }
//
// func TestClose_Error(t *testing.T) {
// 	// Arrange
// 	mockedDB := new(MockDB)
// 	db := DB{client: mockedDB}
// 	mockedDB.On("Close").Return(errors.New("error closing db"))
//
// 	// Act
// 	err := db.client.Close()
//
// 	// Assert
// 	assert.Error(t, err)
// 	mockedDB.AssertExpectations(t)
// }
