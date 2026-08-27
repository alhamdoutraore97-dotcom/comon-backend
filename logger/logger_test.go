package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init()
	log := Get()
	assert.NotNil(t, log)
}

func TestGet(t *testing.T) {
	Init()
	log1 := Get()
	log2 := Get()
	assert.NotNil(t, log1)
	assert.NotNil(t, log2)
}
