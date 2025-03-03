package relay

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-relay/pkg/types"
)

func TestIdentifier_UnmarshalString(t *testing.T) {
	type fields struct {
		Network Network
		ChainID ChainID
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		want    fields
		args    args
		wantErr bool
	}{
		{name: "evm",
			args:    args{s: "evm.1"},
			wantErr: false,
			want:    fields{Network: EVM, ChainID: "1"},
		},
		{name: "bad network",
			args:    args{s: "notANetwork.1"},
			wantErr: true,
		},
		{name: "bad pattern",
			args:    args{s: "evm_1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ID{}
			err := i.UnmarshalString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Identifier.UnmarshalString() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want.Network, i.Network)
			assert.Equal(t, tt.want.ChainID, i.ChainID)
		})
	}
}

func TestNewID(t *testing.T) {
	type args struct {
		n Network
		c ChainID
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		{name: "good evm",
			args: args{n: EVM, c: "1"},
			want: ID{Network: EVM, ChainID: "1"},
		},
		{name: "bad evm",
			args:    args{n: EVM, c: "not a number"},
			want:    ID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.n, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "got id %v", got)
		})
	}
}

type staticMedianProvider struct {
	types.MedianProvider
}

type staticFunctionsProvider struct {
	types.FunctionsProvider
}

type staticMercuryProvider struct {
	types.MercuryProvider
}

type mockRelayer struct {
	types.Relayer
}

func (m *mockRelayer) NewMedianProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.MedianProvider, error) {
	return staticMedianProvider{}, nil
}

func (m *mockRelayer) NewFunctionsProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.FunctionsProvider, error) {
	return staticFunctionsProvider{}, nil
}

func (m *mockRelayer) NewMercuryProvider(rargs types.RelayArgs, pargs types.PluginArgs) (types.MercuryProvider, error) {
	return staticMercuryProvider{}, nil
}

type mockRelayerExt struct {
	RelayerExt
}

func isType[T any](p any) bool {
	_, ok := p.(T)
	return ok
}

func TestRelayerServerAdapter(t *testing.T) {
	r := &mockRelayer{}
	sa := NewRelayerServerAdapter(r, mockRelayerExt{})

	testCases := []struct {
		ProviderType string
		Test         func(p any) bool
		Error        string
	}{
		{
			ProviderType: string(types.Median),
			Test:         isType[types.MedianProvider],
		},
		{
			ProviderType: string(types.Functions),
			Test:         isType[types.FunctionsProvider],
		},
		{
			ProviderType: string(types.Mercury),
			Test:         isType[types.MercuryProvider],
		},
		{
			ProviderType: "unknown",
			Error:        "provider type not supported",
		},
		{
			ProviderType: string(types.GenericPlugin),
			Error:        "unexpected call to NewPluginProvider",
		},
	}

	for _, tc := range testCases {
		pp, err := sa.NewPluginProvider(
			context.Background(),
			types.RelayArgs{ProviderType: tc.ProviderType},
			types.PluginArgs{},
		)

		if tc.Error != "" {
			assert.ErrorContains(t, err, tc.Error)
		} else {
			assert.NoError(t, err)
			assert.True(t, tc.Test(pp))
		}
	}
}
