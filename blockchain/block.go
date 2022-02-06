package blockchain

type BlockChain struct {
	Blocks []*Block
}
type Block struct {
	Hash     []byte //hash is created by using the data and th prev hash
	Data     []byte // data cna be anything eg ledger, document
	PrevHash []byte //prevhash is the reference to the current blockhash that which block is connected to it int the blockchain
	Nonce    int
}

//{
// now creating the hash generator function which would be using the block data as well as the prev hash
// it will be requiring the libraries such as bytes

// since it is done by proof of work algorithm in proof.go  ( func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

// 	//sha256 algorithm is simple as compared to hash calculation in rela world blockchian
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// })

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//} calculating hash

// this method will allow to add a block to the chian

func (chain *BlockChain) AddBlock(data string) {
	//initially adding the prev block to the chain
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)

}

//for creating the first block or genisis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//create the initial blockchain which would start from creating the genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
