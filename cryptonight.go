package cryptonight

import (
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"
)

type Block struct {
	// Properties may vary
	Header     string
	POWHash    string
	Height     uint64
	Nonce      uint64
	Difficulty string
}

func Sum(data *interface{}, length uint8, hash *string) (*big.Int, error) {
	// Currently using sha3 library for setting up tests
	// TODO : Use ryo implementation of cryptonight_gpu algo
	sha3.keccakF1600(data)
	fmt.Println(data)
	return big.NewInt(0), nil
}

func Verify(block Block) (bool, error) {
	return true, nil
}
