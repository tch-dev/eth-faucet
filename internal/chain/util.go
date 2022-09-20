package chain

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}
	return !checksummed || common.HexToAddress(address).Hex() == address
}

func EtherToWei(amount int64) *big.Int {
	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Mul(big.NewInt(amount), ether)
}

func PayoutToWei(payout string) *big.Int{
	amount, err := strconv.ParseFloat(payout, 64)
	_ = err
	wei := big.NewInt(int64(amount * 1000000000000000000))
	return wei
}