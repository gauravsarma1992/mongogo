package mongogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	User struct {
		Name string
		Age  int

		Emails []string
	}
)

func GetDefaultDb() (db *DB) {
	db, _ = NewDB(nil)
	return
}

func TestCreateInit(t *testing.T) {
	db := GetDefaultDb()
	user := User{
		Name:   "Gary",
		Age:    23,
		Emails: []string{"gary@gmail.com", "gary.1@gmail.com"},
	}
	uuid, err := db.Create(user)
	assert.Nil(t, err)
	assert.NotEqual(t, uuid, "")

}

func TestFindOne(t *testing.T) {
	db := GetDefaultDb()
	user := User{
		Name:   "Gary",
		Age:    23,
		Emails: []string{"gary@gmail.com", "gary.1@gmail.com"},
	}
	db.Create(user)
	resUser, err := db.FindOne(user)
	assert.Nil(t, err)
	assert.Equal(t, resUser["name"], "Gary")

}

func TestFindMany(t *testing.T) {
	db := GetDefaultDb()
	user := User{
		Name:   "Gary",
		Age:    23,
		Emails: []string{"gary@gmail.com", "gary.1@gmail.com"},
	}
	db.Create(user)
	userTwo := User{
		Name:   "Bary",
		Age:    23,
		Emails: []string{"gary@gmail.com", "gary.1@gmail.com"},
	}
	db.Create(userTwo)
	resUsers, err := db.FindMany(user)
	assert.Nil(t, err)

	for _, user := range resUsers {
		assert.Equal(t, user["name"], "Gary")
	}

}

func TestUpdateOne(t *testing.T) {
	db := GetDefaultDb()
	user := User{
		Name:   "Gary",
		Age:    23,
		Emails: []string{"gary@gmail.com", "gary.1@gmail.com"},
	}
	db.Create(user)
	updateAttrs := make(map[string]interface{})
	updateAttrs["age"] = 25
	err := db.UpdateOne(updateAttrs, user)
	assert.Nil(t, err)

}
