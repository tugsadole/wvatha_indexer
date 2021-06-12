package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppConfig(t *testing.T) {
	_, err := GetAppConfig()
	assert.Error(t, err, err.Error())
}
