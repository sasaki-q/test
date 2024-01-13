package repositories

import (
	"testing"
)

func TestUser(t *testing.T) {
	tx := DB.Begin()
	var r UserRepository = &UserRepositoryImpl{DB: tx}

	t.Run("Create", func(t *testing.T) {
		_, err := r.Create("username")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("List", func(t *testing.T) {
		_, err := r.List()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("GetById", func(t *testing.T) {
		_, err := r.GetById(1)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	tx.Rollback()
}
