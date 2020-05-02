package bitg

import (
	"github.com/stepollo2/stepollo-explorer/bchain"
	"github.com/stepollo2/stepollo-explorer/bchain/coins/btc"
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	// MainnetMagic is mainnet network constant
	MainnetMagic wire.BitcoinNet = 0x446d47e1
)

var (
	// MainNetParams are parser parameters for mainnet
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	// Address encoding magics
	MainNetParams.PubKeyHashAddrID = []byte{38}
	MainNetParams.ScriptHashAddrID = []byte{23}

}

// Bitgreen handle
type Bitgreen struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewBitgreenParserparams returns new Bitgreen instance
func NewBitgreenParser(params *chaincfg.Params, c *btc.Configuration) *Bitgreen {
	return &Bitgreen{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

// GetChainParams contains network parameters for the main Bitgreen network,
// the regression test Bitgreen network, the test Bitgreen network and
// the simulation test Bitgreen network, in this order
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *Bitgreen) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *Bitgreen) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
