package blockchainingo

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
