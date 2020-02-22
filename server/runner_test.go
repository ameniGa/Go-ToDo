package server

import (
	"github.com/3almadmoon/ameni-assignment/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ttCreate = []struct {
	name       string
	runnerType string
	hasError   bool
}{
	{
		name:       "runner is grpc",
		runnerType: "grpc",
		hasError:   false,
	},
	{
		name:       "runner is http",
		runnerType: "http",
		hasError:   false,
	},
	{
		name:       "unknown runner",
		runnerType: "unknown",
		hasError:   true,
	},
}
var conf *config.Config

func init() {
	conf , _ = config.GetConfig()
}

func TestCreateRunner(t *testing.T) {
	for _, testCase := range ttCreate {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.hasError {
				assert.NotPanics(t, func() { _ = CreateRunner(conf, testCase.runnerType) }, "Creating runner should not panic")
			}
			if testCase.hasError {
				assert.Panics(t, func() { _ = CreateRunner(conf, testCase.runnerType) }, "Creating runner should panic")
			}
		})
	}
}
