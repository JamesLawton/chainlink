package headtracker

import (
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	commontypes "github.com/smartcontractkit/chainlink/v2/common/types"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
)

type heads[H commontypes.HeadTrackerHead[H, BLOCK_HASH], BLOCK_HASH commontypes.Hashable] struct {
	heads  []H
	mu     sync.RWMutex
	getNil func() H
}

func NewEvmHeads() *heads[*evmtypes.Head, common.Hash] {
	return NewHeads[*evmtypes.Head, common.Hash](
		func() *evmtypes.Head { return nil },
	)
}

func NewHeads[
	H commontypes.HeadTrackerHead[H, BLOCK_HASH],
	BLOCK_HASH commontypes.Hashable,
](getNil func() H) *heads[H, BLOCK_HASH] {
	return &heads[H, BLOCK_HASH]{
		getNil: getNil,
	}
}

func (h *heads[H, BLOCK_HASH]) LatestHead() H {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.heads) == 0 {
		return h.getNil()
	}
	return h.heads[0]
}

func (h *heads[H, BLOCK_HASH]) HeadByHash(hash BLOCK_HASH) *H {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, head := range h.heads {
		if head.BlockHash() == hash {
			return &head
		}
	}
	return nil
}

func (h *heads[H, BLOCK_HASH]) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.heads)
}

func (h *heads[H, BLOCK_HASH]) AddHeads(historyDepth uint, newHeads ...H) {
	h.mu.Lock()
	defer h.mu.Unlock()

	headsMap := make(map[BLOCK_HASH]H, len(h.heads)+len(newHeads))
	for _, head := range append(h.heads, newHeads...) {
		if head.BlockHash() == head.GetParentHash() {
			// shouldn't happen but it is untrusted input
			continue
		}
		// copy all head objects to avoid races when a previous head chain is used
		// elsewhere (since we mutate Parent here)
		headCopy := head
		headCopy.SetParent(h.getNil()) // always build it from scratch in case it points to a head too old to be included
		// map eliminates duplicates
		headsMap[head.BlockHash()] = headCopy
	}

	heads := make([]H, len(headsMap))
	// unsorted unique heads
	{
		var i int
		for _, head := range headsMap {
			heads[i] = head
			i++
		}
	}

	// sort the heads
	sort.SliceStable(heads, func(i, j int) bool {
		// sorting from the highest number to lowest
		return heads[i].BlockNumber() > heads[j].BlockNumber()
	})

	// cut off the oldest
	if uint(len(heads)) > historyDepth {
		heads = heads[:historyDepth]
	}

	// assign parents
	for i := 0; i < len(heads)-1; i++ {
		head := heads[i]
		parent, exists := headsMap[head.GetParentHash()]
		if exists {
			head.SetParent(parent)
		}
	}

	// set
	h.heads = heads
}
