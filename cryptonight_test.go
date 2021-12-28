package cryptonight

import (
	"encoding/hex"
	"testing"
)

var testBlock = Block{
	Header:     "82e8f378ea29d152146b6317903249751b809e97c0b6655f86e120b9de95c38a",
	POWHash:    "5a7b00a4c1ccbdc655861ce5e731fdac288aa7bda6d798a7da26f1fa37a5772a",
	Difficulty: hex.EncodeToString([]byte("1")),
	Nonce:      3212448700,
}

func TestCryptoNightHash(t *testing.T) {
	Sum(getTestData(testBlock), 24, &testBlock.POWHash)
}

func getTestData(block Block) *interface{} {
	var hashData interface{} = block.Header
	return &hashData
}
