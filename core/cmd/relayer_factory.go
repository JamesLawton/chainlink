package cmd

import (
	"context"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/sqlx"

	pkgcosmos "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos"
	"github.com/smartcontractkit/chainlink-relay/pkg/loop"
	pkgsolana "github.com/smartcontractkit/chainlink-solana/pkg/solana"
	pkgstarknet "github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/chains/cosmos"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/solana"
	"github.com/smartcontractkit/chainlink/v2/core/chains/starknet"
	"github.com/smartcontractkit/chainlink/v2/core/config/env"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
	evmrelayer "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type relayerFactory struct {
	logger.Logger
	*sqlx.DB
	pg.QConfig
	*plugins.LoopRegistry
	loop.GRPCOpts
}

func (r relayerFactory) NewEVM(ctx context.Context, cfg evm.GeneralConfig, ks evmrelayer.RelayerKeystore, eb pg.EventBroadcaster, mmon *utils.MailboxMonitor) (map[relay.Identifier]evmrelayer.LoopRelayAdapter, error) {
	// TODO impl EVM loop. For now always 'fallback' to an adapter and embedded chainset

	var (
		relayers map[relay.Identifier]evmrelayer.LoopRelayAdapter
	)
	ccOpts := evm.ChainSetOpts{
		Config:           cfg,
		Logger:           r.Logger,
		DB:               r.DB,
		KeyStore:         ks.Eth(),
		EventBroadcaster: eb,
		MailMon:          mmon,
	}

	var ids []utils.Big
	for _, c := range cfg.EVMConfigs() {
		c := c
		ids = append(ids, *c.ChainID)
	}
	if len(ids) > 0 {
		if err := evm.EnsureChains(r.DB, r.Logger, cfg.Database(), ids); err != nil {
			return nil, errors.Wrap(err, "failed to setup EVM chains")
		}
	}

	singleChainChainSets, err := evm.NewTOMLChainSet(ctx, ccOpts)
	if err != nil {
		return nil, err
	}
	for _, s := range singleChainChainSets {
		relayId := relay.Identifier{Network: relay.EVM, ChainID: relay.ChainID(s.Chain().ID().String())}
		relayer := evmrelayer.NewLoopRelayAdapter(evmrelayer.NewRelayer(r.DB, s, r.Logger, ks, eb), s)
		relayers[relayId] = relayer

	}

	return relayers, nil
}

func (r relayerFactory) NewSolana(ks keystore.Solana, chainCfgs solana.SolanaConfigs) (map[relay.Identifier]loop.Relayer, error) {
	var (
		solanaRelayers map[relay.Identifier]loop.Relayer
		ids            []relay.Identifier
		solLggr        = r.Logger.Named("Solana")

		signer = &keystore.SolanaSigner{ks}
	)
	for _, c := range chainCfgs {
		c := c
		ids = append(ids, relay.Identifier{Network: relay.StarkNet, ChainID: relay.ChainID(*c.ChainID)})
	}
	if len(ids) > 0 {
		if err := solana.EnsureChains(r.DB, solLggr, r.QConfig, ids); err != nil {
			return nil, fmt.Errorf("failed to setup Solana chains: %w", err)
		}
	}

	// create one relayer per chain id
	for _, chainCfg := range chainCfgs {
		relayId := relay.Identifier{Network: relay.Solana, ChainID: relay.ChainID(*chainCfg.ChainID)}
		// all the lower level APIs expect chainsets. create a single valued set per id
		singleChainCfg := solana.SolanaConfigs{chainCfg}

		if cmdName := env.SolanaPluginCmd.Get(); cmdName != "" {

			// setup the solana relayer to be a LOOP
			tomls, err := toml.Marshal(struct {
				Solana solana.SolanaConfigs
			}{Solana: singleChainCfg})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal Solana configs: %w", err)
			}

			solCmdFn, err := plugins.NewCmdFactory(r.Register, plugins.CmdConfig{
				ID:  solLggr.Name(),
				Cmd: cmdName,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create Solana LOOP command: %w", err)
			}
			solanaRelayers[relayId] = loop.NewRelayerService(solLggr, r.GRPCOpts, solCmdFn, string(tomls), signer)

		} else {
			// fallback to embedded chainset
			opts := solana.ChainSetOpts{
				Logger:   solLggr,
				KeyStore: signer,
				Configs:  solana.NewConfigs(singleChainCfg),
			}
			chainSet, err := solana.NewChainSet(opts, singleChainCfg)
			if err != nil {
				return nil, fmt.Errorf("failed to load Solana chainset: %w", err)
			}
			solanaRelayers[relayId] = relay.NewRelayerAdapter(pkgsolana.NewRelayer(solLggr, chainSet), chainSet)
		}
	}
	return solanaRelayers, nil
}

