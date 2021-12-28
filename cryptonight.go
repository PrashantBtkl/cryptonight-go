package cryptonight

import (
	"math/big"
)

type Block struct {
	// Properties may vary
	Header     string
	SeedHash   string
	Height     uint64
	Nonce      uint64
	Difficulty string
}

func Sum(Block) (*big.Int, error) {
	return big.NewInt(0), nil
}

func Verify(Block) (bool, error) {
	return true, nil
}
