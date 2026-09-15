package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUser_JSONMarshal teste la sérialisation JSON
// TestUser_JSONMarshal тестирует JSON-сериализацию
func TestUser_JSONMarshal(t *testing.T) {
	user := User{ID: 1, Name: "Alice", Email: "alice@test.com"}
	data, err := json.Marshal(user)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"name":"Alice"`)
	assert.Contains(t, string(data), `"email":"alice@test.com"`)
}

// TestUser_JSONUnmarshal teste la désérialisation JSON
// TestUser_JSONUnmarshal тестирует JSON-десериализацию
func TestUser_JSONUnmarshal(t *testing.T) {
	data := []byte(`{"id":2,"name":"Bob","email":"bob@test.com"}`)
	var user User
	err := json.Unmarshal(data, &user)
	assert.NoError(t, err)
	assert.Equal(t, 2, user.ID)
	assert.Equal(t, "Bob", user.Name)
	assert.Equal(t, "bob@test.com", user.Email)
}

// TestUser_ZeroValue teste la valeur zéro
// TestUser_ZeroValue тестирует нулевое значение
func TestUser_ZeroValue(t *testing.T) {
	var u User
	assert.Equal(t, 0, u.ID)
	assert.Equal(t, "", u.Name)
	assert.Equal(t, "", u.Email)
}
