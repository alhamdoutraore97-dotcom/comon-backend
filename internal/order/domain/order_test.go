package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_JSONMarshal(t *testing.T) {
	o := Order{ID: 1, UserID: 2, Product: "Book", Quantity: 3, Status: "pending"}
	data, err := json.Marshal(o)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"product":"Book"`)
	assert.Contains(t, string(data), `"quantity":3`)
}

func TestOrder_JSONUnmarshal(t *testing.T) {
	data := []byte(`{"id":5,"user_id":1,"product":"Pen","quantity":10,"status":"shipped"}`)
	var o Order
	err := json.Unmarshal(data, &o)
	assert.NoError(t, err)
	assert.Equal(t, 5, o.ID)
	assert.Equal(t, "Pen", o.Product)
	assert.Equal(t, 10, o.Quantity)
	assert.Equal(t, "shipped", o.Status)
}

func TestOrder_ZeroValue(t *testing.T) {
	var o Order
	assert.Equal(t, 0, o.ID)
	assert.Equal(t, "", o.Product)
	assert.Equal(t, 0, o.Quantity)
}
