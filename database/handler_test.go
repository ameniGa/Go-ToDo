package database

import (
	"github.com/3almadmoon/ameni-assignment/config"
	td "github.com/3almadmoon/ameni-assignment/testData"
	"testing"
)

func TestCreateDBhandler(t *testing.T) {
	for _, testCase := range td.TTCreateHandler {
		conf := config.Config{Database: struct {
			Type       string
			Uri        string
			Name       string
			Collection string
		}{
			Type: testCase.Type,
			Uri:  testCase.Uri,
		},
		}

		t.Run(testCase.Name, func(t *testing.T) {
			_, err := CreateDBhandler(&conf)
			if err != nil && !testCase.HasError {
				t.Errorf("expected success , got error: %v", err)
			}
			if err == nil && testCase.HasError {
				t.Error("expected error")
			}
		})
	}
}
