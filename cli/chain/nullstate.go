package chain

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
	"github.com/buuzcoin/blockchain/trie"
	"github.com/buuzcoin/node/db"
)

// InitNullState writes null trie states to local storage
func InitNullState(localStorage *db.LocalStorage) error {
	return localStorage.Env.Update(func(txn *lmdb.Txn) error {
		return txn.Put(localStorage.State, trie.NullTrieHash, trie.NullTrie.ToBytes(), 0)
	})
}
