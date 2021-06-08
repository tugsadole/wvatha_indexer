package config

import (
	"testing"
)

func TestGetAppConfig(t *testing.T) {
	_, err := GetAppConfig()
	if err != nil {
		t.Error(err.Error())
	}
}
