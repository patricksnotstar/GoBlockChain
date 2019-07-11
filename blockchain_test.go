package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
	"encoding/hex"
)

// TODO: some useful tests of Blocks
func TestInitial(t *testing.T){
	diff := uint8(8)
	newBlock := Initial(diff)
	assert.Equal(t, diff, newBlock.Difficulty, "Difficulty is wrong")
	assert.Equal(t, uint64(0), newBlock.Generation, "Generation is wrong")
	//assert.Equal(t, []byte("000000000000000000000000000000000000000000000000000000000000000"), newBlock.PrevHash, "PrevHash is wrong")
}

func TestNext (t *testing.T){
	diff := uint8(8)
	prev := Initial(diff)
	this := prev.Next("message")
	var hash []byte
	assert.Equal(t, diff, this.Difficulty, "Difficulty is wrong")
	assert.Equal(t, uint64(1), this.Generation, "Generation is wrong")
	assert.Equal(t, hash, this.PrevHash, "PrevHash is wrong")
	assert.Equal(t, "message", this.Data, "data is wrong")
	
}


func TestCalcHash (t *testing.T){
	blk := Initial(uint8(16))
	blk.SetProof(56231)
	blk2 := blk.Next("message")
	blk2.SetProof(2159)
	assert.Equal(t, "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000", hex.EncodeToString(blk.CalcHash()), string(blk.Hash))
	assert.Equal(t, "9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000", hex.EncodeToString(blk2.CalcHash()), string(blk2.Hash))
}


func TestValidHash(t *testing.T){
	b0 := Initial(19)
	b0.SetProof(87745)
	assert.Equal(t, true, b0.ValidHash(), "ValidHash is wrong")
	b1 := b0.Next("hash example 1234")
	b1.SetProof(1407891)
	assert.Equal(t, true, b1.ValidHash(), "ValidHash is wrong")
}


func TestMine(t *testing.T){
	b0 := Initial(7)
	b0.Mine(2)
	assert.Equal(t, b0.Proof, uint64(385), "Proof is wrong")
	assert.Equal(t, hex.EncodeToString(b0.Hash), "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300", "Hash is wrong")
}

func TestAdd(t *testing.T){
	bc := new(Blockchain)
	b0 := Initial(17)
	b0.Mine(1)
	b1 := b0.Next("message")
	b1.Mine(1)
	b2 := b1.Next("message2")
	b2.Mine(1)
	bc.Add(b0)
	bc.Add(b1)
	bc.Add(b2)
	
	assert.Equal(t, b0.Generation, bc.Chain[0].Generation, "Wrong Generation for first block")
	assert.Equal(t, b1.Generation, bc.Chain[1].Generation, "Wrong Generation for second block")
	assert.Equal(t, b2.Generation, bc.Chain[2].Generation, "Wrong Generation for third block")
}

func TestIsValid(t *testing.T){
	bc := new(Blockchain)
	b0 := Initial(17)
	b0.Mine(2)
	b1 := b0.Next("message")
	b1.Mine(2)
	b2 := b1.Next("message2")
	b2.Mine(2)
	bc.Add(b0)
	bc.Add(b1)
	bc.Add(b2)
	assert.Equal(t, true, bc.IsValid(), "IsValid is not valid")
}


