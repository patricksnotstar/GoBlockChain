package blockchain

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	blk := new(Block)
	blk.Generation = 0
	blk.Difficulty = difficulty
	for i := 0; i < 32; i++{
		blk.PrevHash = append(blk.PrevHash, '\x00')
	}
	return *blk
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	next := new(Block)
	next.Generation = prev_block.Generation + 1
	next.Difficulty = prev_block.Difficulty
	next.Data = data
	next.PrevHash = prev_block.Hash
	return *next
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	s := (hex.EncodeToString(blk.PrevHash) + ":" + fmt.Sprint(blk.Generation) + ":" + fmt.Sprint(blk.Difficulty) + ":" + blk.Data + ":" + fmt.Sprint(blk.Proof))
	blk.Hash = []byte(s)
	hesh := sha256.New()
	hesh.Write(blk.Hash)
	hash := hesh.Sum(nil)
	return hash
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	arr := blk.Hash[len(blk.Hash) - int(nBytes) : len(blk.Hash)]
	for i:= 0; i < len(arr) - 1; i++{
		if arr[i] != '\x00'{
			return false
		}
	}
	hashByte := blk.Hash[(len(blk.Hash) - int(nBytes) - 1)]
	return  hashByte % (1 << nBits) == 0
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
