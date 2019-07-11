package blockchain


import (
	"work_queue"
)


type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	Block
	start uint64
	end uint64
}


type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	 //TODO
	worker := work_queue.Create(uint(workers), uint(chunks))
	chunkSize := ((end - start) / chunks)
	
	if chunkSize <= 0{
		chunkSize = (end - start)
	}
	var result interface{}
	
	for i := start; i < end; i += chunkSize{
		mineWorker := new(miningWorker)
		mineWorker.Block = blk
		mineWorker.start = i
		mineWorker.end = i + chunkSize
		if mineWorker.end > end{
			mineWorker.end = end
		}
		worker.Enqueue(mineWorker)
		
	}
	for result = range worker.Results{
		if result.(MiningResult).Found{
			worker.Shutdown()
			return result.(MiningResult)
		}
	}
	return MiningResult{}
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}


func (mw miningWorker) Run() interface{} {
	result := MiningResult{}
	for i:= mw.start; i < mw.end; i++{
		mw.SetProof(i)
		if mw.ValidHash(){
			result.Proof = i
			result.Found = true
		}
	}
	return result
}

/*
func (mw miningWorker) IterateRun() MiningResult{
	result := new(MiningResult)
	i := 0
	for true{
		mw.Proof = uint64(i)
		if mw.ValidHash(){
			result.Proof = mw.Proof
			result.Found = true
			return *result
		}
		i++
	}
	return *result
}
*/