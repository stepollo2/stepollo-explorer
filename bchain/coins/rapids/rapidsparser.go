package rapids

import (
	"github.com/grupokindynos/coins-explorer/bchain"
	"github.com/grupokindynos/coins-explorer/bchain/coins/btc"
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xcbf5a261
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	// Address encoding magics
	MainNetParams.AddressMagicLen = 1
	MainNetParams.PubKeyHashAddrID = []byte{61}
	MainNetParams.ScriptHashAddrID = []byte{6}
}

// RapidsParser handle
type RapidsParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewRapidsParser returns new RapidsParser instance
func NewRapidsParser(params *chaincfg.Params, c *btc.Configuration) *RapidsParser {
	return &RapidsParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

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
func (p *RapidsParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *RapidsParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
