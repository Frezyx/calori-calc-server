package store_test

import (
	"testing"

	"github.com/Frezyx/calory-calc-server/internal/app/model"
	"github.com/Frezyx/calory-calc-server/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// func TestUserRepository_FindByEmail(t *testing.T) {
// 	s := teststore.New()
// 	u1 := model.TestUser(t)
// 	_, err := s.User().FindByEmail(u1.Email)
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

// 	s.User().Create(u1)
// 	u2, err := s.User().FindByEmail(u1.Email)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u2)
// }
