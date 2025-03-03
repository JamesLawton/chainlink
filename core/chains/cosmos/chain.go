package cosmos

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/multierr"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/adapters"
	cosmosclient "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/client"
	coscfg "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/config"
	"github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/db"

	"github.com/smartcontractkit/chainlink-relay/pkg/logger"
	"github.com/smartcontractkit/chainlink-relay/pkg/loop"
	relaytypes "github.com/smartcontractkit/chainlink-relay/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains"
	"github.com/smartcontractkit/chainlink/v2/core/chains/cosmos/cosmostxm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/cosmos/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// DefaultRequestTimeout is the default Cosmos client timeout.
// Note that while the cosmos node is processing a heavy block,
// requests can be delayed significantly (https://github.com/tendermint/tendermint/issues/6899),
// however there's nothing we can do but wait until the block is processed.
// So we set a fairly high timeout here.
// TODO(BCI-979): Remove this, or make this configurable with the updated client.
const DefaultRequestTimeout = 30 * time.Second

var (
	// ErrChainIDEmpty is returned when chain is required but was empty.
	ErrChainIDEmpty = errors.New("chain id empty")
	// ErrChainIDInvalid is returned when a chain id does not match any configured chains.
	ErrChainIDInvalid = errors.New("chain id does not match any local chains")
)

// Chain is a wrap for easy use in other places in the core node
type Chain = adapters.Chain

// ChainOpts holds options for configuring a Chain.
type ChainOpts struct {
	QueryConfig      pg.QConfig
	Logger           logger.Logger
	DB               *sqlx.DB
	KeyStore         loop.Keystore
	EventBroadcaster pg.EventBroadcaster
	Configs          types.Configs
}

func (o *ChainOpts) Validate() (err error) {
	required := func(s string) error {
		return fmt.Errorf("%s is required", s)
	}
	if o.QueryConfig == nil {
		err = multierr.Append(err, required("Config"))
	}
	if o.Logger == nil {
		err = multierr.Append(err, required("Logger'"))
	}
	if o.DB == nil {
		err = multierr.Append(err, required("DB"))
	}
	if o.KeyStore == nil {
		err = multierr.Append(err, required("KeyStore"))
	}
	if o.EventBroadcaster == nil {
		err = multierr.Append(err, required("EventBroadcaster"))
	}
	if o.Configs == nil {
		err = multierr.Append(err, required("Configs"))
	}
	return
}

func (o *ChainOpts) ConfigsAndLogger() (chains.Configs[string, db.Node], logger.Logger) {
	return o.Configs, o.Logger
}

func NewChain(cfg *CosmosConfig, opts ChainOpts) (adapters.Chain, error) {
	if !cfg.IsEnabled() {
		return nil, fmt.Errorf("cannot create new chain with ID %s, the chain is disabled", *cfg.ChainID)
	}
	c, err := newChain(*cfg.ChainID, cfg, opts.DB, opts.KeyStore, opts.QueryConfig, opts.EventBroadcaster, opts.Configs, opts.Logger)
	if err != nil {
		return nil, err
	}
	return c, nil
}

var _ adapters.Chain = (*chain)(nil)

type chain struct {
	utils.StartStopOnce
	id  string
	cfg *CosmosConfig
	txm *cosmostxm.Txm
	// TODO remove this dep after BCF-2441
	// cfs implements the loop.Relayer interface that will be removed
	cfgs types.Configs
	lggr logger.Logger
}

func newChain(id string, cfg *CosmosConfig, db *sqlx.DB, ks loop.Keystore, logCfg pg.QConfig, eb pg.EventBroadcaster, cfgs types.Configs, lggr logger.Logger) (*chain, error) {
	lggr = logger.With(lggr, "cosmosChainID", id)
	var ch = chain{
		id:   id,
		cfg:  cfg,
		cfgs: cfgs,
		lggr: logger.Named(lggr, "Chain"),
	}
	tc := func() (cosmosclient.ReaderWriter, error) {
		return ch.getClient("")
	}
	gpe := cosmosclient.NewMustGasPriceEstimator([]cosmosclient.GasPricesEstimator{
		cosmosclient.NewClosureGasPriceEstimator(func() (map[string]sdk.DecCoin, error) {
			return map[string]sdk.DecCoin{
				cfg.FeeToken(): sdk.NewDecCoinFromDec(cfg.FeeToken(), cfg.FallbackGasPrice()),
			}, nil
		}),
	}, lggr)
	ch.txm = cosmostxm.NewTxm(db, tc, *gpe, ch.id, cfg, ks, lggr, logCfg, eb)

	return &ch, nil
}

