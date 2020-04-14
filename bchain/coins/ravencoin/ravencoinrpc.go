package ravencoin

import (
	"encoding/json"
	"github.com/grupokindynos/coins-explorer/bchain"
	"github.com/grupokindynos/coins-explorer/bchain/coins/btc"

	"github.com/golang/glog"
)

// RavencoinRPC is an interface to JSON-RPC bitcoind service.
type RavencoinRPC struct {
	*btc.BitcoinRPC
}

// NewRavencoinRPC returns new RavencoinRPC instance.
func NewRavencoinRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &RavencoinRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

// Initialize initializes RavencoinRPC instance.
func (b *RavencoinRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewRavencoinParser(params, b.ChainConfig)

	// parameters for getInfo request
	b.Testnet = false
	b.Network = "livenet"

	glog.Info("rpc: block chain ", params.Name)

	return nil
}
