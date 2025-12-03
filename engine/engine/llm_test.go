package engine

import (
	"context"
	"testing"
)

func TestLLM(t *testing.T) {
	ctx := context.Background()
	InitLLM(ctx, "sz002594", "20251117")
}
