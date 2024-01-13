package repositories

import (
	"testing"
)

func TestMessage(t *testing.T) {
	tx := DB.Begin()
	var mr MessageRepository = &MessageRepositoryImpl{DB: tx}
	var ur UserRepository = &UserRepositoryImpl{DB: tx}

	t.Run("Create", func(t *testing.T) {
		user, err := ur.Create("username")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		_, err = mr.Create("title", user.ID)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("List", func(t *testing.T) {
		_, err := mr.List()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("GetById", func(t *testing.T) {
		_, err := mr.GetById(1)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	tx.Rollback()
}