func (r relayerFactory) NewStarkNet(ks keystore.StarkNet, chainCfgs starknet.StarknetConfigs) (map[relay.Identifier]loop.Relayer, error) {
	var (
		starknetRelayers map[relay.Identifier]loop.Relayer
		ids              []string
		starkLggr        = r.Logger.Named("StarkNet")
		loopKs           = &keystore.StarknetLooppSigner{StarkNet: ks}
	)
	for _, c := range chainCfgs {
		c := c
		ids = append(ids, *c.ChainID)
	}
	if len(ids) > 0 {
		if err := starknet.EnsureChains(r.DB, starkLggr, r.QConfig, ids); err != nil {
			return nil, fmt.Errorf("failed to setup StarkNet chains: %w", err)
		}
	}

	// create one relayer per chain id
	for _, chainCfg := range chainCfgs {
		relayId := relay.Identifier{Network: relay.StarkNet, ChainID: relay.ChainID(*chainCfg.ChainID)}
		// all the lower level APIs expect chainsets. create a single valued set per id
		singleChainCfg := starknet.StarknetConfigs{chainCfg}

		if cmdName := env.StarknetPluginCmd.Get(); cmdName != "" {
			// setup the starknet relayer to be a LOOP
			tomls, err := toml.Marshal(struct {
				Starknet starknet.StarknetConfigs
			}{Starknet: singleChainCfg})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal StarkNet configs: %w", err)
			}

			starknetCmdFn, err := plugins.NewCmdFactory(r.Register, plugins.CmdConfig{
				ID:  starkLggr.Name(),
				Cmd: cmdName,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create StarkNet LOOP command: %w", err)
			}
			// the starknet relayer service has a delicate keystore dependency. the value that is passed to NewRelayerService must
			// be compatible with instantiating a starknet transaction manager KeystoreAdapter within the LOOPp executable.
			starknetRelayers[relayId] = loop.NewRelayerService(starkLggr, r.GRPCOpts, starknetCmdFn, string(tomls), loopKs)
		} else {
			// fallback to embedded chainset
			opts := starknet.ChainSetOpts{
				Logger:   starkLggr,
				KeyStore: loopKs,
				Configs:  starknet.NewConfigs(singleChainCfg),
			}
			chainSet, err := starknet.NewChainSet(opts, singleChainCfg)
			if err != nil {
				return nil, fmt.Errorf("failed to load StarkNet chainset: %w", err)
			}
			starknetRelayers[relayId] = relay.NewRelayerAdapter(pkgstarknet.NewRelayer(starkLggr, chainSet), chainSet)
		}
	}
	return starknetRelayers, nil

}

func (r relayerFactory) NewCosmos(ks keystore.Cosmos, chainCfgs cosmos.CosmosConfigs, eb pg.EventBroadcaster) (map[relay.Identifier]cosmos.LoopRelayAdapter, error) {
	var (
		relayers map[relay.Identifier]cosmos.LoopRelayAdapter
		ids      []string
		lggr     = r.Logger.Named("Cosmos")
	)
	for _, c := range chainCfgs {
		c := c
		ids = append(ids, *c.ChainID)
	}
	//TODO Cosmos doesn't seem to Ensure chains like others...
	/*
		if len(ids) > 0 {
			if err := cosmos.EnsureChains(r.DB, lggr, r.QConfig, ids); err != nil {
				return nil, fmt.Errorf("failed to setup StarkNet chains: %w", err)
			}
		}
	*/
	// create one relayer per chain id
	for _, chainCfg := range chainCfgs {
		relayId := relay.Identifier{Network: relay.Cosmos, ChainID: relay.ChainID(*chainCfg.ChainID)}
		// all the lower level APIs expect chainsets. create a single valued set per id
		// TODO: Cosmos LOOPp impl. For now, use relayer adapter

		opts := cosmos.ChainSetOpts{
			Config:           r.QConfig,
			Logger:           lggr.Named(relayId.ChainID.String()),
			DB:               r.DB,
			KeyStore:         ks,
			EventBroadcaster: eb,
		}
		opts.Configs = cosmos.NewConfigs(cosmos.CosmosConfigs{chainCfg})
		singleChainChainSet, err := cosmos.NewSingleChainSet(opts, chainCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to load Cosmos chain %q: %w", relayId, err)
		}

		relayers[relayId] = cosmos.NewLoopRelayer(pkgcosmos.NewRelayer(lggr, singleChainChainSet), singleChainChainSet)

	}
	return relayers, nil

}
