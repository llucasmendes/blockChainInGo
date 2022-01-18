package blockchainingo

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block hash with current block data

	hash := sha256.Sum256(info)
	//The actual hashing algorithm

	b.Hash = hash[:]
}
