package blockchain

import (
	"encoding/hex"
)
	
type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	// TODO
	chain.Chain = append(chain.Chain, blk)
}



func (chain Blockchain) IsValid() bool {
	// TODO
	result := make(chan bool, len(chain.Chain))
	if len(chain.Chain) > 0{
		if chain.Chain[0].Generation != 0 {
			return false
		}
		nBytes := chain.Chain[0].Difficulty / 8
		
		for i := 0; i < 32; i++{
			if chain.Chain[0].PrevHash[i] != '\x00'{
				return false
			}
		}
		for i := 1; i < (len(chain.Chain) - 1); i++{
			if (chain.Chain[i].Difficulty != chain.Chain[0].Difficulty) || 
				(chain.Chain[i].Generation != (chain.Chain[i - 1].Generation + 1)) ||
				(hex.EncodeToString(chain.Chain[i].PrevHash) != hex.EncodeToString(chain.Chain[i-1].Hash)) ||
				(hex.EncodeToString(chain.Chain[i].Hash) != hex.EncodeToString(chain.Chain[i].CalcHash())){
				return false
			}
			
			arr := chain.Chain[i].Hash[len(chain.Chain[i].Hash) - int(nBytes) : len(chain.Chain[i].Hash)]
			go checkHash(nBytes, arr,  result)
		}
	}
	
	res := true
	for res = range result{
		if res == false{
			return res
		}
	}
	return res
	
	
}

func checkHash(nBytes uint8, hash []byte, result chan bool){
	for i:= 0; i < len(hash) - 1; i++{
		if hash[i] != '\x00'{
			result <- false
			close(result)
			return
		}
	}
	result <- true
	defer close(result)
}



/*
func (chain Blockchain) IsValid() bool {
	// TODO
	
	if len(chain.Chain) > 0{
		if chain.Chain[0].Generation != 0 {
			return false
		}
		for i := 0; i < 32; i++{
			if chain.Chain[0].PrevHash[i] != '\x00'{
				return false
			}
		}
		for i := 1; i < (len(chain.Chain) - 1); i++{
			if (chain.Chain[i].Difficulty != chain.Chain[0].Difficulty) || 
				(chain.Chain[i].Generation != (chain.Chain[i - 1].Generation + 1)) ||
				(hex.EncodeToString(chain.Chain[i].PrevHash) != hex.EncodeToString(chain.Chain[i-1].Hash)) ||
				(hex.EncodeToString(chain.Chain[i].Hash) != hex.EncodeToString(chain.Chain[i].CalcHash())){
				return false
			}
			nBytes := chain.Chain[i].Difficulty / 8
			arr := chain.Chain[i].Hash[len(chain.Chain[i].Hash) - int(nBytes) : len(chain.Chain[i].Hash)]
			for i:= 0; i < len(arr) - 1; i++{
				if arr[i] != '\x00'{
					return false
				}
			}
		}
	}
	
	return true
	
	
}
*/
