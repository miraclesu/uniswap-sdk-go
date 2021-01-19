package constants

// ChainID network chain id
type ChainID int

//go:generate stringer -type=ChainID -linecomment
const (
	Mainnet ChainID = 1
	Ropsten ChainID = 3
	Rinkeby ChainID = 4
	Goerli  ChainID = 5
	Kovan   ChainID = 42
)
