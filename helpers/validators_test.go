package helpers

import (
	"context"
	"testing"
	"time"
)

func TestCheckTimeout(t *testing.T) {
	t.Run("context without timeout", func(t *testing.T) {
		err := CheckTimeout(context.Background())
		if err == nil {
			t.Errorf("expected error")
		}
	})
	t.Run("context with invalid timeout", func(t *testing.T) {
		ctx,canc := context.WithTimeout(context.Background(),32 * time.Second)
		defer canc()
		err := CheckTimeout(ctx)
		if err == nil {
			t.Errorf("expected error")
		}
	})
	t.Run("valid context", func(t *testing.T) {
		ctx,canc := context.WithTimeout(context.Background(),5 * time.Second)
		defer canc()
		err := CheckTimeout(ctx)
		if err != nil {
			t.Errorf("expected success got %v :",err)
		}
	})
}
