package consensus

import (
	"time"
)

var (
	proposers = [][20]byte{
		{0xa1}, {0xb2}, {0xc3}, {0xd4}, {0xe5},
	}
	blockTime  = 2 * time.Second
	lastBlock  = time.Now().Unix()
)

func GetProposer(blockNum uint64) [20]byte {
	return proposers[blockNum%uint64(len(proposers))]
}

func ShouldPropose(blockNum uint64) bool {
	return time.Now().Unix() > lastBlock+2
}
