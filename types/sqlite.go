package types

type Block struct {
	Id          uint64
	BlockNumber uint64
	BlockHash   string
	BlockMiner  string
	BlockReward string
}

type Stake struct {
	Id              uint64
	BlockNumber     uint64
	TransactionHash string
	Staker          string
	Validator       string
	ActionType      string
	Amount          uint64
}
