package bitg

import (
	"encoding/json"
	"github.com/stepollo2/stepollo-explorer/bchain"
	"github.com/stepollo2/stepollo-explorer/bchain/coins/btc"

	"github.com/golang/glog"
)

// BitGreenRPC is an interface to JSON-RPC bitcoind service.
type BitgreenRPC struct {
	*btc.BitcoinRPC
}

// NewBitgreenRPC returns new BitGreenRPC instance.
func NewBitgreenRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &BitgreenRPC{
		b.(*btc.BitcoinRPC),
	}

	return s, nil
}

// Initialize initializes BGoldRPC instance.
func (b *BitgreenRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewBitgreenParser(params, b.ChainConfig)

	// parameters for getInfo request
	b.Testnet = false
        b.Network = "livenet"

	glog.Info("rpc: block chain ", params.Name)
	return nil
}