func (c *chain) Name() string {
	return c.lggr.Name()
}

func (c *chain) ID() string {
	return c.id
}

func (c *chain) Config() coscfg.Config {
	return c.cfg
}

func (c *chain) TxManager() adapters.TxManager {
	return c.txm
}

func (c *chain) Reader(name string) (cosmosclient.Reader, error) {
	return c.getClient(name)
}

// getClient returns a client, optionally requiring a specific node by name.
func (c *chain) getClient(name string) (cosmosclient.ReaderWriter, error) {
	var node db.Node
	if name == "" { // Any node
		nodes, err := c.cfgs.Nodes(c.id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get nodes")
		}
		if len(nodes) == 0 {
			return nil, errors.New("no nodes available")
		}
		nodeIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(nodes))))
		if err != nil {
			return nil, errors.Wrap(err, "could not generate a random node index")
		}
		node = nodes[nodeIndex.Int64()]
	} else { // Named node
		var err error
		node, err = c.cfgs.Node(name)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get node named %s", name)
		}
		if node.CosmosChainID != c.id {
			return nil, fmt.Errorf("failed to create client for chain %s with node %s: wrong chain id %s", c.id, name, node.CosmosChainID)
		}
	}
	client, err := cosmosclient.NewClient(c.id, node.TendermintURL, DefaultRequestTimeout, logger.Named(c.lggr, "Client."+name))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}
	c.lggr.Debugw("Created client", "name", node.Name, "tendermint-url", node.TendermintURL)
	return client, nil
}

// Start starts cosmos chain.
func (c *chain) Start(ctx context.Context) error {
	return c.StartOnce("Chain", func() error {
		c.lggr.Debug("Starting")
		return c.txm.Start(ctx)
	})
}

func (c *chain) Close() error {
	return c.StopOnce("Chain", func() error {
		c.lggr.Debug("Stopping")
		return c.txm.Close()
	})
}

func (c *chain) Ready() error {
	return multierr.Combine(
		c.StartStopOnce.Ready(),
		c.txm.Ready(),
	)
}

func (c *chain) HealthReport() map[string]error {
	return map[string]error{
		c.Name(): multierr.Combine(
			c.StartStopOnce.Healthy(),
			c.txm.Healthy()),
	}
}

// ChainService interface
func (c *chain) GetChainStatus(ctx context.Context) (relaytypes.ChainStatus, error) {
	toml, err := c.cfg.TOMLString()
	if err != nil {
		return relaytypes.ChainStatus{}, err
	}
	return relaytypes.ChainStatus{
		ID:      c.id,
		Enabled: *c.cfg.Enabled,
		Config:  toml,
	}, nil
}
func (c *chain) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) (stats []relaytypes.NodeStatus, nextPageToken string, total int, err error) {
	return internal.ListNodeStatuses(int(pageSize), pageToken, c.listNodeStatuses)
}

func (c *chain) Transact(ctx context.Context, from, to string, amount *big.Int, balanceCheck bool) error {
	return chains.ErrLOOPPUnsupported
}

// TODO BCF-2602 statuses are static for non-evm chain and should be dynamic
func (c *chain) listNodeStatuses(start, end int) ([]relaytypes.NodeStatus, int, error) {
	stats := make([]relaytypes.NodeStatus, 0)
	total := len(c.cfg.Nodes)
	if start >= total {
		return stats, total, internal.ErrOutOfRange
	}
	if end > total {
		end = total
	}
	nodes := c.cfg.Nodes[start:end]
	for _, node := range nodes {
		stat, err := nodeStatus(node, c.id)
		if err != nil {
			return stats, total, err
		}
		stats = append(stats, stat)
	}
	return stats, total, nil
}
