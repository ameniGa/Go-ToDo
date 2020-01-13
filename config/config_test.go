package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	_,err := GetConfig()
	if err != nil{
		t.Errorf("expected success got %v ",err)
	}

}
