package chainoracles

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/services"
)

// L1Oracle provides interface for fetching L1-specific fee components if the chain is an L2.
// For example, on Optimistic Rollups, this oracle can return rollup-specific l1BaseFee
type L1Oracle interface {
	services.ServiceCtx
	
	L1GasPrice(ctx context.Context) (*assets.Wei, error)
}
