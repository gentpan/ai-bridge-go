package providers

import (
	"context"

	"github.com/gentpan/ai-bridge-go/internal/gateway"
)

type Provider interface {
	Chat(ctx context.Context, request gateway.ChatRequest) (*gateway.ChatResponse, error)
}
