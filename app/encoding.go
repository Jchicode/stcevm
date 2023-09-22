package app

import (
	evmenc "github.com/evmos/ethermint/encoding"

	"github.com/cosmos/cosmos-sdk/simapp/params"
)

// MakeEncodingConfig creates the EncodingConfig for realio network
func MakeEncodingConfig() params.EncodingConfig {
	return evmenc.MakeConfig(ModuleBasics)
}
