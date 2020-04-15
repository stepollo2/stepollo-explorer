package aryacoin

import (
	"github.com/grupokindynos/coins-explorer/bchain"
	"github.com/grupokindynos/coins-explorer/bchain/coins/btc"
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xe7eae188
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	// Address encoding magics
	MainNetParams.AddressMagicLen = 1
	MainNetParams.PubKeyHashAddrID = []byte{23}
	MainNetParams.ScriptHashAddrID = []byte{50}
}

// AryacoinParser handle
type AryacoinParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewAryacoinParser returns new AryacoinParser instance
func NewAryacoinParser(params *chaincfg.Params, c *btc.Configuration) *AryacoinParser {
	return &AryacoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

// GetChainParams contains network parameters for the main SnowGem network,
// the regression test SnowGem network, the test SnowGem network and
// the simulation test SnowGem network, in this order
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

// PackTx packs transaction to byte array using protobuf
func (p *AryacoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *AryacoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
