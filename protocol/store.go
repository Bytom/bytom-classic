package protocol

import (
	"github.com/bytom/bytom-classic/database/storage"
	"github.com/bytom/bytom-classic/protocol/bc"
	"github.com/bytom/bytom-classic/protocol/bc/types"
	"github.com/bytom/bytom-classic/protocol/state"
)

// Store provides storage interface for blockchain data
type Store interface {
	BlockExist(*bc.Hash) bool

	GetBlock(*bc.Hash) (*types.Block, error)
	GetStoreStatus() *BlockStoreState
	GetTransactionStatus(*bc.Hash) (*bc.TransactionStatus, error)
	GetTransactionsUtxo(*state.UtxoViewpoint, []*bc.Tx) error
	GetUtxo(*bc.Hash) (*storage.UtxoEntry, error)

	LoadBlockIndex(uint64) (*state.BlockIndex, error)
	SaveBlock(*types.Block, *bc.TransactionStatus) error
	SaveChainStatus(*state.BlockNode, *state.UtxoViewpoint) error
}

// BlockStoreState represents the core's db status
type BlockStoreState struct {
	Height uint64
	Hash   *bc.Hash
}
