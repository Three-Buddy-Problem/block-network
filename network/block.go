package blockchain

type (
	Block struct {
		Hash     []byte
		Data     []byte
		PrevHash []byte
		Nonce    int
	}
	BlockChain struct {
		Blocks []*Block
	}
)

func CreateBlock(data string, prevHash []byte) (b *Block) {
	b = &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}
	pow := NewProof(b)
	b.Nonce, b.Hash = pow.Run()

	return
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.PrevHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Initial Block", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}
